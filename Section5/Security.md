Security
=======

# 1. Security model

The MCGP protocol's security model assumes that MCGP protocol will be used primarily over a local, trusted network. Second assumption is that the data carried by MCGP protocol (such as temperature) is not sensitive. Therefore, MCGP focuses on authentication rather than encryption.

Much of MCGP's security depends, therefore, on the security of physical network. While on the public Internet, the threat of a malicious intermediary intercepting or modifying the information on-the-fly is high, such attacks are not expected in (W)LANs. A properly secured local home network isolate outsiders and any party able to listen in on the unicast traffic on the network is therefore probably the network owner.

MCGP specifically assumes that:
- TCP connections cannot be hijacked
- TCP/IP packets cannot be rerouted (or ARP-spoofed) to a malicious server
- Unicast packets cannot be intercepted

# 2. Authentication method

Despite the naive assumptions, MCGP takes mesaures to validate incoming client connections. At this moment, only servers can authenticate clients. The clients have to implicitly trust that the serveer they connect to is not an imposter (see security asumptions above).

## Message hashing & signing

For the purpose of identification and authentication, every packet is signed by the client. The PDU includes a field (last field) that carries a secure signature of a hash of the message. The message signature can be obtained in one of two supported methods:

### Shared-secret signature
The hash-signature is obtained by concatenating all fields of the PDU (in the order they are defined and sent) together, appending the byte-representation of the shared-secret (key) and computing an sha256 sum. The sum is then transmitted as the last field of the PDU.

### Public-key signature
The hash-signature is obtained by concatenating all fields of the PDU (in the order they are defined and sent) together and computing an sha256 sum. The sum is then signed using the RSA 2048-bit private-key. The signature is used as the last field of the PDU.

## Message authentication

The server can authenticate a message by computing a hash of the packet received from a client using same rules as defined above. In case of shared-secret signature, the server knows the shared secret and can apply it. In case of the public-key signature, the server knows the public key of the client and can verify the signature.

## Connection authentication

The client can authenticate a connect in its early stage by sending its ID (e.g. a hostname or username) as the authentication payload and signing the packet properly. The server can then determine if the signature matches the claimed name.

# 3. Security risks / vulneraiblities

* Server spoofing / MITM attacks - the protoocl does NOT provide any resiliency against MITM attacks.

* Replay attacks - the protocol defends against replay attacks by sending a new randomly-generated one-time-challenge from the server to the client, to be included (and signed) in the next packet. The security of this solution depends on the randomness of the server.

* Packet tampering - the protocol's message signing method provides sufficient protection against any message tampering.
