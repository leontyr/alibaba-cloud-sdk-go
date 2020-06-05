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

// QueryResourcePackageInstances invokes the bssopenapi.QueryResourcePackageInstances API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryresourcepackageinstances.html
func (client *Client) QueryResourcePackageInstances(request *QueryResourcePackageInstancesRequest) (response *QueryResourcePackageInstancesResponse, err error) {
	response = CreateQueryResourcePackageInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// QueryResourcePackageInstancesWithChan invokes the bssopenapi.QueryResourcePackageInstances API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryresourcepackageinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryResourcePackageInstancesWithChan(request *QueryResourcePackageInstancesRequest) (<-chan *QueryResourcePackageInstancesResponse, <-chan error) {
	responseChan := make(chan *QueryResourcePackageInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryResourcePackageInstances(request)
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

// QueryResourcePackageInstancesWithCallback invokes the bssopenapi.QueryResourcePackageInstances API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryresourcepackageinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryResourcePackageInstancesWithCallback(request *QueryResourcePackageInstancesRequest, callback func(response *QueryResourcePackageInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryResourcePackageInstancesResponse
		var err error
		defer close(result)
		response, err = client.QueryResourcePackageInstances(request)
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

// QueryResourcePackageInstancesRequest is the request struct for api QueryResourcePackageInstances
type QueryResourcePackageInstancesRequest struct {
	*requests.RpcRequest
	ExpiryTimeEnd   string           `position:"Query" name:"ExpiryTimeEnd"`
	ProductCode     string           `position:"Query" name:"ProductCode"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	ExpiryTimeStart string           `position:"Query" name:"ExpiryTimeStart"`
	PageNum         requests.Integer `position:"Query" name:"PageNum"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
}

// QueryResourcePackageInstancesResponse is the response struct for api QueryResourcePackageInstances
type QueryResourcePackageInstancesResponse struct {
	*responses.BaseResponse
	RequestId string                              `json:"RequestId" xml:"RequestId"`
	Success   bool                                `json:"Success" xml:"Success"`
	Code      string                              `json:"Code" xml:"Code"`
	Message   string                              `json:"Message" xml:"Message"`
	Page      int                                 `json:"Page" xml:"Page"`
	PageSize  int                                 `json:"PageSize" xml:"PageSize"`
	Total     int                                 `json:"Total" xml:"Total"`
	Data      DataInQueryResourcePackageInstances `json:"Data" xml:"Data"`
}

// CreateQueryResourcePackageInstancesRequest creates a request to invoke QueryResourcePackageInstances API
func CreateQueryResourcePackageInstancesRequest() (request *QueryResourcePackageInstancesRequest) {
	request = &QueryResourcePackageInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryResourcePackageInstances", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryResourcePackageInstancesResponse creates a response to parse from QueryResourcePackageInstances response
func CreateQueryResourcePackageInstancesResponse() (response *QueryResourcePackageInstancesResponse) {
	response = &QueryResourcePackageInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
