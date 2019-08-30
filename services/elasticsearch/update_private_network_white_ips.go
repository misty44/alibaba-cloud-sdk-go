package elasticsearch

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

// UpdatePrivateNetworkWhiteIps invokes the elasticsearch.UpdatePrivateNetworkWhiteIps API synchronously
// api document: https://help.aliyun.com/api/elasticsearch/updateprivatenetworkwhiteips.html
func (client *Client) UpdatePrivateNetworkWhiteIps(request *UpdatePrivateNetworkWhiteIpsRequest) (response *UpdatePrivateNetworkWhiteIpsResponse, err error) {
	response = CreateUpdatePrivateNetworkWhiteIpsResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePrivateNetworkWhiteIpsWithChan invokes the elasticsearch.UpdatePrivateNetworkWhiteIps API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/updateprivatenetworkwhiteips.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdatePrivateNetworkWhiteIpsWithChan(request *UpdatePrivateNetworkWhiteIpsRequest) (<-chan *UpdatePrivateNetworkWhiteIpsResponse, <-chan error) {
	responseChan := make(chan *UpdatePrivateNetworkWhiteIpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePrivateNetworkWhiteIps(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UpdatePrivateNetworkWhiteIpsWithCallback invokes the elasticsearch.UpdatePrivateNetworkWhiteIps API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/updateprivatenetworkwhiteips.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdatePrivateNetworkWhiteIpsWithCallback(request *UpdatePrivateNetworkWhiteIpsRequest, callback func(response *UpdatePrivateNetworkWhiteIpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePrivateNetworkWhiteIpsResponse
		var err error
		defer close(result)
		response, err = client.UpdatePrivateNetworkWhiteIps(request)
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

// UpdatePrivateNetworkWhiteIpsRequest is the request struct for api UpdatePrivateNetworkWhiteIps
type UpdatePrivateNetworkWhiteIpsRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
}

// UpdatePrivateNetworkWhiteIpsResponse is the response struct for api UpdatePrivateNetworkWhiteIps
type UpdatePrivateNetworkWhiteIpsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateUpdatePrivateNetworkWhiteIpsRequest creates a request to invoke UpdatePrivateNetworkWhiteIps API
func CreateUpdatePrivateNetworkWhiteIpsRequest() (request *UpdatePrivateNetworkWhiteIpsRequest) {
	request = &UpdatePrivateNetworkWhiteIpsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "UpdatePrivateNetworkWhiteIps", "/openapi/instances/[InstanceId]/private-network-white-ips", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdatePrivateNetworkWhiteIpsResponse creates a response to parse from UpdatePrivateNetworkWhiteIps response
func CreateUpdatePrivateNetworkWhiteIpsResponse() (response *UpdatePrivateNetworkWhiteIpsResponse) {
	response = &UpdatePrivateNetworkWhiteIpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
