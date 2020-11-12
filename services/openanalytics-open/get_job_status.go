package openanalytics_open

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

// GetJobStatus invokes the openanalytics_open.GetJobStatus API synchronously
func (client *Client) GetJobStatus(request *GetJobStatusRequest) (response *GetJobStatusResponse, err error) {
	response = CreateGetJobStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetJobStatusWithChan invokes the openanalytics_open.GetJobStatus API asynchronously
func (client *Client) GetJobStatusWithChan(request *GetJobStatusRequest) (<-chan *GetJobStatusResponse, <-chan error) {
	responseChan := make(chan *GetJobStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetJobStatus(request)
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

// GetJobStatusWithCallback invokes the openanalytics_open.GetJobStatus API asynchronously
func (client *Client) GetJobStatusWithCallback(request *GetJobStatusRequest, callback func(response *GetJobStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetJobStatusResponse
		var err error
		defer close(result)
		response, err = client.GetJobStatus(request)
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

// GetJobStatusRequest is the request struct for api GetJobStatus
type GetJobStatusRequest struct {
	*requests.RpcRequest
	JobId  string `position:"Body" name:"JobId"`
	VcName string `position:"Body" name:"VcName"`
}

// GetJobStatusResponse is the response struct for api GetJobStatus
type GetJobStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateGetJobStatusRequest creates a request to invoke GetJobStatus API
func CreateGetJobStatusRequest() (request *GetJobStatusRequest) {
	request = &GetJobStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "GetJobStatus", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetJobStatusResponse creates a response to parse from GetJobStatus response
func CreateGetJobStatusResponse() (response *GetJobStatusResponse) {
	response = &GetJobStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
