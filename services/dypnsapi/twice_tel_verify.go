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

// TwiceTelVerify invokes the dypnsapi.TwiceTelVerify API synchronously
// api document: https://help.aliyun.com/api/dypnsapi/twicetelverify.html
func (client *Client) TwiceTelVerify(request *TwiceTelVerifyRequest) (response *TwiceTelVerifyResponse, err error) {
	response = CreateTwiceTelVerifyResponse()
	err = client.DoAction(request, response)
	return
}

// TwiceTelVerifyWithChan invokes the dypnsapi.TwiceTelVerify API asynchronously
// api document: https://help.aliyun.com/api/dypnsapi/twicetelverify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TwiceTelVerifyWithChan(request *TwiceTelVerifyRequest) (<-chan *TwiceTelVerifyResponse, <-chan error) {
	responseChan := make(chan *TwiceTelVerifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TwiceTelVerify(request)
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

// TwiceTelVerifyWithCallback invokes the dypnsapi.TwiceTelVerify API asynchronously
// api document: https://help.aliyun.com/api/dypnsapi/twicetelverify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TwiceTelVerifyWithCallback(request *TwiceTelVerifyRequest, callback func(response *TwiceTelVerifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TwiceTelVerifyResponse
		var err error
		defer close(result)
		response, err = client.TwiceTelVerify(request)
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

// TwiceTelVerifyRequest is the request struct for api TwiceTelVerify
type TwiceTelVerifyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PhoneNumber          string           `position:"Query" name:"PhoneNumber"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Since                string           `position:"Query" name:"Since"`
}

// TwiceTelVerifyResponse is the response struct for api TwiceTelVerify
type TwiceTelVerifyResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	Code                 string               `json:"Code" xml:"Code"`
	Message              string               `json:"Message" xml:"Message"`
	TwiceTelVerifyResult TwiceTelVerifyResult `json:"TwiceTelVerifyResult" xml:"TwiceTelVerifyResult"`
}

// CreateTwiceTelVerifyRequest creates a request to invoke TwiceTelVerify API
func CreateTwiceTelVerifyRequest() (request *TwiceTelVerifyRequest) {
	request = &TwiceTelVerifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi", "2017-05-25", "TwiceTelVerify", "dypns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTwiceTelVerifyResponse creates a response to parse from TwiceTelVerify response
func CreateTwiceTelVerifyResponse() (response *TwiceTelVerifyResponse) {
	response = &TwiceTelVerifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
