package nas

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

func (client *Client) DescribeAccessRules(request *DescribeAccessRulesRequest) (response *DescribeAccessRulesResponse, err error) {
	response = CreateDescribeAccessRulesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeAccessRulesWithChan(request *DescribeAccessRulesRequest) (<-chan *DescribeAccessRulesResponse, <-chan error) {
	responseChan := make(chan *DescribeAccessRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAccessRules(request)
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

func (client *Client) DescribeAccessRulesWithCallback(request *DescribeAccessRulesRequest, callback func(response *DescribeAccessRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAccessRulesResponse
		var err error
		defer close(result)
		response, err = client.DescribeAccessRules(request)
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

type DescribeAccessRulesRequest struct {
	*requests.RpcRequest
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	PageSize        string `position:"Query" name:"PageSize"`
	AccessRuleId    string `position:"Query" name:"AccessRuleId"`
	PageNumber      string `position:"Query" name:"PageNumber"`
}

type DescribeAccessRulesResponse struct {
	*responses.BaseResponse
	RequestId   string          `json:"RequestId" xml:"RequestId"`
	TotalCount  request.Integer `json:"TotalCount" xml:"TotalCount"`
	PageSize    request.Integer `json:"PageSize" xml:"PageSize"`
	PageNumber  request.Integer `json:"PageNumber" xml:"PageNumber"`
	AccessRules struct {
		AccessRule []struct {
			SourceCidrIp string          `json:"SourceCidrIp" xml:"SourceCidrIp"`
			Priority     request.Integer `json:"Priority" xml:"Priority"`
			AccessRuleId string          `json:"AccessRuleId" xml:"AccessRuleId"`
			RWAccess     string          `json:"RWAccess" xml:"RWAccess"`
			UserAccess   string          `json:"UserAccess" xml:"UserAccess"`
		} `json:"AccessRule" xml:"AccessRule"`
	} `json:"AccessRules" xml:"AccessRules"`
}

func CreateDescribeAccessRulesRequest() (request *DescribeAccessRulesRequest) {
	request = &DescribeAccessRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "DescribeAccessRules", "", "")
	return
}

func CreateDescribeAccessRulesResponse() (response *DescribeAccessRulesResponse) {
	response = &DescribeAccessRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
