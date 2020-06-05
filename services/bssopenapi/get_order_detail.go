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

// GetOrderDetail invokes the bssopenapi.GetOrderDetail API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/getorderdetail.html
func (client *Client) GetOrderDetail(request *GetOrderDetailRequest) (response *GetOrderDetailResponse, err error) {
	response = CreateGetOrderDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetOrderDetailWithChan invokes the bssopenapi.GetOrderDetail API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/getorderdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetOrderDetailWithChan(request *GetOrderDetailRequest) (<-chan *GetOrderDetailResponse, <-chan error) {
	responseChan := make(chan *GetOrderDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOrderDetail(request)
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

// GetOrderDetailWithCallback invokes the bssopenapi.GetOrderDetail API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/getorderdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetOrderDetailWithCallback(request *GetOrderDetailRequest, callback func(response *GetOrderDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOrderDetailResponse
		var err error
		defer close(result)
		response, err = client.GetOrderDetail(request)
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

// GetOrderDetailRequest is the request struct for api GetOrderDetail
type GetOrderDetailRequest struct {
	*requests.RpcRequest
	OrderId string           `position:"Query" name:"OrderId"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// GetOrderDetailResponse is the response struct for api GetOrderDetail
type GetOrderDetailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetOrderDetailRequest creates a request to invoke GetOrderDetail API
func CreateGetOrderDetailRequest() (request *GetOrderDetailRequest) {
	request = &GetOrderDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "GetOrderDetail", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetOrderDetailResponse creates a response to parse from GetOrderDetail response
func CreateGetOrderDetailResponse() (response *GetOrderDetailResponse) {
	response = &GetOrderDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
