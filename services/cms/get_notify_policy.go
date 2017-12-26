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

func (client *Client) GetNotifyPolicy(request *GetNotifyPolicyRequest) (response *GetNotifyPolicyResponse, err error) {
	response = CreateGetNotifyPolicyResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetNotifyPolicyWithChan(request *GetNotifyPolicyRequest) (<-chan *GetNotifyPolicyResponse, <-chan error) {
	responseChan := make(chan *GetNotifyPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNotifyPolicy(request)
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

func (client *Client) GetNotifyPolicyWithCallback(request *GetNotifyPolicyRequest, callback func(response *GetNotifyPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNotifyPolicyResponse
		var err error
		defer close(result)
		response, err = client.GetNotifyPolicy(request)
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

type GetNotifyPolicyRequest struct {
	*requests.RpcRequest
	Id         string `position:"Query" name:"Id"`
	AlertName  string `position:"Query" name:"AlertName"`
	PolicyType string `position:"Query" name:"PolicyType"`
	Dimensions string `position:"Query" name:"Dimensions"`
}

type GetNotifyPolicyResponse struct {
	*responses.BaseResponse
	Code    string `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
	Success string `json:"success" xml:"success"`
	TraceId string `json:"traceId" xml:"traceId"`
	Result  struct {
		AlertName  string          `json:"AlertName" xml:"AlertName"`
		Dimensions string          `json:"Dimensions" xml:"Dimensions"`
		Type       string          `json:"Type" xml:"Type"`
		Id         string          `json:"Id" xml:"Id"`
		StartTime  request.Integer `json:"StartTime" xml:"StartTime"`
		EndTime    request.Integer `json:"EndTime" xml:"EndTime"`
	} `json:"Result" xml:"Result"`
}

func CreateGetNotifyPolicyRequest() (request *GetNotifyPolicyRequest) {
	request = &GetNotifyPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "GetNotifyPolicy", "", "")
	return
}

func CreateGetNotifyPolicyResponse() (response *GetNotifyPolicyResponse) {
	response = &GetNotifyPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
