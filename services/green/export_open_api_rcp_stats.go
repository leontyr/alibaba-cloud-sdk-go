package green

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

// ExportOpenApiRcpStats invokes the green.ExportOpenApiRcpStats API synchronously
func (client *Client) ExportOpenApiRcpStats(request *ExportOpenApiRcpStatsRequest) (response *ExportOpenApiRcpStatsResponse, err error) {
	response = CreateExportOpenApiRcpStatsResponse()
	err = client.DoAction(request, response)
	return
}

// ExportOpenApiRcpStatsWithChan invokes the green.ExportOpenApiRcpStats API asynchronously
func (client *Client) ExportOpenApiRcpStatsWithChan(request *ExportOpenApiRcpStatsRequest) (<-chan *ExportOpenApiRcpStatsResponse, <-chan error) {
	responseChan := make(chan *ExportOpenApiRcpStatsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportOpenApiRcpStats(request)
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

// ExportOpenApiRcpStatsWithCallback invokes the green.ExportOpenApiRcpStats API asynchronously
func (client *Client) ExportOpenApiRcpStatsWithCallback(request *ExportOpenApiRcpStatsRequest, callback func(response *ExportOpenApiRcpStatsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportOpenApiRcpStatsResponse
		var err error
		defer close(result)
		response, err = client.ExportOpenApiRcpStats(request)
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

// ExportOpenApiRcpStatsRequest is the request struct for api ExportOpenApiRcpStats
type ExportOpenApiRcpStatsRequest struct {
	*requests.RpcRequest
	StartDate    string `position:"Query" name:"StartDate"`
	ResourceType string `position:"Query" name:"ResourceType"`
	BizType      string `position:"Query" name:"BizType"`
	EndDate      string `position:"Query" name:"EndDate"`
	SourceIp     string `position:"Query" name:"SourceIp"`
}

// ExportOpenApiRcpStatsResponse is the response struct for api ExportOpenApiRcpStats
type ExportOpenApiRcpStatsResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	DownloadUrl string `json:"DownloadUrl" xml:"DownloadUrl"`
}

// CreateExportOpenApiRcpStatsRequest creates a request to invoke ExportOpenApiRcpStats API
func CreateExportOpenApiRcpStatsRequest() (request *ExportOpenApiRcpStatsRequest) {
	request = &ExportOpenApiRcpStatsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "ExportOpenApiRcpStats", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExportOpenApiRcpStatsResponse creates a response to parse from ExportOpenApiRcpStats response
func CreateExportOpenApiRcpStatsResponse() (response *ExportOpenApiRcpStatsResponse) {
	response = &ExportOpenApiRcpStatsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
