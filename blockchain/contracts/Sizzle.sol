// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.6 <0.9.0;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract Sizzle {
    enum CertStatus { Invalid, Valid, Revoked }

    struct CertMetadata {
        address owner;
        string domain;
        string pubKey;
        int reputation;
        int reputationMax;
        CertStatus status;
    }

    struct CertParticipation {
        EnumerableSet.AddressSet endUser;
        EnumerableSet.AddressSet endorser;
        EnumerableSet.AddressSet denier;
    }

    struct PeerMetadata{
        address addr;
        int reputation;
    }

    int REPUTATION_THRESHOLD = 2;
    int PEER_REPUTATION_RATING_COUNT = 30;
    int PEER_REPUTATION_MAX = 100;
    int PEER_REPUTATION_PRECISION = 10000;

    mapping(string => CertMetadata) certs;
    mapping(string => CertParticipation) participations;
    mapping(address => PeerMetadata) peers;
    mapping(address => int[]) peersRating;
    
    event CertPublishRequestCreated(address owner, string domain, string pubKey);
    event CertRekeyed(address owner, string domain, string pubKey);
    event CertRevoked(address owner, string domain);
    event CertValid(address owner, string domain, string pubKey, int reputation);
    event CertEndorsed(string domain, address peer);
    event CertDenied(string domain, address peer);

    constructor() {
        PeerMetadata storage peer = peers[msg.sender];
        peer.addr = msg.sender;
        peer.reputation = PEER_REPUTATION_MAX;

        for (int i = 0; i < PEER_REPUTATION_RATING_COUNT; i++) {
            peersRating[msg.sender].push(1);
        }
    }
    
    function certPublishRequest(string memory domain, string memory pubKey) public {
        CertMetadata storage c = certs[domain];
        require(c.owner == address(0));

        c.owner = msg.sender;
        c.domain = domain;
        c.pubKey = pubKey;
        c.reputation = 0;
        c.reputationMax = 0;
        c.status = CertStatus.Invalid;

        emit CertPublishRequestCreated(c.owner, c.domain, c.pubKey);
    }

    function certRekey(string memory domain, string memory pubKey) public {
        CertMetadata storage c = certs[domain];
        require(c.owner == address(0));
        
        c.pubKey = pubKey;
        emit CertRekeyed(c.owner, c.domain, c.pubKey);
    }

    function certRevoke(string memory domain) public {
        CertMetadata storage c = certs[domain];
        require(c.owner == address(0));

        c.status = CertStatus.Revoked;
        emit CertRevoked(c.owner, c.domain);
    }

    function calculateCertValidity(string memory domain) private {
        CertMetadata storage c = certs[domain];
        
        if (c.status != CertStatus.Valid && c.reputation * REPUTATION_THRESHOLD >= c.reputationMax) {
            c.status = CertStatus.Valid;
            emit CertValid(c.owner, c.domain, c.pubKey, c.reputation);
        }
    }

    function certEndorseByPeer(string memory domain) public {
        PeerMetadata storage peer = peers[msg.sender];
        CertMetadata storage cert = certs[domain];
        CertParticipation storage participation = participations[domain];

        require(cert.owner != address(0));
        require(cert.owner != msg.sender);
        require(!EnumerableSet.contains(participation.endUser, msg.sender));
        require(!EnumerableSet.contains(participation.endorser, msg.sender));
        require(!EnumerableSet.contains(participation.denier, msg.sender));
        require(peer.addr != address(0));

        cert.reputation += peer.reputation;
        cert.reputationMax += PEER_REPUTATION_MAX;
        EnumerableSet.add(participation.endorser, msg.sender);

        calculateCertValidity(domain);

        emit CertEndorsed(domain, peer.addr);
    }

    function certDenyByPeer(string memory domain) public {
        PeerMetadata storage peer = peers[msg.sender];
        CertMetadata storage cert = certs[domain];
        CertParticipation storage participation = participations[domain];

        require(cert.owner != address(0));
        require(cert.owner != msg.sender);
        require(!EnumerableSet.contains(participation.endUser, msg.sender));
        require(!EnumerableSet.contains(participation.endorser, msg.sender));
        require(!EnumerableSet.contains(participation.denier, msg.sender));
        require(peer.addr != address(0));

        cert.reputation -= peer.reputation;
        cert.reputationMax += PEER_REPUTATION_MAX;
        EnumerableSet.add(participation.denier, msg.sender);

        calculateCertValidity(domain);

        emit CertDenied(domain, peer.addr);
    }

    function calculatePeerReputation(address addr) private {
        int[] storage peerRating = peersRating[addr];
        int ratingLen = int(peerRating.length);
        int startIdx = ratingLen - PEER_REPUTATION_RATING_COUNT;
        if (startIdx < 0) {
            startIdx = 0;
        }
        int significantRatingLen = ratingLen - startIdx;

        int sumF = (1 + significantRatingLen) / 2;
        int sumR = 0;
        if (sumF != 0) {
            for (int i = startIdx; i < ratingLen; i++) {
                int p = (PEER_REPUTATION_PRECISION * (i + 1) / significantRatingLen) / sumF;
                sumR += p * peerRating[uint(i)];
            }
        }

        PeerMetadata storage peer = peers[addr];
        peer.reputation = sumR / (PEER_REPUTATION_PRECISION / PEER_REPUTATION_MAX);
    }

    function certEndorseByUser(string memory domain) public {
        CertMetadata storage cert = certs[domain];
        CertParticipation storage participation = participations[domain];

        require(cert.owner != address(0));
        require(cert.owner != msg.sender);
        require(!EnumerableSet.contains(participation.endUser, msg.sender));
        require(!EnumerableSet.contains(participation.endorser, msg.sender));
        require(!EnumerableSet.contains(participation.denier, msg.sender));

        int rating = 1;
        uint endorserLen = EnumerableSet.length(participation.endorser);
        for (uint i = 0; i < endorserLen; i++) {
            address addr = EnumerableSet.at(participation.endorser, i);
            peersRating[addr].push(rating);
            calculatePeerReputation(addr);
        }
        
        uint denierLen = EnumerableSet.length(participation.denier);
        for (uint i = 0; i < denierLen; i++) {
            address addr = EnumerableSet.at(participation.denier, i);
            peersRating[addr].push(-1 * rating);
            calculatePeerReputation(addr);
        }

        EnumerableSet.add(participation.endUser, msg.sender);
    }

    function certDenyByUser(string memory domain) public {
        CertMetadata storage cert = certs[domain];
        CertParticipation storage participation = participations[domain];

        require(cert.owner != address(0));
        require(cert.owner != msg.sender);
        require(!EnumerableSet.contains(participation.endUser, msg.sender));
        require(!EnumerableSet.contains(participation.endorser, msg.sender));
        require(!EnumerableSet.contains(participation.denier, msg.sender));

        int rating = -1;
        uint endorserLen = EnumerableSet.length(participation.endorser);
        for (uint i = 0; i < endorserLen; i++) {
            address addr = EnumerableSet.at(participation.endorser, i);
            peersRating[addr].push(rating);
            calculatePeerReputation(addr);
        }

        uint denierLen = EnumerableSet.length(participation.denier);
        for (uint i = 0; i < denierLen; i++) {
            address addr = EnumerableSet.at(participation.denier, i);
            peersRating[addr].push(-1 * rating);
            calculatePeerReputation(addr);
        }

        EnumerableSet.add(participation.endUser, msg.sender);
    }

    function certQuery(string memory domain) public view returns (CertMetadata memory cert) {
        return certs[domain];
    }

    function peerRegister() public {
        PeerMetadata storage peer = peers[msg.sender];
        require(peer.addr == address(0));
        
        peer.addr = msg.sender;
        peer.reputation = 0;
    }
    
    function peerQuery(address addr) public view returns (PeerMetadata memory peer) {
        return peers[addr];
    }
}