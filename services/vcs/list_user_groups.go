package vcs

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

// ListUserGroups invokes the vcs.ListUserGroups API synchronously
func (client *Client) ListUserGroups(request *ListUserGroupsRequest) (response *ListUserGroupsResponse, err error) {
	response = CreateListUserGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserGroupsWithChan invokes the vcs.ListUserGroups API asynchronously
func (client *Client) ListUserGroupsWithChan(request *ListUserGroupsRequest) (<-chan *ListUserGroupsResponse, <-chan error) {
	responseChan := make(chan *ListUserGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserGroups(request)
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

// ListUserGroupsWithCallback invokes the vcs.ListUserGroups API asynchronously
func (client *Client) ListUserGroupsWithCallback(request *ListUserGroupsRequest, callback func(response *ListUserGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListUserGroups(request)
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

// ListUserGroupsRequest is the request struct for api ListUserGroups
type ListUserGroupsRequest struct {
	*requests.RpcRequest
	IsvSubId string `position:"Query" name:"IsvSubId"`
	CorpId   string `position:"Query" name:"CorpId"`
}

// ListUserGroupsResponse is the response struct for api ListUserGroups
type ListUserGroupsResponse struct {
	*responses.BaseResponse
	Code      string                     `json:"Code" xml:"Code"`
	Message   string                     `json:"Message" xml:"Message"`
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	Data      []DataItemInListUserGroups `json:"Data" xml:"Data"`
}

// CreateListUserGroupsRequest creates a request to invoke ListUserGroups API
func CreateListUserGroupsRequest() (request *ListUserGroupsRequest) {
	request = &ListUserGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "ListUserGroups", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListUserGroupsResponse creates a response to parse from ListUserGroups response
func CreateListUserGroupsResponse() (response *ListUserGroupsResponse) {
	response = &ListUserGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}