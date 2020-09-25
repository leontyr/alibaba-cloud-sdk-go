package sas

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

// HandleSecurityEvents invokes the sas.HandleSecurityEvents API synchronously
func (client *Client) HandleSecurityEvents(request *HandleSecurityEventsRequest) (response *HandleSecurityEventsResponse, err error) {
	response = CreateHandleSecurityEventsResponse()
	err = client.DoAction(request, response)
	return
}

// HandleSecurityEventsWithChan invokes the sas.HandleSecurityEvents API asynchronously
func (client *Client) HandleSecurityEventsWithChan(request *HandleSecurityEventsRequest) (<-chan *HandleSecurityEventsResponse, <-chan error) {
	responseChan := make(chan *HandleSecurityEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.HandleSecurityEvents(request)
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

// HandleSecurityEventsWithCallback invokes the sas.HandleSecurityEvents API asynchronously
func (client *Client) HandleSecurityEventsWithCallback(request *HandleSecurityEventsRequest, callback func(response *HandleSecurityEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *HandleSecurityEventsResponse
		var err error
		defer close(result)
		response, err = client.HandleSecurityEvents(request)
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

// HandleSecurityEventsRequest is the request struct for api HandleSecurityEvents
type HandleSecurityEventsRequest struct {
	*requests.RpcRequest
	MarkMissParam    string           `position:"Query" name:"MarkMissParam"`
	ResourceOwnerId  requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityEventIds *[]string        `position:"Query" name:"SecurityEventIds"  type:"Repeated"`
	SourceIp         string           `position:"Query" name:"SourceIp"`
	OperationCode    string           `position:"Query" name:"OperationCode"`
	OperationParams  string           `position:"Query" name:"OperationParams"`
}

// HandleSecurityEventsResponse is the response struct for api HandleSecurityEvents
type HandleSecurityEventsResponse struct {
	*responses.BaseResponse
	RequestId                    string                       `json:"RequestId" xml:"RequestId"`
	HandleSecurityEventsResponse HandleSecurityEventsResponse `json:"HandleSecurityEventsResponse" xml:"HandleSecurityEventsResponse"`
}

// CreateHandleSecurityEventsRequest creates a request to invoke HandleSecurityEvents API
func CreateHandleSecurityEventsRequest() (request *HandleSecurityEventsRequest) {
	request = &HandleSecurityEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "HandleSecurityEvents", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateHandleSecurityEventsResponse creates a response to parse from HandleSecurityEvents response
func CreateHandleSecurityEventsResponse() (response *HandleSecurityEventsResponse) {
	response = &HandleSecurityEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}