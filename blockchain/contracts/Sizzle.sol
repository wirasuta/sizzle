// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.6 <0.9.0;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract Sizzle {
    struct CertMetadata {
        address owner;
        string domain;
        string pubKey;
        int reputation;
        int reputationMax;
        bool valid;
    }

    struct CertParticipation {
        EnumerableSet.AddressSet endUser;
        EnumerableSet.AddressSet endorser;
        EnumerableSet.AddressSet denier;
    }

    struct Peer {
        address addr;
        int reputation;
    }

    int REPUTATION_THRESHOLD = 2;
    int PEER_REPUTATION_RATING_COUNT = 50;
    int PEER_REPUTATION_MAX = 100;
    int PEER_REPUTATION_PRECISION = 10000;

    mapping(string => CertMetadata) certs;
    mapping(string => CertParticipation) participations;
    mapping(address => Peer) peers;
    mapping(address => int[]) peersRating;
    
    event CertPublishRequestCreated(CertMetadata cert);
    event CertValid(CertMetadata cert);
    event CertEndorsed(string domain, Peer peer);
    event CertDenied(string domain, Peer peer);
    
    function certPublishRequest(string memory domain, string memory pubKey) public {
        require(certs[domain].owner == address(0));

        CertMetadata storage c = certs[domain];
        c.owner = msg.sender;
        c.domain = domain;
        c.pubKey = pubKey;
        c.reputation = 0;
        c.reputationMax = 0;
        c.valid = false;

        // emit CertPublishRequestCreated(c);
    }

    function calculateCertValidity(string memory domain) private {
        CertMetadata storage c = certs[domain];
        
        if (!c.valid && c.reputation * REPUTATION_THRESHOLD >= c.reputationMax) {
            c.valid = true;
            // emit CertValid(c);
        }
    }

    function certEndorseByPeer(string memory domain) public {
        Peer storage peer = peers[msg.sender];
        CertMetadata storage cert = certs[domain];
        CertParticipation storage participation = participations[domain];

        require(cert.owner != msg.sender);
        require(!EnumerableSet.contains(participation.endUser, msg.sender));
        require(!EnumerableSet.contains(participation.endorser, msg.sender));
        require(!EnumerableSet.contains(participation.denier, msg.sender));
        require(peer.addr != address(0));

        cert.reputation += peer.reputation;
        cert.reputationMax += PEER_REPUTATION_MAX;
        EnumerableSet.add(participation.endorser, msg.sender);

        calculateCertValidity(domain);

        // emit CertEndorsed(domain, peer);
    }

    function certDenyByPeer(string memory domain) public {
        Peer storage peer = peers[msg.sender];
        CertMetadata storage cert = certs[domain];
        CertParticipation storage participation = participations[domain];

        require(cert.owner != msg.sender);
        require(!EnumerableSet.contains(participation.endUser, msg.sender));
        require(!EnumerableSet.contains(participation.endorser, msg.sender));
        require(!EnumerableSet.contains(participation.denier, msg.sender));
        require(peer.addr != address(0));

        cert.reputation -= peer.reputation;
        cert.reputationMax += PEER_REPUTATION_MAX;
        EnumerableSet.add(participation.denier, msg.sender);

        calculateCertValidity(domain);

        // emit CertDenied(domain, peer);
    }

    function calculatePeerReputation(address addr) private {
        int ratingLen = int(peersRating[addr].length);
        int startIdx = ratingLen - PEER_REPUTATION_RATING_COUNT;
        if (startIdx < 0) {
            startIdx = 0;
        }

        int sumF = (PEER_REPUTATION_RATING_COUNT/2) * (1 + PEER_REPUTATION_RATING_COUNT);
        int sumR = 0;
        for (int i = startIdx; i < ratingLen; i++) {
            int p = (PEER_REPUTATION_PRECISION * i / PEER_REPUTATION_RATING_COUNT) / sumF;
            sumR += p * peersRating[addr][uint(i)];
        }

        peers[addr].reputation = sumR / (PEER_REPUTATION_PRECISION / PEER_REPUTATION_MAX);
    }

    function certEndorseByUser(string memory domain) public {
        CertMetadata storage cert = certs[domain];
        CertParticipation storage participation = participations[domain];

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
        require(peers[msg.sender].addr == address(0));
        
        Peer storage peer = peers[msg.sender];
        peer.addr = msg.sender;
        peer.reputation = 0;
    }
}