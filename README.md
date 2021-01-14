[//]: # (SPDX-License-Identifier: CC-BY-4.0)

# Veritas Blockchain 

You can use Fabric samples to get started working with Hyperledger Fabric, explore important Fabric features, and learn how to build applications that can interact with blockchain networks using the Fabric SDKs. To learn more about Hyperledger Fabric, visit the [Fabric documentation](https://hyperledger-fabric.readthedocs.io/en/latest).

## Veritas network

### Clean up
Navigate to the test-network dir and run the following command.  This will bring down all the docker containers running the blockchain and clear any data we have.
```
cd /root/fabric-samples/test-network
./network.sh down
```


### Bringing up the blockchain
Make sure the following environment var is set
```
export FABRIC_CFG_PATH=/root/fabric-samples/test-network/../config/
```

To start a new blockchain run the following command. This brings up your CA, registers, and enrolls your peers and orderers creating the genesis blockchain.  Also it will create a channel.  At this point you have a blockchain but no smart contracts to execute on it.
```
./network.sh up createChannel -c mychannel -ca
```


### Deploy the smart contract (chaincode)
To deploy the chaincode run the following command.  At this point you should see a new docker container for each peer as this is your deployed chaincode.  You can now run your applications that can execute smartcontracts.
```
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go/ -ccl go
```


## Applications
Run the nodejs based application.  This will start the Application for interacting with the blockchain.  The api will run on port 3000.
```
cd /root/fabric-samples/asset-transfer-basic/application-javascript
node app.js
```

