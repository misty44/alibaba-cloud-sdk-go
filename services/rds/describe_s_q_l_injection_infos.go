package rds

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

func (client *Client) DescribeSQLInjectionInfos(request *DescribeSQLInjectionInfosRequest) (response *DescribeSQLInjectionInfosResponse, err error) {
	response = CreateDescribeSQLInjectionInfosResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeSQLInjectionInfosWithChan(request *DescribeSQLInjectionInfosRequest) (<-chan *DescribeSQLInjectionInfosResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLInjectionInfosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLInjectionInfos(request)
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

func (client *Client) DescribeSQLInjectionInfosWithCallback(request *DescribeSQLInjectionInfosRequest, callback func(response *DescribeSQLInjectionInfosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLInjectionInfosResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLInjectionInfos(request)
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

type DescribeSQLInjectionInfosRequest struct {
	*requests.RpcRequest
	EndTime              string `position:"Query" name:"EndTime"`
	PageSize             string `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DescribeSQLInjectionInfosResponse struct {
	*responses.BaseResponse
	RequestId        string          `json:"RequestId" xml:"RequestId"`
	Engine           string          `json:"Engine" xml:"Engine"`
	TotalRecordCount request.Integer `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       request.Integer `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  request.Integer `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            struct {
		SQLInjectionInfo []struct {
			DBName         string `json:"DBName" xml:"DBName"`
			SQLText        string `json:"SQLText" xml:"SQLText"`
			LatencyTime    string `json:"LatencyTime" xml:"LatencyTime"`
			HostAddress    string `json:"HostAddress" xml:"HostAddress"`
			ExecuteTime    string `json:"ExecuteTime" xml:"ExecuteTime"`
			AccountName    string `json:"AccountName" xml:"AccountName"`
			EffectRowCount string `json:"EffectRowCount" xml:"EffectRowCount"`
		} `json:"SQLInjectionInfo" xml:"SQLInjectionInfo"`
	} `json:"Items" xml:"Items"`
}

func CreateDescribeSQLInjectionInfosRequest() (request *DescribeSQLInjectionInfosRequest) {
	request = &DescribeSQLInjectionInfosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLInjectionInfos", "", "")
	return
}

func CreateDescribeSQLInjectionInfosResponse() (response *DescribeSQLInjectionInfosResponse) {
	response = &DescribeSQLInjectionInfosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
