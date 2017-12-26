package slb

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

func (client *Client) SetLoadBalancerHTTPSListenerAttribute(request *SetLoadBalancerHTTPSListenerAttributeRequest) (response *SetLoadBalancerHTTPSListenerAttributeResponse, err error) {
	response = CreateSetLoadBalancerHTTPSListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetLoadBalancerHTTPSListenerAttributeWithChan(request *SetLoadBalancerHTTPSListenerAttributeRequest) (<-chan *SetLoadBalancerHTTPSListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *SetLoadBalancerHTTPSListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLoadBalancerHTTPSListenerAttribute(request)
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

func (client *Client) SetLoadBalancerHTTPSListenerAttributeWithCallback(request *SetLoadBalancerHTTPSListenerAttributeRequest, callback func(response *SetLoadBalancerHTTPSListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLoadBalancerHTTPSListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.SetLoadBalancerHTTPSListenerAttribute(request)
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

type SetLoadBalancerHTTPSListenerAttributeRequest struct {
	*requests.RpcRequest
	VServerGroup           string `position:"Query" name:"VServerGroup"`
	XForwardedFor          string `position:"Query" name:"XForwardedFor"`
	UnhealthyThreshold     string `position:"Query" name:"UnhealthyThreshold"`
	Bandwidth              string `position:"Query" name:"Bandwidth"`
	HealthCheck            string `position:"Query" name:"HealthCheck"`
	XForwardedForSLBIP     string `position:"Query" name:"XForwardedFor_SLBIP"`
	HealthCheckDomain      string `position:"Query" name:"HealthCheckDomain"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	StickySession          string `position:"Query" name:"StickySession"`
	ResourceOwnerId        string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	XForwardedForSLBID     string `position:"Query" name:"XForwardedFor_SLBID"`
	Tags                   string `position:"Query" name:"Tags"`
	HealthCheckTimeout     string `position:"Query" name:"HealthCheckTimeout"`
	ServerCertificateId    string `position:"Query" name:"ServerCertificateId"`
	HealthCheckHttpCode    string `position:"Query" name:"HealthCheckHttpCode"`
	Gzip                   string `position:"Query" name:"Gzip"`
	Scheduler              string `position:"Query" name:"Scheduler"`
	OwnerId                string `position:"Query" name:"OwnerId"`
	XForwardedForProto     string `position:"Query" name:"XForwardedFor_proto"`
	VServerGroupId         string `position:"Query" name:"VServerGroupId"`
	Cookie                 string `position:"Query" name:"Cookie"`
	HealthCheckInterval    string `position:"Query" name:"HealthCheckInterval"`
	CACertificateId        string `position:"Query" name:"CACertificateId"`
	ListenerPort           string `position:"Query" name:"ListenerPort"`
	HealthCheckURI         string `position:"Query" name:"HealthCheckURI"`
	AccessKeyId            string `position:"Query" name:"access_key_id"`
	MaxConnection          string `position:"Query" name:"MaxConnection"`
	CookieTimeout          string `position:"Query" name:"CookieTimeout"`
	StickySessionType      string `position:"Query" name:"StickySessionType"`
	HealthCheckConnectPort string `position:"Query" name:"HealthCheckConnectPort"`
	LoadBalancerId         string `position:"Query" name:"LoadBalancerId"`
	HealthyThreshold       string `position:"Query" name:"HealthyThreshold"`
}

type SetLoadBalancerHTTPSListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSetLoadBalancerHTTPSListenerAttributeRequest() (request *SetLoadBalancerHTTPSListenerAttributeRequest) {
	request = &SetLoadBalancerHTTPSListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerHTTPSListenerAttribute", "", "")
	return
}

func CreateSetLoadBalancerHTTPSListenerAttributeResponse() (response *SetLoadBalancerHTTPSListenerAttributeResponse) {
	response = &SetLoadBalancerHTTPSListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
