package cloudphoto

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

func (client *Client) EditPhotos(request *EditPhotosRequest) (response *EditPhotosResponse, err error) {
	response = CreateEditPhotosResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) EditPhotosWithChan(request *EditPhotosRequest) (<-chan *EditPhotosResponse, <-chan error) {
	responseChan := make(chan *EditPhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EditPhotos(request)
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

func (client *Client) EditPhotosWithCallback(request *EditPhotosRequest, callback func(response *EditPhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EditPhotosResponse
		var err error
		defer close(result)
		response, err = client.EditPhotos(request)
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

type EditPhotosRequest struct {
	*requests.RpcRequest
	Title           string    `position:"Query" name:"Title"`
	ShareExpireTime string    `position:"Query" name:"ShareExpireTime"`
	Remark          string    `position:"Query" name:"Remark"`
	LibraryId       string    `position:"Query" name:"LibraryId"`
	StoreName       string    `position:"Query" name:"StoreName"`
	PhotoId         *[]string `position:"Query" name:"PhotoId"  type:"Repeated"`
}

type EditPhotosResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Results   []struct {
		Id      request.Integer `json:"Id" xml:"Id"`
		Code    string          `json:"Code" xml:"Code"`
		Message string          `json:"Message" xml:"Message"`
	} `json:"Results" xml:"Results"`
}

func CreateEditPhotosRequest() (request *EditPhotosRequest) {
	request = &EditPhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "EditPhotos", "", "")
	return
}

func CreateEditPhotosResponse() (response *EditPhotosResponse) {
	response = &EditPhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
