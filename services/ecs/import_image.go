package ecs

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

func (client *Client) ImportImage(request *ImportImageRequest) (response *ImportImageResponse, err error) {
	response = CreateImportImageResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ImportImageWithChan(request *ImportImageRequest) (<-chan *ImportImageResponse, <-chan error) {
	responseChan := make(chan *ImportImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportImage(request)
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

func (client *Client) ImportImageWithCallback(request *ImportImageRequest, callback func(response *ImportImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportImageResponse
		var err error
		defer close(result)
		response, err = client.ImportImage(request)
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

type ImportImageRequest struct {
	*requests.RpcRequest
	Platform             string                          `position:"Query" name:"Platform"`
	RoleName             string                          `position:"Query" name:"RoleName"`
	ResourceOwnerAccount string                          `position:"Query" name:"ResourceOwnerAccount"`
	OSType               string                          `position:"Query" name:"OSType"`
	Description          string                          `position:"Query" name:"Description"`
	DiskDeviceMapping    *[]ImportImageDiskDeviceMapping `position:"Query" name:"DiskDeviceMapping"  type:"Repeated"`
	Architecture         string                          `position:"Query" name:"Architecture"`
	ResourceOwnerId      string                          `position:"Query" name:"ResourceOwnerId"`
	OwnerId              string                          `position:"Query" name:"OwnerId"`
	ImageName            string                          `position:"Query" name:"ImageName"`
}

type ImportImageDiskDeviceMapping struct {
	Format        string `name:"Format"`
	OSSBucket     string `name:"OSSBucket"`
	OSSObject     string `name:"OSSObject"`
	DiskImSize    string `name:"DiskImSize"`
	DiskImageSize string `name:"DiskImageSize"`
	Device        string `name:"Device"`
}

type ImportImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
	RegionId  string `json:"RegionId" xml:"RegionId"`
	ImageId   string `json:"ImageId" xml:"ImageId"`
}

func CreateImportImageRequest() (request *ImportImageRequest) {
	request = &ImportImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ImportImage", "", "")
	return
}

func CreateImportImageResponse() (response *ImportImageResponse) {
	response = &ImportImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
