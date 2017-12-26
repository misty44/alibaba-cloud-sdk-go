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

func (client *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
	response = CreateDescribeTasksResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeTasksWithChan(request *DescribeTasksRequest) (<-chan *DescribeTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTasks(request)
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

func (client *Client) DescribeTasksWithCallback(request *DescribeTasksRequest, callback func(response *DescribeTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeTasks(request)
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

type DescribeTasksRequest struct {
	*requests.RpcRequest
	EndTime              string `position:"Query" name:"EndTime"`
	PageSize             string `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	Status               string `position:"Query" name:"Status"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	TaskAction           string `position:"Query" name:"TaskAction"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
}

type DescribeTasksResponse struct {
	*responses.BaseResponse
	RequestId        string          `json:"RequestId" xml:"RequestId"`
	TotalRecordCount request.Integer `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       request.Integer `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  request.Integer `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            struct {
		TaskProgressInfo []struct {
			DBName             string `json:"DBName" xml:"DBName"`
			BeginTime          string `json:"BeginTime" xml:"BeginTime"`
			ProgressInfo       string `json:"ProgressInfo" xml:"ProgressInfo"`
			FinishTime         string `json:"FinishTime" xml:"FinishTime"`
			TaskAction         string `json:"TaskAction" xml:"TaskAction"`
			TaskId             string `json:"TaskId" xml:"TaskId"`
			Progress           string `json:"Progress" xml:"Progress"`
			ExpectedFinishTime string `json:"ExpectedFinishTime" xml:"ExpectedFinishTime"`
			Status             string `json:"Status" xml:"Status"`
			TaskErrorCode      string `json:"TaskErrorCode" xml:"TaskErrorCode"`
			TaskErrorMessage   string `json:"TaskErrorMessage" xml:"TaskErrorMessage"`
		} `json:"TaskProgressInfo" xml:"TaskProgressInfo"`
	} `json:"Items" xml:"Items"`
}

func CreateDescribeTasksRequest() (request *DescribeTasksRequest) {
	request = &DescribeTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeTasks", "", "")
	return
}

func CreateDescribeTasksResponse() (response *DescribeTasksResponse) {
	response = &DescribeTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
