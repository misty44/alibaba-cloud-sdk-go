package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) StartDBInstanceDiagnose(request *StartDBInstanceDiagnoseRequest) (response *StartDBInstanceDiagnoseResponse, err error) {
	response = CreateStartDBInstanceDiagnoseResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) StartDBInstanceDiagnoseWithChan(request *StartDBInstanceDiagnoseRequest) (<-chan *StartDBInstanceDiagnoseResponse, <-chan error) {
	responseChan := make(chan *StartDBInstanceDiagnoseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartDBInstanceDiagnose(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) StartDBInstanceDiagnoseWithCallback(request *StartDBInstanceDiagnoseRequest, callback func(response *StartDBInstanceDiagnoseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartDBInstanceDiagnoseResponse
		var err error
		defer close(result)
		response, err = client.StartDBInstanceDiagnose(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type StartDBInstanceDiagnoseRequest struct {
	*requests.RpcRequest
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	ProxyId              string `position:"Query" name:"proxyId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type StartDBInstanceDiagnoseResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
	DBInstanceId   string `json:"DBInstanceId" xml:"DBInstanceId"`
}

func CreateStartDBInstanceDiagnoseRequest() (request *StartDBInstanceDiagnoseRequest) {
	request = &StartDBInstanceDiagnoseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "StartDBInstanceDiagnose", "", "")
	return
}

func CreateStartDBInstanceDiagnoseResponse() (response *StartDBInstanceDiagnoseResponse) {
	response = &StartDBInstanceDiagnoseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
