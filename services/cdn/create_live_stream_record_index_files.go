package cdn

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

func (client *Client) CreateLiveStreamRecordIndexFiles(request *CreateLiveStreamRecordIndexFilesRequest) (response *CreateLiveStreamRecordIndexFilesResponse, err error) {
	response = CreateCreateLiveStreamRecordIndexFilesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateLiveStreamRecordIndexFilesWithChan(request *CreateLiveStreamRecordIndexFilesRequest) (<-chan *CreateLiveStreamRecordIndexFilesResponse, <-chan error) {
	responseChan := make(chan *CreateLiveStreamRecordIndexFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLiveStreamRecordIndexFiles(request)
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

func (client *Client) CreateLiveStreamRecordIndexFilesWithCallback(request *CreateLiveStreamRecordIndexFilesRequest, callback func(response *CreateLiveStreamRecordIndexFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLiveStreamRecordIndexFilesResponse
		var err error
		defer close(result)
		response, err = client.CreateLiveStreamRecordIndexFiles(request)
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

type CreateLiveStreamRecordIndexFilesRequest struct {
	*requests.RpcRequest
	EndTime       string `position:"Query" name:"EndTime"`
	StreamName    string `position:"Query" name:"StreamName"`
	StartTime     string `position:"Query" name:"StartTime"`
	OssBucket     string `position:"Query" name:"OssBucket"`
	DomainName    string `position:"Query" name:"DomainName"`
	OssEndpoint   string `position:"Query" name:"OssEndpoint"`
	AppName       string `position:"Query" name:"AppName"`
	OssObject     string `position:"Query" name:"OssObject"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type CreateLiveStreamRecordIndexFilesResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	RecordInfo struct {
		RecordId   string          `json:"RecordId" xml:"RecordId"`
		RecordUrl  string          `json:"RecordUrl" xml:"RecordUrl"`
		Duration   request.Float   `json:"Duration" xml:"Duration"`
		Height     request.Integer `json:"Height" xml:"Height"`
		Width      request.Integer `json:"Width" xml:"Width"`
		CreateTime string          `json:"CreateTime" xml:"CreateTime"`
	} `json:"RecordInfo" xml:"RecordInfo"`
}

func CreateCreateLiveStreamRecordIndexFilesRequest() (request *CreateLiveStreamRecordIndexFilesRequest) {
	request = &CreateLiveStreamRecordIndexFilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "CreateLiveStreamRecordIndexFiles", "", "")
	return
}

func CreateCreateLiveStreamRecordIndexFilesResponse() (response *CreateLiveStreamRecordIndexFilesResponse) {
	response = &CreateLiveStreamRecordIndexFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
