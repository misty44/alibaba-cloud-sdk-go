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

func (client *Client) GetPhotos(request *GetPhotosRequest) (response *GetPhotosResponse, err error) {
	response = CreateGetPhotosResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetPhotosWithChan(request *GetPhotosRequest) (<-chan *GetPhotosResponse, <-chan error) {
	responseChan := make(chan *GetPhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPhotos(request)
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

func (client *Client) GetPhotosWithCallback(request *GetPhotosRequest, callback func(response *GetPhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPhotosResponse
		var err error
		defer close(result)
		response, err = client.GetPhotos(request)
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

type GetPhotosRequest struct {
	*requests.RpcRequest
	LibraryId string    `position:"Query" name:"LibraryId"`
	StoreName string    `position:"Query" name:"StoreName"`
	PhotoId   *[]string `position:"Query" name:"PhotoId"  type:"Repeated"`
}

type GetPhotosResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Photos    []struct {
		Id              request.Integer `json:"Id" xml:"Id"`
		Title           string          `json:"Title" xml:"Title"`
		FileId          string          `json:"FileId" xml:"FileId"`
		State           string          `json:"State" xml:"State"`
		Md5             string          `json:"Md5" xml:"Md5"`
		IsVideo         request.Boolean `json:"IsVideo" xml:"IsVideo"`
		Remark          string          `json:"Remark" xml:"Remark"`
		Width           request.Integer `json:"Width" xml:"Width"`
		Height          request.Integer `json:"Height" xml:"Height"`
		Ctime           request.Integer `json:"Ctime" xml:"Ctime"`
		Mtime           request.Integer `json:"Mtime" xml:"Mtime"`
		TakenAt         request.Integer `json:"TakenAt" xml:"TakenAt"`
		InactiveTime    request.Integer `json:"InactiveTime" xml:"InactiveTime"`
		ShareExpireTime request.Integer `json:"ShareExpireTime" xml:"ShareExpireTime"`
		Like            request.Integer `json:"Like" xml:"Like"`
	} `json:"Photos" xml:"Photos"`
}

func CreateGetPhotosRequest() (request *GetPhotosRequest) {
	request = &GetPhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetPhotos", "", "")
	return
}

func CreateGetPhotosResponse() (response *GetPhotosResponse) {
	response = &GetPhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
