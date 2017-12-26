package cms

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

func (client *Client) CreateAlarm(request *CreateAlarmRequest) (response *CreateAlarmResponse, err error) {
	response = CreateCreateAlarmResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateAlarmWithChan(request *CreateAlarmRequest) (<-chan *CreateAlarmResponse, <-chan error) {
	responseChan := make(chan *CreateAlarmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAlarm(request)
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

func (client *Client) CreateAlarmWithCallback(request *CreateAlarmRequest, callback func(response *CreateAlarmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAlarmResponse
		var err error
		defer close(result)
		response, err = client.CreateAlarm(request)
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

type CreateAlarmRequest struct {
	*requests.RpcRequest
	EndTime            string `position:"Query" name:"EndTime"`
	ComparisonOperator string `position:"Query" name:"ComparisonOperator"`
	StartTime          string `position:"Query" name:"StartTime"`
	NotifyType         string `position:"Query" name:"NotifyType"`
	Period             string `position:"Query" name:"Period"`
	Namespace          string `position:"Query" name:"Namespace"`
	Statistics         string `position:"Query" name:"Statistics"`
	Threshold          string `position:"Query" name:"Threshold"`
	MetricName         string `position:"Query" name:"MetricName"`
	Webhook            string `position:"Query" name:"Webhook"`
	Name               string `position:"Query" name:"Name"`
	EvaluationCount    string `position:"Query" name:"EvaluationCount"`
	Dimensions         string `position:"Query" name:"Dimensions"`
	SilenceTime        string `position:"Query" name:"SilenceTime"`
	ContactGroups      string `position:"Query" name:"ContactGroups"`
	CallbyCmsOwner     string `position:"Query" name:"callby_cms_owner"`
}

type CreateAlarmResponse struct {
	*responses.BaseResponse
	Success   request.Boolean `json:"Success" xml:"Success"`
	Code      string          `json:"Code" xml:"Code"`
	Message   string          `json:"Message" xml:"Message"`
	RequestId string          `json:"RequestId" xml:"RequestId"`
	Data      string          `json:"Data" xml:"Data"`
}

func CreateCreateAlarmRequest() (request *CreateAlarmRequest) {
	request = &CreateAlarmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "CreateAlarm", "", "")
	return
}

func CreateCreateAlarmResponse() (response *CreateAlarmResponse) {
	response = &CreateAlarmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
