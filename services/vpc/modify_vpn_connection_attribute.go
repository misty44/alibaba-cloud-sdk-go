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

func (client *Client) ModifyVpnConnectionAttribute(request *ModifyVpnConnectionAttributeRequest) (response *ModifyVpnConnectionAttributeResponse, err error) {
	response = CreateModifyVpnConnectionAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyVpnConnectionAttributeWithChan(request *ModifyVpnConnectionAttributeRequest) (<-chan *ModifyVpnConnectionAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyVpnConnectionAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVpnConnectionAttribute(request)
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

func (client *Client) ModifyVpnConnectionAttributeWithCallback(request *ModifyVpnConnectionAttributeRequest, callback func(response *ModifyVpnConnectionAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVpnConnectionAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyVpnConnectionAttribute(request)
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

type ModifyVpnConnectionAttributeRequest struct {
	*requests.RpcRequest
	ClientToken          string `position:"Query" name:"ClientToken"`
	LocalSubnet          string `position:"Query" name:"LocalSubnet"`
	EffectImmediately    string `position:"Query" name:"EffectImmediately"`
	RemoteSubnet         string `position:"Query" name:"RemoteSubnet"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	IpsecConfig          string `position:"Query" name:"IpsecConfig"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	Name                 string `position:"Query" name:"Name"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IkeConfig            string `position:"Query" name:"IkeConfig"`
}

type ModifyVpnConnectionAttributeResponse struct {
	*responses.BaseResponse
	RequestId         string          `json:"RequestId" xml:"RequestId"`
	VpnConnectionId   string          `json:"VpnConnectionId" xml:"VpnConnectionId"`
	CustomerGatewayId string          `json:"CustomerGatewayId" xml:"CustomerGatewayId"`
	VpnGatewayId      string          `json:"VpnGatewayId" xml:"VpnGatewayId"`
	Name              string          `json:"Name" xml:"Name"`
	Description       string          `json:"Description" xml:"Description"`
	LocalSubnet       string          `json:"LocalSubnet" xml:"LocalSubnet"`
	RemoteSubnet      string          `json:"RemoteSubnet" xml:"RemoteSubnet"`
	CreateTime        request.Integer `json:"CreateTime" xml:"CreateTime"`
	EffectImmediately request.Boolean `json:"EffectImmediately" xml:"EffectImmediately"`
	IkeConfig         struct {
		Psk         string          `json:"Psk" xml:"Psk"`
		IkeVersion  string          `json:"IkeVersion" xml:"IkeVersion"`
		IkeMode     string          `json:"IkeMode" xml:"IkeMode"`
		IkeEncAlg   string          `json:"IkeEncAlg" xml:"IkeEncAlg"`
		IkeAuthAlg  string          `json:"IkeAuthAlg" xml:"IkeAuthAlg"`
		IkePfs      string          `json:"IkePfs" xml:"IkePfs"`
		IkeLifetime request.Integer `json:"IkeLifetime" xml:"IkeLifetime"`
		LocalId     string          `json:"LocalId" xml:"LocalId"`
		RemoteId    string          `json:"RemoteId" xml:"RemoteId"`
	} `json:"IkeConfig" xml:"IkeConfig"`
	IpsecConfig struct {
		IpsecEncAlg   string          `json:"IpsecEncAlg" xml:"IpsecEncAlg"`
		IpsecAuthAlg  string          `json:"IpsecAuthAlg" xml:"IpsecAuthAlg"`
		IpsecPfs      string          `json:"IpsecPfs" xml:"IpsecPfs"`
		IpsecLifetime request.Integer `json:"IpsecLifetime" xml:"IpsecLifetime"`
	} `json:"IpsecConfig" xml:"IpsecConfig"`
}

func CreateModifyVpnConnectionAttributeRequest() (request *ModifyVpnConnectionAttributeRequest) {
	request = &ModifyVpnConnectionAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVpnConnectionAttribute", "", "")
	return
}

func CreateModifyVpnConnectionAttributeResponse() (response *ModifyVpnConnectionAttributeResponse) {
	response = &ModifyVpnConnectionAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
