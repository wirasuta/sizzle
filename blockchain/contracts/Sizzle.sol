// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.6 <0.9.0;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract Sizzle {
    struct CertMetadata {
        address owner;
        string domain;
        string cert;
        int reputation;
        int reputationMax;
        bool valid;
        EnumerableSet.AddressSet endorser;
        EnumerableSet.AddressSet denier;
    }

    struct Peer {
        address p;
        int reputation;
    }

    int reputationThreshold = 2;

    mapping(string => CertMetadata) certs;
    mapping(address => Peer) peers;
    
    event CertPublishRequestCreated(string domain);
    event CertValid(string domain);
    event CertEndorsed(string domain, Peer peer);
    event CertDenied(string domain, Peer peer);
    
    function certPublishRequest(string memory domain, string memory certStr) public {
        require(certs[domain].owner == address(0));

        CertMetadata storage c = certs[domain];
        c.owner = msg.sender;
        c.domain = domain;
        c.cert = certStr;
        c.reputation = 0;
        c.reputationMax = 0;
        c.valid = false;

        emit CertPublishRequestCreated(c.domain);
    }

    function calculateCertValidity(string memory domain) private {
        CertMetadata storage c = certs[domain];
        
        if (!c.valid && c.reputation * reputationThreshold >= c.reputationMax) {
            c.valid = true;
            emit CertValid(domain);
        }
    }

    function certEndorseByPeer(string memory domain) public {
        Peer storage peer = peers[msg.sender];
        CertMetadata storage cert = certs[domain];

        require(cert.owner != msg.sender);
        require(!EnumerableSet.contains(cert.endorser, msg.sender));
        require(!EnumerableSet.contains(cert.denier, msg.sender));
        require(peer.p != address(0));

        // TODO: Update reputation calculation
        cert.reputation += peer.reputation;
        cert.reputationMax += 10;
        EnumerableSet.add(cert.endorser, msg.sender);

        calculateCertValidity(domain);

        emit CertEndorsed(domain, peer);
    }

    function certDenyByPeer(string memory domain) public {
        Peer storage peer = peers[msg.sender];
        CertMetadata storage cert = certs[domain];

        require(cert.owner != msg.sender);
        require(!EnumerableSet.contains(cert.endorser, msg.sender));
        require(!EnumerableSet.contains(cert.denier, msg.sender));
        require(peer.p != address(0));

        // TODO: Update reputation calculation
        cert.reputation -= peer.reputation;
        cert.reputationMax += 10;
        EnumerableSet.add(cert.denier, msg.sender);

        calculateCertValidity(domain);

        emit CertDenied(domain, peer);
    }

    function certEndorseByUser(string memory domain) public {
    }

    function certDenyByUser(string memory domain) public {
    }

    function certQuery(string memory domain) public view returns (string memory cert) {
        // TODO: Change data structure to certmetadata and certsupport?
        CertMetadata storage c = certs[domain];
        return c.cert;
    }

    function peerRegister(string memory pubkey) public {
        require(peers[msg.sender].p == address(0));
        
        Peer storage peer = peers[msg.sender];
        peer.p = msg.sender;
        peer.reputation = 0;
    }
}