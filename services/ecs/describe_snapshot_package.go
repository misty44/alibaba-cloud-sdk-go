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

func (client *Client) DescribeSnapshotPackage(request *DescribeSnapshotPackageRequest) (response *DescribeSnapshotPackageResponse, err error) {
	response = CreateDescribeSnapshotPackageResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeSnapshotPackageWithChan(request *DescribeSnapshotPackageRequest) (<-chan *DescribeSnapshotPackageResponse, <-chan error) {
	responseChan := make(chan *DescribeSnapshotPackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSnapshotPackage(request)
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

func (client *Client) DescribeSnapshotPackageWithCallback(request *DescribeSnapshotPackageRequest, callback func(response *DescribeSnapshotPackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSnapshotPackageResponse
		var err error
		defer close(result)
		response, err = client.DescribeSnapshotPackage(request)
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

type DescribeSnapshotPackageRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DescribeSnapshotPackageResponse struct {
	*responses.BaseResponse
	RequestId        string          `json:"RequestId" xml:"RequestId"`
	TotalCount       request.Integer `json:"TotalCount" xml:"TotalCount"`
	PageNumber       request.Integer `json:"PageNumber" xml:"PageNumber"`
	PageSize         request.Integer `json:"PageSize" xml:"PageSize"`
	SnapshotPackages struct {
		SnapshotPackage []struct {
			StartTime    string          `json:"StartTime" xml:"StartTime"`
			EndTime      string          `json:"EndTime" xml:"EndTime"`
			InitCapacity request.Integer `json:"InitCapacity" xml:"InitCapacity"`
			DisplayName  string          `json:"DisplayName" xml:"DisplayName"`
		} `json:"SnapshotPackage" xml:"SnapshotPackage"`
	} `json:"SnapshotPackages" xml:"SnapshotPackages"`
}

func CreateDescribeSnapshotPackageRequest() (request *DescribeSnapshotPackageRequest) {
	request = &DescribeSnapshotPackageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSnapshotPackage", "", "")
	return
}

func CreateDescribeSnapshotPackageResponse() (response *DescribeSnapshotPackageResponse) {
	response = &DescribeSnapshotPackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
