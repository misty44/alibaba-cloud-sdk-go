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

func (client *Client) DescribeIntranetAttributeKb(request *DescribeIntranetAttributeKbRequest) (response *DescribeIntranetAttributeKbResponse, err error) {
	response = CreateDescribeIntranetAttributeKbResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeIntranetAttributeKbWithChan(request *DescribeIntranetAttributeKbRequest) (<-chan *DescribeIntranetAttributeKbResponse, <-chan error) {
	responseChan := make(chan *DescribeIntranetAttributeKbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIntranetAttributeKb(request)
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

func (client *Client) DescribeIntranetAttributeKbWithCallback(request *DescribeIntranetAttributeKbRequest, callback func(response *DescribeIntranetAttributeKbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIntranetAttributeKbResponse
		var err error
		defer close(result)
		response, err = client.DescribeIntranetAttributeKb(request)
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

type DescribeIntranetAttributeKbRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
}

type DescribeIntranetAttributeKbResponse struct {
	*responses.BaseResponse
	RequestId               string          `json:"RequestId" xml:"RequestId"`
	InstanceId              string          `json:"InstanceId" xml:"InstanceId"`
	VlanId                  string          `json:"VlanId" xml:"VlanId"`
	IntranetIpAddress       string          `json:"IntranetIpAddress" xml:"IntranetIpAddress"`
	IntranetMaxBandwidthIn  request.Integer `json:"IntranetMaxBandwidthIn" xml:"IntranetMaxBandwidthIn"`
	IntranetMaxBandwidthOut request.Integer `json:"IntranetMaxBandwidthOut" xml:"IntranetMaxBandwidthOut"`
}

func CreateDescribeIntranetAttributeKbRequest() (request *DescribeIntranetAttributeKbRequest) {
	request = &DescribeIntranetAttributeKbRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeIntranetAttributeKb", "", "")
	return
}

func CreateDescribeIntranetAttributeKbResponse() (response *DescribeIntranetAttributeKbResponse) {
	response = &DescribeIntranetAttributeKbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
