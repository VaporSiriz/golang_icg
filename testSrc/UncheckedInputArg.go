package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type BadChaincode struct {
}

func (t BadChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	result , err := stub.GetState(args[0])
	b := result

	if err != nil {
		return shim.Error(err.Error())
	}

	return  shim.Success(b)
}


