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

func (client *Client) DescribeDomainFileSizeProportionData(request *DescribeDomainFileSizeProportionDataRequest) (response *DescribeDomainFileSizeProportionDataResponse, err error) {
	response = CreateDescribeDomainFileSizeProportionDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainFileSizeProportionDataWithChan(request *DescribeDomainFileSizeProportionDataRequest) (<-chan *DescribeDomainFileSizeProportionDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainFileSizeProportionDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainFileSizeProportionData(request)
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

func (client *Client) DescribeDomainFileSizeProportionDataWithCallback(request *DescribeDomainFileSizeProportionDataRequest, callback func(response *DescribeDomainFileSizeProportionDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainFileSizeProportionDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainFileSizeProportionData(request)
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

type DescribeDomainFileSizeProportionDataRequest struct {
	*requests.RpcRequest
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type DescribeDomainFileSizeProportionDataResponse struct {
	*responses.BaseResponse
	RequestId                      string `json:"RequestId" xml:"RequestId"`
	DomainName                     string `json:"DomainName" xml:"DomainName"`
	DataInterval                   string `json:"DataInterval" xml:"DataInterval"`
	StartTime                      string `json:"StartTime" xml:"StartTime"`
	EndTime                        string `json:"EndTime" xml:"EndTime"`
	FileSizeProportionDataInterval struct {
		UsageData []struct {
			TimeStamp string `json:"TimeStamp" xml:"TimeStamp"`
			Value     struct {
				FileSizeProportionData []struct {
					FileSize   string `json:"FileSize" xml:"FileSize"`
					Proportion string `json:"Proportion" xml:"Proportion"`
				} `json:"FileSizeProportionData" xml:"FileSizeProportionData"`
			} `json:"Value" xml:"Value"`
		} `json:"UsageData" xml:"UsageData"`
	} `json:"FileSizeProportionDataInterval" xml:"FileSizeProportionDataInterval"`
}

func CreateDescribeDomainFileSizeProportionDataRequest() (request *DescribeDomainFileSizeProportionDataRequest) {
	request = &DescribeDomainFileSizeProportionDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainFileSizeProportionData", "", "")
	return
}

func CreateDescribeDomainFileSizeProportionDataResponse() (response *DescribeDomainFileSizeProportionDataResponse) {
	response = &DescribeDomainFileSizeProportionDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
