package alimt

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

// GetDocTranslateTask invokes the alimt.GetDocTranslateTask API synchronously
// api document: https://help.aliyun.com/api/alimt/getdoctranslatetask.html
func (client *Client) GetDocTranslateTask(request *GetDocTranslateTaskRequest) (response *GetDocTranslateTaskResponse, err error) {
	response = CreateGetDocTranslateTaskResponse()
	err = client.DoAction(request, response)
	return
}

// GetDocTranslateTaskWithChan invokes the alimt.GetDocTranslateTask API asynchronously
// api document: https://help.aliyun.com/api/alimt/getdoctranslatetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDocTranslateTaskWithChan(request *GetDocTranslateTaskRequest) (<-chan *GetDocTranslateTaskResponse, <-chan error) {
	responseChan := make(chan *GetDocTranslateTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDocTranslateTask(request)
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

// GetDocTranslateTaskWithCallback invokes the alimt.GetDocTranslateTask API asynchronously
// api document: https://help.aliyun.com/api/alimt/getdoctranslatetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDocTranslateTaskWithCallback(request *GetDocTranslateTaskRequest, callback func(response *GetDocTranslateTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDocTranslateTaskResponse
		var err error
		defer close(result)
		response, err = client.GetDocTranslateTask(request)
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

// GetDocTranslateTaskRequest is the request struct for api GetDocTranslateTask
type GetDocTranslateTaskRequest struct {
	*requests.RpcRequest
	TaskId string `position:"Query" name:"TaskId"`
}

// GetDocTranslateTaskResponse is the response struct for api GetDocTranslateTask
type GetDocTranslateTaskResponse struct {
	*responses.BaseResponse
	RequestId             string `json:"RequestId" xml:"RequestId"`
	TaskId                string `json:"TaskId" xml:"TaskId"`
	Status                string `json:"Status" xml:"Status"`
	TranslateFileUrl      string `json:"TranslateFileUrl" xml:"TranslateFileUrl"`
	TranslateErrorCode    string `json:"TranslateErrorCode" xml:"TranslateErrorCode"`
	TranslateErrorMessage string `json:"TranslateErrorMessage" xml:"TranslateErrorMessage"`
	PageCount             int    `json:"PageCount" xml:"PageCount"`
}

// CreateGetDocTranslateTaskRequest creates a request to invoke GetDocTranslateTask API
func CreateGetDocTranslateTaskRequest() (request *GetDocTranslateTaskRequest) {
	request = &GetDocTranslateTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alimt", "2018-10-12", "GetDocTranslateTask", "alimt", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetDocTranslateTaskResponse creates a response to parse from GetDocTranslateTask response
func CreateGetDocTranslateTaskResponse() (response *GetDocTranslateTaskResponse) {
	response = &GetDocTranslateTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
