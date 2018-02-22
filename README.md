# Notes

1.  Hyperledger Fabric runs chaincode which is comparable to Ethereum's smart contracts. The consensus protocol is pluggable.

2.  https://s3-eu-west-1.amazonaws.com/b9-academy-assets/course-assets/HLF-FREE-0/table-hlf-traits.png

3. Nodes:
	- Client - Clients are the end-user facing nodes. Hyperledger Fabric architecture provides multiple interfaces to the blockchain.
	- Peer - Peers maintain the state of the ledger. Peers execute chaincode and participate in consensus formation. Chaincode is installed on peers.
	- Orderer - The OS creates new blocks by ordering the transactions. It provides a shared communication channel to clients and peers. 

4. In Hyperledger Fabric, the ledger is maintained and validated by the Peers.

5. Channels - Hyperledger Fabric introduces channels. You can think of a channel as a kind of separate blockchain.

6. Shim - The package we will use to write chaincode is the package shim. It provides the type Chaincode interface, which is an interface.



## ERRORS
- Error starting Simple chaincode: error sending chaincode REGISTER
!!!: CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=mycc:0  ./chaincode_example02