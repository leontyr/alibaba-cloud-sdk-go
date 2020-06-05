package bssopenapi

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

// QueryUserOmsData invokes the bssopenapi.QueryUserOmsData API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryuseromsdata.html
func (client *Client) QueryUserOmsData(request *QueryUserOmsDataRequest) (response *QueryUserOmsDataResponse, err error) {
	response = CreateQueryUserOmsDataResponse()
	err = client.DoAction(request, response)
	return
}

// QueryUserOmsDataWithChan invokes the bssopenapi.QueryUserOmsData API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryuseromsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryUserOmsDataWithChan(request *QueryUserOmsDataRequest) (<-chan *QueryUserOmsDataResponse, <-chan error) {
	responseChan := make(chan *QueryUserOmsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryUserOmsData(request)
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

// QueryUserOmsDataWithCallback invokes the bssopenapi.QueryUserOmsData API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryuseromsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryUserOmsDataWithCallback(request *QueryUserOmsDataRequest, callback func(response *QueryUserOmsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryUserOmsDataResponse
		var err error
		defer close(result)
		response, err = client.QueryUserOmsData(request)
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

// QueryUserOmsDataRequest is the request struct for api QueryUserOmsData
type QueryUserOmsDataRequest struct {
	*requests.RpcRequest
	EndTime   string           `position:"Query" name:"EndTime"`
	StartTime string           `position:"Query" name:"StartTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	DataType  string           `position:"Query" name:"DataType"`
	Marker    string           `position:"Query" name:"Marker"`
	PageSize  requests.Integer `position:"Query" name:"PageSize"`
	Table     string           `position:"Query" name:"Table"`
}

// QueryUserOmsDataResponse is the response struct for api QueryUserOmsData
type QueryUserOmsDataResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Success   bool                   `json:"Success" xml:"Success"`
	Code      string                 `json:"Code" xml:"Code"`
	Message   string                 `json:"Message" xml:"Message"`
	Data      DataInQueryUserOmsData `json:"Data" xml:"Data"`
}

// CreateQueryUserOmsDataRequest creates a request to invoke QueryUserOmsData API
func CreateQueryUserOmsDataRequest() (request *QueryUserOmsDataRequest) {
	request = &QueryUserOmsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryUserOmsData", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryUserOmsDataResponse creates a response to parse from QueryUserOmsData response
func CreateQueryUserOmsDataResponse() (response *QueryUserOmsDataResponse) {
	response = &QueryUserOmsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
