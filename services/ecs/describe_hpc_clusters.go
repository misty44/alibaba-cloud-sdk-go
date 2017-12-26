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

func (client *Client) DescribeHpcClusters(request *DescribeHpcClustersRequest) (response *DescribeHpcClustersResponse, err error) {
	response = CreateDescribeHpcClustersResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeHpcClustersWithChan(request *DescribeHpcClustersRequest) (<-chan *DescribeHpcClustersResponse, <-chan error) {
	responseChan := make(chan *DescribeHpcClustersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHpcClusters(request)
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

func (client *Client) DescribeHpcClustersWithCallback(request *DescribeHpcClustersRequest, callback func(response *DescribeHpcClustersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHpcClustersResponse
		var err error
		defer close(result)
		response, err = client.DescribeHpcClusters(request)
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

type DescribeHpcClustersRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	HpcClusterIds        string `position:"Query" name:"HpcClusterIds"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DescribeHpcClustersResponse struct {
	*responses.BaseResponse
	RequestId   string          `json:"RequestId" xml:"RequestId"`
	TotalCount  request.Integer `json:"TotalCount" xml:"TotalCount"`
	PageNumber  request.Integer `json:"PageNumber" xml:"PageNumber"`
	PageSize    request.Integer `json:"PageSize" xml:"PageSize"`
	HpcClusters struct {
		HpcCluster []struct {
			HpcClusterId string `json:"HpcClusterId" xml:"HpcClusterId"`
			Name         string `json:"Name" xml:"Name"`
			Description  string `json:"Description" xml:"Description"`
		} `json:"HpcCluster" xml:"HpcCluster"`
	} `json:"HpcClusters" xml:"HpcClusters"`
}

func CreateDescribeHpcClustersRequest() (request *DescribeHpcClustersRequest) {
	request = &DescribeHpcClustersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeHpcClusters", "", "")
	return
}

func CreateDescribeHpcClustersResponse() (response *DescribeHpcClustersResponse) {
	response = &DescribeHpcClustersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
