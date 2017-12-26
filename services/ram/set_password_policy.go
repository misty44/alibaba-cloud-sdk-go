package ram

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

func (client *Client) SetPasswordPolicy(request *SetPasswordPolicyRequest) (response *SetPasswordPolicyResponse, err error) {
	response = CreateSetPasswordPolicyResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetPasswordPolicyWithChan(request *SetPasswordPolicyRequest) (<-chan *SetPasswordPolicyResponse, <-chan error) {
	responseChan := make(chan *SetPasswordPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetPasswordPolicy(request)
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

func (client *Client) SetPasswordPolicyWithCallback(request *SetPasswordPolicyRequest, callback func(response *SetPasswordPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetPasswordPolicyResponse
		var err error
		defer close(result)
		response, err = client.SetPasswordPolicy(request)
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

type SetPasswordPolicyRequest struct {
	*requests.RpcRequest
	MaxLoginAttemps            string `position:"Query" name:"MaxLoginAttemps"`
	RequireUppercaseCharacters string `position:"Query" name:"RequireUppercaseCharacters"`
	MinimumPasswordLength      string `position:"Query" name:"MinimumPasswordLength"`
	MaxPasswordAge             string `position:"Query" name:"MaxPasswordAge"`
	RequireNumbers             string `position:"Query" name:"RequireNumbers"`
	RequireLowercaseCharacters string `position:"Query" name:"RequireLowercaseCharacters"`
	PasswordReusePrevention    string `position:"Query" name:"PasswordReusePrevention"`
	HardExpiry                 string `position:"Query" name:"HardExpiry"`
	RequireSymbols             string `position:"Query" name:"RequireSymbols"`
}

type SetPasswordPolicyResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	PasswordPolicy struct {
		MinimumPasswordLength      request.Integer `json:"MinimumPasswordLength" xml:"MinimumPasswordLength"`
		RequireLowercaseCharacters request.Boolean `json:"RequireLowercaseCharacters" xml:"RequireLowercaseCharacters"`
		RequireUppercaseCharacters request.Boolean `json:"RequireUppercaseCharacters" xml:"RequireUppercaseCharacters"`
		RequireNumbers             request.Boolean `json:"RequireNumbers" xml:"RequireNumbers"`
		RequireSymbols             request.Boolean `json:"RequireSymbols" xml:"RequireSymbols"`
		HardExpiry                 request.Boolean `json:"HardExpiry" xml:"HardExpiry"`
		MaxPasswordAge             request.Integer `json:"MaxPasswordAge" xml:"MaxPasswordAge"`
		PasswordReusePrevention    request.Integer `json:"PasswordReusePrevention" xml:"PasswordReusePrevention"`
		MaxLoginAttemps            request.Integer `json:"MaxLoginAttemps" xml:"MaxLoginAttemps"`
	} `json:"PasswordPolicy" xml:"PasswordPolicy"`
}

func CreateSetPasswordPolicyRequest() (request *SetPasswordPolicyRequest) {
	request = &SetPasswordPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "SetPasswordPolicy", "", "")
	return
}

func CreateSetPasswordPolicyResponse() (response *SetPasswordPolicyResponse) {
	response = &SetPasswordPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
