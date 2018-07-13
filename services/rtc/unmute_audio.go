package rtc

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

// UnmuteAudio invokes the rtc.UnmuteAudio API synchronously
// api document: https://help.aliyun.com/api/rtc/unmuteaudio.html
func (client *Client) UnmuteAudio(request *UnmuteAudioRequest) (response *UnmuteAudioResponse, err error) {
	response = CreateUnmuteAudioResponse()
	err = client.DoAction(request, response)
	return
}

// UnmuteAudioWithChan invokes the rtc.UnmuteAudio API asynchronously
// api document: https://help.aliyun.com/api/rtc/unmuteaudio.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnmuteAudioWithChan(request *UnmuteAudioRequest) (<-chan *UnmuteAudioResponse, <-chan error) {
	responseChan := make(chan *UnmuteAudioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnmuteAudio(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UnmuteAudioWithCallback invokes the rtc.UnmuteAudio API asynchronously
// api document: https://help.aliyun.com/api/rtc/unmuteaudio.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnmuteAudioWithCallback(request *UnmuteAudioRequest, callback func(response *UnmuteAudioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnmuteAudioResponse
		var err error
		defer close(result)
		response, err = client.UnmuteAudio(request)
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

// UnmuteAudioRequest is the request struct for api UnmuteAudio
type UnmuteAudioRequest struct {
	*requests.RpcRequest
	ParticipantIds *[]string        `position:"Query" name:"ParticipantIds"  type:"Repeated"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	ConferenceId   string           `position:"Query" name:"ConferenceId"`
	AppId          string           `position:"Query" name:"AppId"`
}

// UnmuteAudioResponse is the response struct for api UnmuteAudio
type UnmuteAudioResponse struct {
	*responses.BaseResponse
	RequestId    string                    `json:"RequestId" xml:"RequestId"`
	ConferenceId string                    `json:"ConferenceId" xml:"ConferenceId"`
	Participants ParticipantsInUnmuteAudio `json:"Participants" xml:"Participants"`
}

// CreateUnmuteAudioRequest creates a request to invoke UnmuteAudio API
func CreateUnmuteAudioRequest() (request *UnmuteAudioRequest) {
	request = &UnmuteAudioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "UnmuteAudio", "rtc", "openAPI")
	return
}

// CreateUnmuteAudioResponse creates a response to parse from UnmuteAudio response
func CreateUnmuteAudioResponse() (response *UnmuteAudioResponse) {
	response = &UnmuteAudioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}