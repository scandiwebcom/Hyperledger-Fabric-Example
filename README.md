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

## Download Platform-specific Binaries
- curl -sSL https://goo.gl/6wtTN5 | bash -s 1.1.0-alpha

The command above downloads and executes a bash script that will download and extract all of the platform-specific binaries you will need to set up your network and place them into the cloned repo you created above. It retrieves four platform-specific binaries and places them in the bin sub-directory of the current working directory.

You may want to add that to your PATH environment variable so that these can be picked up without fully qualifying the path to each binary. e.g.:
- export PATH=<path to download location>/bin:$PATH

## Stop and remove all docker containers and images
- docker rm -f $(docker ps -aq)
- docker images
- docker rmi imageId

## ERRORS
- Error starting Simple chaincode: error sending chaincode REGISTER
!!!: CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=mycc:0  ./chaincode_example02