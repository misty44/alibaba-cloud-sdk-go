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

func (client *Client) GetTrackList(request *GetTrackListRequest) (response *GetTrackListResponse, err error) {
	response = CreateGetTrackListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetTrackListWithChan(request *GetTrackListRequest) (<-chan *GetTrackListResponse, <-chan error) {
	responseChan := make(chan *GetTrackListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTrackList(request)
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

func (client *Client) GetTrackListWithCallback(request *GetTrackListRequest, callback func(response *GetTrackListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTrackListResponse
		var err error
		defer close(result)
		response, err = client.GetTrackList(request)
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

type GetTrackListRequest struct {
	*requests.RpcRequest
	Total                string `position:"Query" name:"Total"`
	EndTime              string `position:"Query" name:"EndTime"`
	PageSize             string `position:"Query" name:"PageSize"`
	StartTime            string `position:"Query" name:"StartTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	OffsetCreateTimeDesc string `position:"Query" name:"OffsetCreateTimeDesc"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	Offset               string `position:"Query" name:"Offset"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	OffsetCreateTime     string `position:"Query" name:"OffsetCreateTime"`
}

type GetTrackListResponse struct {
	*responses.BaseResponse
	RequestId            string          `json:"RequestId" xml:"RequestId"`
	Total                request.Integer `json:"Total" xml:"Total"`
	PageNo               request.Integer `json:"PageNo" xml:"PageNo"`
	PageSize             request.Integer `json:"PageSize" xml:"PageSize"`
	OffsetCreateTime     string          `json:"OffsetCreateTime" xml:"OffsetCreateTime"`
	OffsetCreateTimeDesc string          `json:"OffsetCreateTimeDesc" xml:"OffsetCreateTimeDesc"`
	Data                 struct {
		Stat []struct {
			CreateTime           string `json:"CreateTime" xml:"CreateTime"`
			RcptClickCount       string `json:"RcptClickCount" xml:"RcptClickCount"`
			RcptClickRate        string `json:"RcptClickRate" xml:"RcptClickRate"`
			RcptOpenCount        string `json:"RcptOpenCount" xml:"RcptOpenCount"`
			RcptOpenRate         string `json:"RcptOpenRate" xml:"RcptOpenRate"`
			RcptUniqueClickCount string `json:"RcptUniqueClickCount" xml:"RcptUniqueClickCount"`
			RcptUniqueClickRate  string `json:"RcptUniqueClickRate" xml:"RcptUniqueClickRate"`
			RcptUniqueOpenCount  string `json:"RcptUniqueOpenCount" xml:"RcptUniqueOpenCount"`
			RcptUniqueOpenRate   string `json:"RcptUniqueOpenRate" xml:"RcptUniqueOpenRate"`
			TotalNumber          string `json:"TotalNumber" xml:"TotalNumber"`
		} `json:"stat" xml:"stat"`
	} `json:"data" xml:"data"`
}

func CreateGetTrackListRequest() (request *GetTrackListRequest) {
	request = &GetTrackListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "GetTrackList", "", "")
	return
}

func CreateGetTrackListResponse() (response *GetTrackListResponse) {
	response = &GetTrackListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
