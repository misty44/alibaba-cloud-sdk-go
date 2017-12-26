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

func (client *Client) DescribePhysicalConnections(request *DescribePhysicalConnectionsRequest) (response *DescribePhysicalConnectionsResponse, err error) {
	response = CreateDescribePhysicalConnectionsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribePhysicalConnectionsWithChan(request *DescribePhysicalConnectionsRequest) (<-chan *DescribePhysicalConnectionsResponse, <-chan error) {
	responseChan := make(chan *DescribePhysicalConnectionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePhysicalConnections(request)
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

func (client *Client) DescribePhysicalConnectionsWithCallback(request *DescribePhysicalConnectionsRequest, callback func(response *DescribePhysicalConnectionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePhysicalConnectionsResponse
		var err error
		defer close(result)
		response, err = client.DescribePhysicalConnections(request)
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

type DescribePhysicalConnectionsRequest struct {
	*requests.RpcRequest
	PageSize             string                               `position:"Query" name:"PageSize"`
	ClientToken          string                               `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string                               `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string                               `position:"Query" name:"PageNumber"`
	UserCidr             string                               `position:"Query" name:"UserCidr"`
	ResourceOwnerId      string                               `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string                               `position:"Query" name:"OwnerAccount"`
	OwnerId              string                               `position:"Query" name:"OwnerId"`
	Filter               *[]DescribePhysicalConnectionsFilter `position:"Query" name:"Filter"  type:"Repeated"`
}

type DescribePhysicalConnectionsFilter struct {
	Key   string    `name:"Key"`
	Value *[]string `name:"Value" type:"Repeated"`
}

type DescribePhysicalConnectionsResponse struct {
	*responses.BaseResponse
	RequestId             string          `json:"RequestId" xml:"RequestId"`
	PageNumber            request.Integer `json:"PageNumber" xml:"PageNumber"`
	PageSize              request.Integer `json:"PageSize" xml:"PageSize"`
	TotalCount            request.Integer `json:"TotalCount" xml:"TotalCount"`
	PhysicalConnectionSet struct {
		PhysicalConnectionType []struct {
			PhysicalConnectionId          string          `json:"PhysicalConnectionId" xml:"PhysicalConnectionId"`
			AccessPointId                 string          `json:"AccessPointId" xml:"AccessPointId"`
			Type                          string          `json:"Type" xml:"Type"`
			Status                        string          `json:"Status" xml:"Status"`
			BusinessStatus                string          `json:"BusinessStatus" xml:"BusinessStatus"`
			CreationTime                  string          `json:"CreationTime" xml:"CreationTime"`
			EnabledTime                   string          `json:"EnabledTime" xml:"EnabledTime"`
			LineOperator                  string          `json:"LineOperator" xml:"LineOperator"`
			Spec                          string          `json:"Spec" xml:"Spec"`
			PeerLocation                  string          `json:"PeerLocation" xml:"PeerLocation"`
			PortType                      string          `json:"PortType" xml:"PortType"`
			RedundantPhysicalConnectionId string          `json:"RedundantPhysicalConnectionId" xml:"RedundantPhysicalConnectionId"`
			Name                          string          `json:"Name" xml:"Name"`
			Description                   string          `json:"Description" xml:"Description"`
			AdLocation                    string          `json:"AdLocation" xml:"AdLocation"`
			PortNumber                    string          `json:"PortNumber" xml:"PortNumber"`
			CircuitCode                   string          `json:"CircuitCode" xml:"CircuitCode"`
			Bandwidth                     request.Integer `json:"Bandwidth" xml:"Bandwidth"`
		} `json:"PhysicalConnectionType" xml:"PhysicalConnectionType"`
	} `json:"PhysicalConnectionSet" xml:"PhysicalConnectionSet"`
}

func CreateDescribePhysicalConnectionsRequest() (request *DescribePhysicalConnectionsRequest) {
	request = &DescribePhysicalConnectionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribePhysicalConnections", "", "")
	return
}

func CreateDescribePhysicalConnectionsResponse() (response *DescribePhysicalConnectionsResponse) {
	response = &DescribePhysicalConnectionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
