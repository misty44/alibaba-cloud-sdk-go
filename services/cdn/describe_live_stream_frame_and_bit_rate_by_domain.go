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

func (client *Client) DescribeLiveStreamFrameAndBitRateByDomain(request *DescribeLiveStreamFrameAndBitRateByDomainRequest) (response *DescribeLiveStreamFrameAndBitRateByDomainResponse, err error) {
	response = CreateDescribeLiveStreamFrameAndBitRateByDomainResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamFrameAndBitRateByDomainWithChan(request *DescribeLiveStreamFrameAndBitRateByDomainRequest) (<-chan *DescribeLiveStreamFrameAndBitRateByDomainResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamFrameAndBitRateByDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamFrameAndBitRateByDomain(request)
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

func (client *Client) DescribeLiveStreamFrameAndBitRateByDomainWithCallback(request *DescribeLiveStreamFrameAndBitRateByDomainRequest, callback func(response *DescribeLiveStreamFrameAndBitRateByDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamFrameAndBitRateByDomainResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamFrameAndBitRateByDomain(request)
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

type DescribeLiveStreamFrameAndBitRateByDomainRequest struct {
	*requests.RpcRequest
	PageSize      string `position:"Query" name:"PageSize"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageNumber    string `position:"Query" name:"PageNumber"`
	AppName       string `position:"Query" name:"AppName"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type DescribeLiveStreamFrameAndBitRateByDomainResponse struct {
	*responses.BaseResponse
	RequestId                string          `json:"RequestId" xml:"RequestId"`
	Count                    request.Integer `json:"Count" xml:"Count"`
	PageNumber               request.Integer `json:"pageNumber" xml:"pageNumber"`
	PageSize                 request.Integer `json:"pageSize" xml:"pageSize"`
	FrameRateAndBitRateInfos struct {
		FrameRateAndBitRateInfo []struct {
			StreamUrl      string        `json:"StreamUrl" xml:"StreamUrl"`
			VideoFrameRate request.Float `json:"VideoFrameRate" xml:"VideoFrameRate"`
			AudioFrameRate request.Float `json:"AudioFrameRate" xml:"AudioFrameRate"`
			BitRate        request.Float `json:"BitRate" xml:"BitRate"`
			Time           string        `json:"Time" xml:"Time"`
		} `json:"FrameRateAndBitRateInfo" xml:"FrameRateAndBitRateInfo"`
	} `json:"FrameRateAndBitRateInfos" xml:"FrameRateAndBitRateInfos"`
}

func CreateDescribeLiveStreamFrameAndBitRateByDomainRequest() (request *DescribeLiveStreamFrameAndBitRateByDomainRequest) {
	request = &DescribeLiveStreamFrameAndBitRateByDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamFrameAndBitRateByDomain", "", "")
	return
}

func CreateDescribeLiveStreamFrameAndBitRateByDomainResponse() (response *DescribeLiveStreamFrameAndBitRateByDomainResponse) {
	response = &DescribeLiveStreamFrameAndBitRateByDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
