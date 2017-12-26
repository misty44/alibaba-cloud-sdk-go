package ccc

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

func (client *Client) RequestLoginInfo(request *RequestLoginInfoRequest) (response *RequestLoginInfoResponse, err error) {
	response = CreateRequestLoginInfoResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) RequestLoginInfoWithChan(request *RequestLoginInfoRequest) (<-chan *RequestLoginInfoResponse, <-chan error) {
	responseChan := make(chan *RequestLoginInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RequestLoginInfo(request)
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

func (client *Client) RequestLoginInfoWithCallback(request *RequestLoginInfoRequest, callback func(response *RequestLoginInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RequestLoginInfoResponse
		var err error
		defer close(result)
		response, err = client.RequestLoginInfo(request)
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

type RequestLoginInfoRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

type RequestLoginInfoResponse struct {
	*responses.BaseResponse
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	Success        request.Boolean `json:"Success" xml:"Success"`
	Code           string          `json:"Code" xml:"Code"`
	Message        string          `json:"Message" xml:"Message"`
	HttpStatusCode request.Integer `json:"HttpStatusCode" xml:"HttpStatusCode"`
	LoginInfo      struct {
		UserName       string `json:"UserName" xml:"UserName"`
		DisplayName    string `json:"DisplayName" xml:"DisplayName"`
		Region         string `json:"Region" xml:"Region"`
		WebRtcUrl      string `json:"WebRtcUrl" xml:"WebRtcUrl"`
		AgentServerUrl string `json:"AgentServerUrl" xml:"AgentServerUrl"`
		Extension      string `json:"Extension" xml:"Extension"`
		TenantId       string `json:"TenantId" xml:"TenantId"`
		Signature      string `json:"Signature" xml:"Signature"`
		SignData       string `json:"SignData" xml:"SignData"`
	} `json:"LoginInfo" xml:"LoginInfo"`
}

func CreateRequestLoginInfoRequest() (request *RequestLoginInfoRequest) {
	request = &RequestLoginInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "RequestLoginInfo", "", "")
	return
}

func CreateRequestLoginInfoResponse() (response *RequestLoginInfoResponse) {
	response = &RequestLoginInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
