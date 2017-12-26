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

func (client *Client) ListMoments(request *ListMomentsRequest) (response *ListMomentsResponse, err error) {
	response = CreateListMomentsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListMomentsWithChan(request *ListMomentsRequest) (<-chan *ListMomentsResponse, <-chan error) {
	responseChan := make(chan *ListMomentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMoments(request)
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

func (client *Client) ListMomentsWithCallback(request *ListMomentsRequest, callback func(response *ListMomentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMomentsResponse
		var err error
		defer close(result)
		response, err = client.ListMoments(request)
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

type ListMomentsRequest struct {
	*requests.RpcRequest
	Cursor    string `position:"Query" name:"Cursor"`
	Direction string `position:"Query" name:"Direction"`
	State     string `position:"Query" name:"State"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	Size      string `position:"Query" name:"Size"`
}

type ListMomentsResponse struct {
	*responses.BaseResponse
	Code       string          `json:"Code" xml:"Code"`
	Message    string          `json:"Message" xml:"Message"`
	NextCursor string          `json:"NextCursor" xml:"NextCursor"`
	TotalCount request.Integer `json:"TotalCount" xml:"TotalCount"`
	RequestId  string          `json:"RequestId" xml:"RequestId"`
	Action     string          `json:"Action" xml:"Action"`
	Moments    []struct {
		Id           request.Integer `json:"Id" xml:"Id"`
		LocationName string          `json:"LocationName" xml:"LocationName"`
		PhotosCount  request.Integer `json:"PhotosCount" xml:"PhotosCount"`
		State        string          `json:"State" xml:"State"`
		TakenAt      request.Integer `json:"TakenAt" xml:"TakenAt"`
		Ctime        request.Integer `json:"Ctime" xml:"Ctime"`
		Mtime        request.Integer `json:"Mtime" xml:"Mtime"`
	} `json:"Moments" xml:"Moments"`
}

func CreateListMomentsRequest() (request *ListMomentsRequest) {
	request = &ListMomentsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListMoments", "", "")
	return
}

func CreateListMomentsResponse() (response *ListMomentsResponse) {
	response = &ListMomentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
