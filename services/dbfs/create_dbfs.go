package dbfs

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

// CreateDbfs invokes the dbfs.CreateDbfs API synchronously
// api document: https://help.aliyun.com/api/dbfs/createdbfs.html
func (client *Client) CreateDbfs(request *CreateDbfsRequest) (response *CreateDbfsResponse, err error) {
	response = CreateCreateDbfsResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDbfsWithChan invokes the dbfs.CreateDbfs API asynchronously
// api document: https://help.aliyun.com/api/dbfs/createdbfs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDbfsWithChan(request *CreateDbfsRequest) (<-chan *CreateDbfsResponse, <-chan error) {
	responseChan := make(chan *CreateDbfsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDbfs(request)
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

// CreateDbfsWithCallback invokes the dbfs.CreateDbfs API asynchronously
// api document: https://help.aliyun.com/api/dbfs/createdbfs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDbfsWithCallback(request *CreateDbfsRequest, callback func(response *CreateDbfsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDbfsResponse
		var err error
		defer close(result)
		response, err = client.CreateDbfs(request)
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

// CreateDbfsRequest is the request struct for api CreateDbfs
type CreateDbfsRequest struct {
	*requests.RpcRequest
	SizeG       requests.Integer `position:"Query" name:"SizeG"`
	ClientToken string           `position:"Query" name:"ClientToken"`
	FsName      string           `position:"Query" name:"FsName"`
	ZoneId      string           `position:"Query" name:"ZoneId"`
	Category    string           `position:"Query" name:"Category"`
}

// CreateDbfsResponse is the response struct for api CreateDbfs
type CreateDbfsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	FsId      string `json:"FsId" xml:"FsId"`
}

// CreateCreateDbfsRequest creates a request to invoke CreateDbfs API
func CreateCreateDbfsRequest() (request *CreateDbfsRequest) {
	request = &CreateDbfsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-02-19", "CreateDbfs", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDbfsResponse creates a response to parse from CreateDbfs response
func CreateCreateDbfsResponse() (response *CreateDbfsResponse) {
	response = &CreateDbfsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}