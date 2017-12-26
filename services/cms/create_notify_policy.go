package cms

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

func (client *Client) CreateNotifyPolicy(request *CreateNotifyPolicyRequest) (response *CreateNotifyPolicyResponse, err error) {
	response = CreateCreateNotifyPolicyResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateNotifyPolicyWithChan(request *CreateNotifyPolicyRequest) (<-chan *CreateNotifyPolicyResponse, <-chan error) {
	responseChan := make(chan *CreateNotifyPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNotifyPolicy(request)
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

func (client *Client) CreateNotifyPolicyWithCallback(request *CreateNotifyPolicyRequest, callback func(response *CreateNotifyPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNotifyPolicyResponse
		var err error
		defer close(result)
		response, err = client.CreateNotifyPolicy(request)
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

type CreateNotifyPolicyRequest struct {
	*requests.RpcRequest
	EndTime    string `position:"Query" name:"EndTime"`
	AlertName  string `position:"Query" name:"AlertName"`
	StartTime  string `position:"Query" name:"StartTime"`
	PolicyType string `position:"Query" name:"PolicyType"`
	Dimensions string `position:"Query" name:"Dimensions"`
}

type CreateNotifyPolicyResponse struct {
	*responses.BaseResponse
	Code    string          `json:"code" xml:"code"`
	Message string          `json:"message" xml:"message"`
	Success string          `json:"success" xml:"success"`
	TraceId string          `json:"traceId" xml:"traceId"`
	Result  request.Integer `json:"result" xml:"result"`
}

func CreateCreateNotifyPolicyRequest() (request *CreateNotifyPolicyRequest) {
	request = &CreateNotifyPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "CreateNotifyPolicy", "", "")
	return
}

func CreateCreateNotifyPolicyResponse() (response *CreateNotifyPolicyResponse) {
	response = &CreateNotifyPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
