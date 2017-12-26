package nas

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

func (client *Client) DeleteFileSystem(request *DeleteFileSystemRequest) (response *DeleteFileSystemResponse, err error) {
	response = CreateDeleteFileSystemResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteFileSystemWithChan(request *DeleteFileSystemRequest) (<-chan *DeleteFileSystemResponse, <-chan error) {
	responseChan := make(chan *DeleteFileSystemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFileSystem(request)
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

func (client *Client) DeleteFileSystemWithCallback(request *DeleteFileSystemRequest, callback func(response *DeleteFileSystemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFileSystemResponse
		var err error
		defer close(result)
		response, err = client.DeleteFileSystem(request)
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

type DeleteFileSystemRequest struct {
	*requests.RpcRequest
	FileSystemId string `position:"Query" name:"FileSystemId"`
}

type DeleteFileSystemResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteFileSystemRequest() (request *DeleteFileSystemRequest) {
	request = &DeleteFileSystemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "DeleteFileSystem", "", "")
	return
}

func CreateDeleteFileSystemResponse() (response *DeleteFileSystemResponse) {
	response = &DeleteFileSystemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
