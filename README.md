

# Hyperledger Fabric Asset Transfer Using Go

This project demonstrates how to use *Go chaincode* with *Hyperledger Fabric Test Network* to manage a simple asset transfer ledger.

## Prerequisites

Make sure the following are installed:

- Git
- cURL
- Docker & Docker Compose
- Go (1.17+)
- Node.js (for REST API, optional)
- Hyperledger Fabric binaries (v2.5 or compatible)

## Project Structure

fabric-samples/ ├── test-network/                     # Fabric test network ├── asset-transfer-basic/ │   ├── chaincode-go/                # Go smart contract │   ├── application-javascript/      # Optional: REST API

## Step 1: Download Fabric Samples and Binaries

```bash
curl -sSL https://bit.ly/HyperledgerFabricInstaller | bash -s

This will create fabric-samples/ with binaries and config.

Step 2: Start the Test Network

cd fabric-samples/test-network
./network.sh up createChannel -c mychannel -ca

Step 3: Deploy the Go Chaincode

./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go -ccl go

Step 4: Interact with Chaincode

Set environment variables for Org1:

export PATH=$PATH:../bin
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=$PWD/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=$PWD/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

Invoke InitLedger and query assets:

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com \
--tls --cafile "$PWD/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" \
-C mychannel -n basic -c '{"function":"InitLedger","Args":[]}'

peer chaincode query -C mychannel -n basic -c '{"Args":["GetAllAssets"]}'

Optional: Build REST API (JavaScript)

Navigate to:

cd ../asset-transfer-basic/application-javascript
npm install
node app.js

