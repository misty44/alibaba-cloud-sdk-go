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

func (client *Client) DescibeImportsFromDatabase(request *DescibeImportsFromDatabaseRequest) (response *DescibeImportsFromDatabaseResponse, err error) {
	response = CreateDescibeImportsFromDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescibeImportsFromDatabaseWithChan(request *DescibeImportsFromDatabaseRequest) (<-chan *DescibeImportsFromDatabaseResponse, <-chan error) {
	responseChan := make(chan *DescibeImportsFromDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescibeImportsFromDatabase(request)
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

func (client *Client) DescibeImportsFromDatabaseWithCallback(request *DescibeImportsFromDatabaseRequest, callback func(response *DescibeImportsFromDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescibeImportsFromDatabaseResponse
		var err error
		defer close(result)
		response, err = client.DescibeImportsFromDatabase(request)
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

type DescibeImportsFromDatabaseRequest struct {
	*requests.RpcRequest
	EndTime              string `position:"Query" name:"EndTime"`
	PageSize             string `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	ImportId             string `position:"Query" name:"ImportId"`
	StartTime            string `position:"Query" name:"StartTime"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	Engine               string `position:"Query" name:"Engine"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
}

type DescibeImportsFromDatabaseResponse struct {
	*responses.BaseResponse
	RequestId        string          `json:"RequestId" xml:"RequestId"`
	TotalRecordCount request.Integer `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       request.Integer `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  request.Integer `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            struct {
		ImportResultFromDB []struct {
			ImportId                    request.Integer `json:"ImportId" xml:"ImportId"`
			ImportDataType              string          `json:"ImportDataType" xml:"ImportDataType"`
			ImportDataStatus            string          `json:"ImportDataStatus" xml:"ImportDataStatus"`
			ImportDataStatusDescription string          `json:"ImportDataStatusDescription" xml:"ImportDataStatusDescription"`
			IncrementalImportingTime    string          `json:"IncrementalImportingTime" xml:"IncrementalImportingTime"`
		} `json:"ImportResultFromDB" xml:"ImportResultFromDB"`
	} `json:"Items" xml:"Items"`
}

func CreateDescibeImportsFromDatabaseRequest() (request *DescibeImportsFromDatabaseRequest) {
	request = &DescibeImportsFromDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescibeImportsFromDatabase", "", "")
	return
}

func CreateDescibeImportsFromDatabaseResponse() (response *DescibeImportsFromDatabaseResponse) {
	response = &DescibeImportsFromDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
