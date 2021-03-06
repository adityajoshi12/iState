/*
	Copyright 2020 Motoreq Infotech Pvt Ltd

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"testing"
)

func CreateState(test *testing.T, stub *shim.MockStub, input TestStruct, txID int) pb.Response {
	inputString, err := MarshalAndStringify(input)
	if err != nil {
		test.FailNow()
	}
	return Invoke(test, stub, txID, "CreateState", inputString)
}

func QueryState(test *testing.T, stub *shim.MockStub, input interface{}, txID int) pb.Response {
	inputString, err := MarshalAndStringify(input)
	if err != nil {
		test.FailNow()
	}
	return Invoke(test, stub, txID, "QueryState", inputString)
}
