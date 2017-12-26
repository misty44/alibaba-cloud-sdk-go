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

func (client *Client) DescribeDeploymentSets(request *DescribeDeploymentSetsRequest) (response *DescribeDeploymentSetsResponse, err error) {
	response = CreateDescribeDeploymentSetsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDeploymentSetsWithChan(request *DescribeDeploymentSetsRequest) (<-chan *DescribeDeploymentSetsResponse, <-chan error) {
	responseChan := make(chan *DescribeDeploymentSetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDeploymentSets(request)
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

func (client *Client) DescribeDeploymentSetsWithCallback(request *DescribeDeploymentSetsRequest, callback func(response *DescribeDeploymentSetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDeploymentSetsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDeploymentSets(request)
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

type DescribeDeploymentSetsRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	DeploymentSetIds     string `position:"Query" name:"DeploymentSetIds"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	Strategy             string `position:"Query" name:"Strategy"`
	DeploymentSetName    string `position:"Query" name:"DeploymentSetName"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Granularity          string `position:"Query" name:"Granularity"`
	Domain               string `position:"Query" name:"Domain"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
}

type DescribeDeploymentSetsResponse struct {
	*responses.BaseResponse
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	RegionId       string          `json:"RegionId" xml:"RegionId"`
	TotalCount     request.Integer `json:"TotalCount" xml:"TotalCount"`
	PageNumber     request.Integer `json:"PageNumber" xml:"PageNumber"`
	PageSize       request.Integer `json:"PageSize" xml:"PageSize"`
	DeploymentSets struct {
		DeploymentSet []struct {
			DeploymentSetId          string          `json:"DeploymentSetId" xml:"DeploymentSetId"`
			DeploymentSetDescription string          `json:"DeploymentSetDescription" xml:"DeploymentSetDescription"`
			DeploymentSetName        string          `json:"DeploymentSetName" xml:"DeploymentSetName"`
			Strategy                 string          `json:"Strategy" xml:"Strategy"`
			Domain                   string          `json:"Domain" xml:"Domain"`
			Granularity              string          `json:"Granularity" xml:"Granularity"`
			InstanceAmount           request.Integer `json:"InstanceAmount" xml:"InstanceAmount"`
			CreationTime             string          `json:"CreationTime" xml:"CreationTime"`
		} `json:"DeploymentSet" xml:"DeploymentSet"`
	} `json:"DeploymentSets" xml:"DeploymentSets"`
}

func CreateDescribeDeploymentSetsRequest() (request *DescribeDeploymentSetsRequest) {
	request = &DescribeDeploymentSetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDeploymentSets", "", "")
	return
}

func CreateDescribeDeploymentSetsResponse() (response *DescribeDeploymentSetsResponse) {
	response = &DescribeDeploymentSetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
