package cs

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

func (client *Client) ScaleCluster(request *ScaleClusterRequest) (response *ScaleClusterResponse, err error) {
	response = CreateScaleClusterResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ScaleClusterWithChan(request *ScaleClusterRequest) (<-chan *ScaleClusterResponse, <-chan error) {
	responseChan := make(chan *ScaleClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ScaleCluster(request)
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

func (client *Client) ScaleClusterWithCallback(request *ScaleClusterRequest, callback func(response *ScaleClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ScaleClusterResponse
		var err error
		defer close(result)
		response, err = client.ScaleCluster(request)
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

type ScaleClusterRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

type ScaleClusterResponse struct {
	*responses.BaseResponse
}

func CreateScaleClusterRequest() (request *ScaleClusterRequest) {
	request = &ScaleClusterRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "ScaleCluster", "/clusters/[ClusterId]", "", "")
	return
}

func CreateScaleClusterResponse() (response *ScaleClusterResponse) {
	response = &ScaleClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
