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

func (client *Client) ProfileSet(request *ProfileSetRequest) (response *ProfileSetResponse, err error) {
	response = CreateProfileSetResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ProfileSetWithChan(request *ProfileSetRequest) (<-chan *ProfileSetResponse, <-chan error) {
	responseChan := make(chan *ProfileSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ProfileSet(request)
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

func (client *Client) ProfileSetWithCallback(request *ProfileSetRequest, callback func(response *ProfileSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ProfileSetResponse
		var err error
		defer close(result)
		response, err = client.ProfileSet(request)
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

type ProfileSetRequest struct {
	*requests.RpcRequest
	AutoInstall              string `position:"Query" name:"AutoInstall"`
	UserId                   string `position:"Query" name:"UserId"`
	EnableInstallAgentNewECS string `position:"Query" name:"EnableInstallAgentNewECS"`
}

type ProfileSetResponse struct {
	*responses.BaseResponse
	ErrorCode    request.Integer `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string          `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      request.Boolean `json:"Success" xml:"Success"`
	RequestId    string          `json:"RequestId" xml:"RequestId"`
}

func CreateProfileSetRequest() (request *ProfileSetRequest) {
	request = &ProfileSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "ProfileSet", "", "")
	return
}

func CreateProfileSetResponse() (response *ProfileSetResponse) {
	response = &ProfileSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
