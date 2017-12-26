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

func (client *Client) DeleteSkillGroup(request *DeleteSkillGroupRequest) (response *DeleteSkillGroupResponse, err error) {
	response = CreateDeleteSkillGroupResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteSkillGroupWithChan(request *DeleteSkillGroupRequest) (<-chan *DeleteSkillGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteSkillGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSkillGroup(request)
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

func (client *Client) DeleteSkillGroupWithCallback(request *DeleteSkillGroupRequest, callback func(response *DeleteSkillGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSkillGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteSkillGroup(request)
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

type DeleteSkillGroupRequest struct {
	*requests.RpcRequest
	SkillGroupId string `position:"Query" name:"SkillGroupId"`
	InstanceId   string `position:"Query" name:"InstanceId"`
}

type DeleteSkillGroupResponse struct {
	*responses.BaseResponse
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	Success        request.Boolean `json:"Success" xml:"Success"`
	Code           string          `json:"Code" xml:"Code"`
	Message        string          `json:"Message" xml:"Message"`
	HttpStatusCode request.Integer `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

func CreateDeleteSkillGroupRequest() (request *DeleteSkillGroupRequest) {
	request = &DeleteSkillGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "DeleteSkillGroup", "", "")
	return
}

func CreateDeleteSkillGroupResponse() (response *DeleteSkillGroupResponse) {
	response = &DeleteSkillGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
