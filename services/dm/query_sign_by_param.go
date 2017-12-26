package dm

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

func (client *Client) QuerySignByParam(request *QuerySignByParamRequest) (response *QuerySignByParamResponse, err error) {
	response = CreateQuerySignByParamResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QuerySignByParamWithChan(request *QuerySignByParamRequest) (<-chan *QuerySignByParamResponse, <-chan error) {
	responseChan := make(chan *QuerySignByParamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySignByParam(request)
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

func (client *Client) QuerySignByParamWithCallback(request *QuerySignByParamRequest, callback func(response *QuerySignByParamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySignByParamResponse
		var err error
		defer close(result)
		response, err = client.QuerySignByParam(request)
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

type QuerySignByParamRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	KeyWord              string `position:"Query" name:"KeyWord"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	PageNo               string `position:"Query" name:"PageNo"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	FromType             string `position:"Query" name:"FromType"`
}

type QuerySignByParamResponse struct {
	*responses.BaseResponse
	RequestId  string          `json:"RequestId" xml:"RequestId"`
	PageNumber request.Integer `json:"PageNumber" xml:"PageNumber"`
	PageSize   request.Integer `json:"PageSize" xml:"PageSize"`
	Data       struct {
		Sign []struct {
			SignId     request.Integer `json:"SignId" xml:"SignId"`
			OrderId    string          `json:"OrderId" xml:"OrderId"`
			Remark     string          `json:"Remark" xml:"Remark"`
			AuditState string          `json:"AuditState" xml:"AuditState"`
			SignType   string          `json:"SignType" xml:"SignType"`
			GmtCreate  string          `json:"GmtCreate" xml:"GmtCreate"`
			SignName   string          `json:"SignName" xml:"SignName"`
			RejectInfo string          `json:"RejectInfo" xml:"RejectInfo"`
		} `json:"sign" xml:"sign"`
	} `json:"data" xml:"data"`
}

func CreateQuerySignByParamRequest() (request *QuerySignByParamRequest) {
	request = &QuerySignByParamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "QuerySignByParam", "", "")
	return
}

func CreateQuerySignByParamResponse() (response *QuerySignByParamResponse) {
	response = &QuerySignByParamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
