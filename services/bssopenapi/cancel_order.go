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

// CancelOrder invokes the bssopenapi.CancelOrder API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/cancelorder.html
func (client *Client) CancelOrder(request *CancelOrderRequest) (response *CancelOrderResponse, err error) {
	response = CreateCancelOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CancelOrderWithChan invokes the bssopenapi.CancelOrder API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/cancelorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelOrderWithChan(request *CancelOrderRequest) (<-chan *CancelOrderResponse, <-chan error) {
	responseChan := make(chan *CancelOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelOrder(request)
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

// CancelOrderWithCallback invokes the bssopenapi.CancelOrder API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/cancelorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelOrderWithCallback(request *CancelOrderRequest, callback func(response *CancelOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelOrderResponse
		var err error
		defer close(result)
		response, err = client.CancelOrder(request)
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

// CancelOrderRequest is the request struct for api CancelOrder
type CancelOrderRequest struct {
	*requests.RpcRequest
	OrderId string           `position:"Query" name:"OrderId"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// CancelOrderResponse is the response struct for api CancelOrder
type CancelOrderResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCancelOrderRequest creates a request to invoke CancelOrder API
func CreateCancelOrderRequest() (request *CancelOrderRequest) {
	request = &CancelOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "CancelOrder", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelOrderResponse creates a response to parse from CancelOrder response
func CreateCancelOrderResponse() (response *CancelOrderResponse) {
	response = &CancelOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
