package ccc

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

// ModifyPrimaryTrunksOfSkillGroup invokes the ccc.ModifyPrimaryTrunksOfSkillGroup API synchronously
func (client *Client) ModifyPrimaryTrunksOfSkillGroup(request *ModifyPrimaryTrunksOfSkillGroupRequest) (response *ModifyPrimaryTrunksOfSkillGroupResponse, err error) {
	response = CreateModifyPrimaryTrunksOfSkillGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPrimaryTrunksOfSkillGroupWithChan invokes the ccc.ModifyPrimaryTrunksOfSkillGroup API asynchronously
func (client *Client) ModifyPrimaryTrunksOfSkillGroupWithChan(request *ModifyPrimaryTrunksOfSkillGroupRequest) (<-chan *ModifyPrimaryTrunksOfSkillGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyPrimaryTrunksOfSkillGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPrimaryTrunksOfSkillGroup(request)
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

// ModifyPrimaryTrunksOfSkillGroupWithCallback invokes the ccc.ModifyPrimaryTrunksOfSkillGroup API asynchronously
func (client *Client) ModifyPrimaryTrunksOfSkillGroupWithCallback(request *ModifyPrimaryTrunksOfSkillGroupRequest, callback func(response *ModifyPrimaryTrunksOfSkillGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPrimaryTrunksOfSkillGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyPrimaryTrunksOfSkillGroup(request)
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

// ModifyPrimaryTrunksOfSkillGroupRequest is the request struct for api ModifyPrimaryTrunksOfSkillGroup
type ModifyPrimaryTrunksOfSkillGroupRequest struct {
	*requests.RpcRequest
	PrimaryProviderName *[]string `position:"Query" name:"PrimaryProviderName"  type:"Repeated"`
	InstanceId          string    `position:"Query" name:"InstanceId"`
	SkillGroupId        string    `position:"Query" name:"SkillGroupId"`
}

// ModifyPrimaryTrunksOfSkillGroupResponse is the response struct for api ModifyPrimaryTrunksOfSkillGroup
type ModifyPrimaryTrunksOfSkillGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateModifyPrimaryTrunksOfSkillGroupRequest creates a request to invoke ModifyPrimaryTrunksOfSkillGroup API
func CreateModifyPrimaryTrunksOfSkillGroupRequest() (request *ModifyPrimaryTrunksOfSkillGroupRequest) {
	request = &ModifyPrimaryTrunksOfSkillGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ModifyPrimaryTrunksOfSkillGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyPrimaryTrunksOfSkillGroupResponse creates a response to parse from ModifyPrimaryTrunksOfSkillGroup response
func CreateModifyPrimaryTrunksOfSkillGroupResponse() (response *ModifyPrimaryTrunksOfSkillGroupResponse) {
	response = &ModifyPrimaryTrunksOfSkillGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
