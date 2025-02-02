package dypnsapi

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
	"github.com/misty44/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/misty44/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeVerifyScheme invokes the dypnsapi.DescribeVerifyScheme API synchronously
// api document: https://help.aliyun.com/api/dypnsapi/describeverifyscheme.html
func (client *Client) DescribeVerifyScheme(request *DescribeVerifySchemeRequest) (response *DescribeVerifySchemeResponse, err error) {
	response = CreateDescribeVerifySchemeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVerifySchemeWithChan invokes the dypnsapi.DescribeVerifyScheme API asynchronously
// api document: https://help.aliyun.com/api/dypnsapi/describeverifyscheme.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVerifySchemeWithChan(request *DescribeVerifySchemeRequest) (<-chan *DescribeVerifySchemeResponse, <-chan error) {
	responseChan := make(chan *DescribeVerifySchemeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVerifyScheme(request)
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

// DescribeVerifySchemeWithCallback invokes the dypnsapi.DescribeVerifyScheme API asynchronously
// api document: https://help.aliyun.com/api/dypnsapi/describeverifyscheme.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVerifySchemeWithCallback(request *DescribeVerifySchemeRequest, callback func(response *DescribeVerifySchemeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVerifySchemeResponse
		var err error
		defer close(result)
		response, err = client.DescribeVerifyScheme(request)
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

// DescribeVerifySchemeRequest is the request struct for api DescribeVerifyScheme
type DescribeVerifySchemeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SchemeCode           string           `position:"Query" name:"SchemeCode"`
	CustomerId           requests.Integer `position:"Query" name:"CustomerId"`
}

// DescribeVerifySchemeResponse is the response struct for api DescribeVerifyScheme
type DescribeVerifySchemeResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	Code                 string               `json:"Code" xml:"Code"`
	Message              string               `json:"Message" xml:"Message"`
	SchemeQueryResultDTO SchemeQueryResultDTO `json:"SchemeQueryResultDTO" xml:"SchemeQueryResultDTO"`
}

// CreateDescribeVerifySchemeRequest creates a request to invoke DescribeVerifyScheme API
func CreateDescribeVerifySchemeRequest() (request *DescribeVerifySchemeRequest) {
	request = &DescribeVerifySchemeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi", "2017-05-25", "DescribeVerifyScheme", "dypns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVerifySchemeResponse creates a response to parse from DescribeVerifyScheme response
func CreateDescribeVerifySchemeResponse() (response *DescribeVerifySchemeResponse) {
	response = &DescribeVerifySchemeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
