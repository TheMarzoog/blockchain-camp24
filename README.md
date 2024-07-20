# blockchain-camp24



## Blockchain Overview
- [Bitcoin Whitepaper](https://bitcoin.org/bitcoin.pdf)
- [Blockchain Technology](https://chain.link/education-hub/blockchain)
- [Distributed System](https://www.atlassian.com/microservices/microservices-architecture/distributed-architecture)
- [Blockchain Demo](https://andersbrownworth.com/blockchain/hash)
- [Signature Demo](https://elearning.kba.ai/custom_app)

#### Hashing & Cryptograpy

```go
// Tx creatation
data = someData
hash = SHA256(data)
signature = Encrypt(hash, senderPrivateKey)

Tx = "data:hash:signature"

// Tx validation
validateHash = Decrypt(signature, senderPublicKey)
validateHash == Hash
```

## Hyperledger Fabric
- [Hyperledger](https://www.hyperledger.org/)
- [Hyperledger Fabric](https://hyperledger-fabric.readthedocs.io/en/release-2.5/)
- [Hyperledger Fabric Repository](https://github.com/hyperledger/fabric)
- [Fabric Samples](https://github.com/hyperledger/fabric-samples)
