package mts

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

func (client *Client) SubmitAsrJob(request *SubmitAsrJobRequest) (response *SubmitAsrJobResponse, err error) {
	response = CreateSubmitAsrJobResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SubmitAsrJobWithChan(request *SubmitAsrJobRequest) (<-chan *SubmitAsrJobResponse, <-chan error) {
	responseChan := make(chan *SubmitAsrJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitAsrJob(request)
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

func (client *Client) SubmitAsrJobWithCallback(request *SubmitAsrJobRequest, callback func(response *SubmitAsrJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitAsrJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitAsrJob(request)
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

type SubmitAsrJobRequest struct {
	*requests.RpcRequest
	UserData             string `position:"Query" name:"UserData"`
	Input                string `position:"Query" name:"Input"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	AsrConfig            string `position:"Query" name:"AsrConfig"`
}

type SubmitAsrJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

func CreateSubmitAsrJobRequest() (request *SubmitAsrJobRequest) {
	request = &SubmitAsrJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitAsrJob", "", "")
	return
}

func CreateSubmitAsrJobResponse() (response *SubmitAsrJobResponse) {
	response = &SubmitAsrJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
