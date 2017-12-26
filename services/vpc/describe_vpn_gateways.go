package vpc

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

func (client *Client) DescribeVpnGateways(request *DescribeVpnGatewaysRequest) (response *DescribeVpnGatewaysResponse, err error) {
	response = CreateDescribeVpnGatewaysResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeVpnGatewaysWithChan(request *DescribeVpnGatewaysRequest) (<-chan *DescribeVpnGatewaysResponse, <-chan error) {
	responseChan := make(chan *DescribeVpnGatewaysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVpnGateways(request)
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

func (client *Client) DescribeVpnGatewaysWithCallback(request *DescribeVpnGatewaysRequest, callback func(response *DescribeVpnGatewaysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVpnGatewaysResponse
		var err error
		defer close(result)
		response, err = client.DescribeVpnGateways(request)
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

type DescribeVpnGatewaysRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	Status               string `position:"Query" name:"Status"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	BusinessStatus       string `position:"Query" name:"BusinessStatus"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
}

type DescribeVpnGatewaysResponse struct {
	*responses.BaseResponse
	RequestId   string          `json:"RequestId" xml:"RequestId"`
	TotalCount  request.Integer `json:"TotalCount" xml:"TotalCount"`
	PageNumber  request.Integer `json:"PageNumber" xml:"PageNumber"`
	PageSize    request.Integer `json:"PageSize" xml:"PageSize"`
	VpnGateways struct {
		VpnGateway []struct {
			VpnGatewayId   string          `json:"VpnGatewayId" xml:"VpnGatewayId"`
			VpcId          string          `json:"VpcId" xml:"VpcId"`
			VSwitchId      string          `json:"VSwitchId" xml:"VSwitchId"`
			InternetIp     string          `json:"InternetIp" xml:"InternetIp"`
			CreateTime     request.Integer `json:"CreateTime" xml:"CreateTime"`
			EndTime        request.Integer `json:"EndTime" xml:"EndTime"`
			Spec           string          `json:"Spec" xml:"Spec"`
			Name           string          `json:"Name" xml:"Name"`
			Description    string          `json:"Description" xml:"Description"`
			Status         string          `json:"Status" xml:"Status"`
			BusinessStatus string          `json:"BusinessStatus" xml:"BusinessStatus"`
			ChargeType     string          `json:"ChargeType" xml:"ChargeType"`
		} `json:"VpnGateway" xml:"VpnGateway"`
	} `json:"VpnGateways" xml:"VpnGateways"`
}

func CreateDescribeVpnGatewaysRequest() (request *DescribeVpnGatewaysRequest) {
	request = &DescribeVpnGatewaysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpnGateways", "", "")
	return
}

func CreateDescribeVpnGatewaysResponse() (response *DescribeVpnGatewaysResponse) {
	response = &DescribeVpnGatewaysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
