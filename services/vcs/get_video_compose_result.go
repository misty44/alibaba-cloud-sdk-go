package vcs

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

// GetVideoComposeResult invokes the vcs.GetVideoComposeResult API synchronously
// api document: https://help.aliyun.com/api/vcs/getvideocomposeresult.html
func (client *Client) GetVideoComposeResult(request *GetVideoComposeResultRequest) (response *GetVideoComposeResultResponse, err error) {
	response = CreateGetVideoComposeResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetVideoComposeResultWithChan invokes the vcs.GetVideoComposeResult API asynchronously
// api document: https://help.aliyun.com/api/vcs/getvideocomposeresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetVideoComposeResultWithChan(request *GetVideoComposeResultRequest) (<-chan *GetVideoComposeResultResponse, <-chan error) {
	responseChan := make(chan *GetVideoComposeResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVideoComposeResult(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetVideoComposeResultWithCallback invokes the vcs.GetVideoComposeResult API asynchronously
// api document: https://help.aliyun.com/api/vcs/getvideocomposeresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetVideoComposeResultWithCallback(request *GetVideoComposeResultRequest, callback func(response *GetVideoComposeResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVideoComposeResultResponse
		var err error
		defer close(result)
		response, err = client.GetVideoComposeResult(request)
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

// GetVideoComposeResultRequest is the request struct for api GetVideoComposeResult
type GetVideoComposeResultRequest struct {
	*requests.RpcRequest
	CorpId        string `position:"Body" name:"CorpId"`
	TaskRequestId string `position:"Body" name:"TaskRequestId"`
}

// GetVideoComposeResultResponse is the response struct for api GetVideoComposeResult
type GetVideoComposeResultResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	VideoUrl  string `json:"VideoUrl" xml:"VideoUrl"`
	Code      string `json:"Code" xml:"Code"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateGetVideoComposeResultRequest creates a request to invoke GetVideoComposeResult API
func CreateGetVideoComposeResultRequest() (request *GetVideoComposeResultRequest) {
	request = &GetVideoComposeResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "GetVideoComposeResult", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetVideoComposeResultResponse creates a response to parse from GetVideoComposeResult response
func CreateGetVideoComposeResultResponse() (response *GetVideoComposeResultResponse) {
	response = &GetVideoComposeResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
