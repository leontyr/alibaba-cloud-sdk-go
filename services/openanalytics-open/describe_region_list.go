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

// DescribeRegionList invokes the openanalytics_open.DescribeRegionList API synchronously
func (client *Client) DescribeRegionList(request *DescribeRegionListRequest) (response *DescribeRegionListResponse, err error) {
	response = CreateDescribeRegionListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRegionListWithChan invokes the openanalytics_open.DescribeRegionList API asynchronously
func (client *Client) DescribeRegionListWithChan(request *DescribeRegionListRequest) (<-chan *DescribeRegionListResponse, <-chan error) {
	responseChan := make(chan *DescribeRegionListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRegionList(request)
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

// DescribeRegionListWithCallback invokes the openanalytics_open.DescribeRegionList API asynchronously
func (client *Client) DescribeRegionListWithCallback(request *DescribeRegionListRequest, callback func(response *DescribeRegionListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRegionListResponse
		var err error
		defer close(result)
		response, err = client.DescribeRegionList(request)
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

// DescribeRegionListRequest is the request struct for api DescribeRegionList
type DescribeRegionListRequest struct {
	*requests.RpcRequest
}

// DescribeRegionListResponse is the response struct for api DescribeRegionList
type DescribeRegionListResponse struct {
	*responses.BaseResponse
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	RegionId   string           `json:"RegionId" xml:"RegionId"`
	RegionList []RegionListItem `json:"RegionList" xml:"RegionList"`
}

// CreateDescribeRegionListRequest creates a request to invoke DescribeRegionList API
func CreateDescribeRegionListRequest() (request *DescribeRegionListRequest) {
	request = &DescribeRegionListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "DescribeRegionList", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRegionListResponse creates a response to parse from DescribeRegionList response
func CreateDescribeRegionListResponse() (response *DescribeRegionListResponse) {
	response = &DescribeRegionListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
