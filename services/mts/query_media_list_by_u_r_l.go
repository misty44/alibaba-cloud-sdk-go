package mts

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

func (client *Client) QueryMediaListByURL(request *QueryMediaListByURLRequest) (response *QueryMediaListByURLResponse, err error) {
	response = CreateQueryMediaListByURLResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryMediaListByURLWithChan(request *QueryMediaListByURLRequest) (<-chan *QueryMediaListByURLResponse, <-chan error) {
	responseChan := make(chan *QueryMediaListByURLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMediaListByURL(request)
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

func (client *Client) QueryMediaListByURLWithCallback(request *QueryMediaListByURLRequest, callback func(response *QueryMediaListByURLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMediaListByURLResponse
		var err error
		defer close(result)
		response, err = client.QueryMediaListByURL(request)
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

type QueryMediaListByURLRequest struct {
	*requests.RpcRequest
	IncludeMediaInfo     string `position:"Query" name:"IncludeMediaInfo"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	FileURLs             string `position:"Query" name:"FileURLs"`
	IncludeSnapshotList  string `position:"Query" name:"IncludeSnapshotList"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IncludePlayList      string `position:"Query" name:"IncludePlayList"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type QueryMediaListByURLResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	NonExistFileURLs struct {
		FileURL []string `json:"FileURL" xml:"FileURL"`
	} `json:"NonExistFileURLs" xml:"NonExistFileURLs"`
	MediaList struct {
		Media []struct {
			MediaId      string          `json:"MediaId" xml:"MediaId"`
			Title        string          `json:"Title" xml:"Title"`
			Description  string          `json:"Description" xml:"Description"`
			CoverURL     string          `json:"CoverURL" xml:"CoverURL"`
			CateId       request.Integer `json:"CateId" xml:"CateId"`
			Duration     string          `json:"Duration" xml:"Duration"`
			Format       string          `json:"Format" xml:"Format"`
			Size         string          `json:"Size" xml:"Size"`
			Bitrate      string          `json:"Bitrate" xml:"Bitrate"`
			Width        string          `json:"Width" xml:"Width"`
			Height       string          `json:"Height" xml:"Height"`
			Fps          string          `json:"Fps" xml:"Fps"`
			PublishState string          `json:"PublishState" xml:"PublishState"`
			CreationTime string          `json:"CreationTime" xml:"CreationTime"`
			Tags         struct {
				Tag []string `json:"Tag" xml:"Tag"`
			} `json:"Tags" xml:"Tags"`
			RunIdList struct {
				RunId []string `json:"RunId" xml:"RunId"`
			} `json:"RunIdList" xml:"RunIdList"`
			File struct {
				URL   string `json:"URL" xml:"URL"`
				State string `json:"State" xml:"State"`
			} `json:"File" xml:"File"`
			MediaInfo struct {
				Streams struct {
					VideoStreamList struct {
						VideoStream []struct {
							Index          string `json:"Index" xml:"Index"`
							CodecName      string `json:"CodecName" xml:"CodecName"`
							CodecLongName  string `json:"CodecLongName" xml:"CodecLongName"`
							Profile        string `json:"Profile" xml:"Profile"`
							CodecTimeBase  string `json:"CodecTimeBase" xml:"CodecTimeBase"`
							CodecTagString string `json:"CodecTagString" xml:"CodecTagString"`
							CodecTag       string `json:"CodecTag" xml:"CodecTag"`
							Width          string `json:"Width" xml:"Width"`
							Height         string `json:"Height" xml:"Height"`
							HasBFrames     string `json:"HasBFrames" xml:"HasBFrames"`
							Sar            string `json:"Sar" xml:"Sar"`
							Dar            string `json:"Dar" xml:"Dar"`
							PixFmt         string `json:"PixFmt" xml:"PixFmt"`
							Level          string `json:"Level" xml:"Level"`
							Fps            string `json:"Fps" xml:"Fps"`
							AvgFPS         string `json:"AvgFPS" xml:"AvgFPS"`
							Timebase       string `json:"Timebase" xml:"Timebase"`
							StartTime      string `json:"StartTime" xml:"StartTime"`
							Duration       string `json:"Duration" xml:"Duration"`
							Bitrate        string `json:"Bitrate" xml:"Bitrate"`
							NumFrames      string `json:"NumFrames" xml:"NumFrames"`
							Lang           string `json:"Lang" xml:"Lang"`
							Rotate         string `json:"Rotate" xml:"Rotate"`
							NetworkCost    struct {
								PreloadTime   string `json:"PreloadTime" xml:"PreloadTime"`
								CostBandwidth string `json:"CostBandwidth" xml:"CostBandwidth"`
								AvgBitrate    string `json:"AvgBitrate" xml:"AvgBitrate"`
							} `json:"NetworkCost" xml:"NetworkCost"`
						} `json:"VideoStream" xml:"VideoStream"`
					} `json:"VideoStreamList" xml:"VideoStreamList"`
					AudioStreamList struct {
						AudioStream []struct {
							Index          string `json:"Index" xml:"Index"`
							CodecName      string `json:"CodecName" xml:"CodecName"`
							CodecTimeBase  string `json:"CodecTimeBase" xml:"CodecTimeBase"`
							CodecLongName  string `json:"CodecLongName" xml:"CodecLongName"`
							CodecTagString string `json:"CodecTagString" xml:"CodecTagString"`
							CodecTag       string `json:"CodecTag" xml:"CodecTag"`
							SampleFmt      string `json:"SampleFmt" xml:"SampleFmt"`
							Samplerate     string `json:"Samplerate" xml:"Samplerate"`
							Channels       string `json:"Channels" xml:"Channels"`
							ChannelLayout  string `json:"ChannelLayout" xml:"ChannelLayout"`
							Timebase       string `json:"Timebase" xml:"Timebase"`
							StartTime      string `json:"StartTime" xml:"StartTime"`
							Duration       string `json:"Duration" xml:"Duration"`
							Bitrate        string `json:"Bitrate" xml:"Bitrate"`
							NumFrames      string `json:"NumFrames" xml:"NumFrames"`
							Lang           string `json:"Lang" xml:"Lang"`
						} `json:"AudioStream" xml:"AudioStream"`
					} `json:"AudioStreamList" xml:"AudioStreamList"`
					SubtitleStreamList struct {
						SubtitleStream []struct {
							Index string `json:"Index" xml:"Index"`
							Lang  string `json:"Lang" xml:"Lang"`
						} `json:"SubtitleStream" xml:"SubtitleStream"`
					} `json:"SubtitleStreamList" xml:"SubtitleStreamList"`
				} `json:"Streams" xml:"Streams"`
				Format struct {
					NumStreams     string `json:"NumStreams" xml:"NumStreams"`
					NumPrograms    string `json:"NumPrograms" xml:"NumPrograms"`
					FormatName     string `json:"FormatName" xml:"FormatName"`
					FormatLongName string `json:"FormatLongName" xml:"FormatLongName"`
					StartTime      string `json:"StartTime" xml:"StartTime"`
					Duration       string `json:"Duration" xml:"Duration"`
					Size           string `json:"Size" xml:"Size"`
					Bitrate        string `json:"Bitrate" xml:"Bitrate"`
				} `json:"Format" xml:"Format"`
			} `json:"MediaInfo" xml:"MediaInfo"`
			PlayList struct {
				Play []struct {
					ActivityName      string `json:"ActivityName" xml:"ActivityName"`
					MediaWorkflowId   string `json:"MediaWorkflowId" xml:"MediaWorkflowId"`
					MediaWorkflowName string `json:"MediaWorkflowName" xml:"MediaWorkflowName"`
					Duration          string `json:"Duration" xml:"Duration"`
					Format            string `json:"Format" xml:"Format"`
					Size              string `json:"Size" xml:"Size"`
					Bitrate           string `json:"Bitrate" xml:"Bitrate"`
					Width             string `json:"Width" xml:"Width"`
					Height            string `json:"Height" xml:"Height"`
					Fps               string `json:"Fps" xml:"Fps"`
					Encryption        string `json:"Encryption" xml:"Encryption"`
					File1             struct {
						URL   string `json:"URL" xml:"URL"`
						State string `json:"State" xml:"State"`
					} `json:"File" xml:"File"`
				} `json:"Play" xml:"Play"`
			} `json:"PlayList" xml:"PlayList"`
			SnapshotList struct {
				Snapshot []struct {
					Type              string `json:"Type" xml:"Type"`
					MediaWorkflowId   string `json:"MediaWorkflowId" xml:"MediaWorkflowId"`
					MediaWorkflowName string `json:"MediaWorkflowName" xml:"MediaWorkflowName"`
					ActivityName      string `json:"ActivityName" xml:"ActivityName"`
					Count             string `json:"Count" xml:"Count"`
					File2             struct {
						URL   string `json:"URL" xml:"URL"`
						State string `json:"State" xml:"State"`
					} `json:"File" xml:"File"`
				} `json:"Snapshot" xml:"Snapshot"`
			} `json:"SnapshotList" xml:"SnapshotList"`
		} `json:"Media" xml:"Media"`
	} `json:"MediaList" xml:"MediaList"`
}

func CreateQueryMediaListByURLRequest() (request *QueryMediaListByURLRequest) {
	request = &QueryMediaListByURLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaListByURL", "", "")
	return
}

func CreateQueryMediaListByURLResponse() (response *QueryMediaListByURLResponse) {
	response = &QueryMediaListByURLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
