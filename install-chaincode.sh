#!/bin/bash

export CHANNEL_NAME=blockchain-camp
export CHAINCODE_PATH=${PWD}/studentcc/ # update it based on your chaincode
export CURRENT_PATH=${PWD}
export CHAINCODE_NAME=codecc

fabric-samples/test-network/network.sh down

fabric-samples/test-network/network.sh up createChannel -c $CHANNEL_NAME


echo "navigatting to chaincode..."
sleep 3

cd $CHAINCODE_PATH
GO111MODULE=on go mod vendor

echo "vendor created successfully..."
sleep 2

cd $CURRENT_PATH/fabric-samples/test-network

export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
echo "loading peer CLI tool"


peer lifecycle chaincode package $CHAINCODE_NAME.tar.gz --path $CHAINCODE_PATH --lang golang --label "${CHAINCODE_NAME}_1.0"

echo "packaging chaincode $CHAINCODE_NAME.tar.gz"

org1() {
    export CORE_PEER_TLS_ENABLED=true
    export CORE_PEER_LOCALMSPID="Org1MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE="${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt"
    export CORE_PEER_MSPCONFIGPATH="${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp"
    export CORE_PEER_ADDRESS="localhost:7051"
}

org2() {
    export CORE_PEER_TLS_ENABLED=true
    export CORE_PEER_LOCALMSPID=Org2MSP
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
    export CORE_PEER_ADDRESS=localhost:9051
}

echo "loading Org1 info"
org1
sleep 3

peer lifecycle chaincode install $CHAINCODE_NAME.tar.gz


echo "loading Org2 info"
org2
sleep 3

peer lifecycle chaincode install $CHAINCODE_NAME.tar.gz


echo "Approving a chaincode"
sleep 3

export CC_PACKAGE_ID=$(peer lifecycle chaincode queryinstalled | sed -n 's/.*Package ID: \([^,]*\).*/\1/p')

echo "package_id is ${CC_PACKAGE_ID}"
sleep 2

echo "approving for org 1"
sleep 1
org1
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"


echo "approving for org 2"
sleep 1
org2
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"



echo "Committing the chaincode"
sleep 3

peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --output json


peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"


peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name $CHAINCODE_NAME
