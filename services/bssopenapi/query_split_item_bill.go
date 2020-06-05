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

// QuerySplitItemBill invokes the bssopenapi.QuerySplitItemBill API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/querysplititembill.html
func (client *Client) QuerySplitItemBill(request *QuerySplitItemBillRequest) (response *QuerySplitItemBillResponse, err error) {
	response = CreateQuerySplitItemBillResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySplitItemBillWithChan invokes the bssopenapi.QuerySplitItemBill API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querysplititembill.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QuerySplitItemBillWithChan(request *QuerySplitItemBillRequest) (<-chan *QuerySplitItemBillResponse, <-chan error) {
	responseChan := make(chan *QuerySplitItemBillResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySplitItemBill(request)
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

// QuerySplitItemBillWithCallback invokes the bssopenapi.QuerySplitItemBill API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querysplititembill.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QuerySplitItemBillWithCallback(request *QuerySplitItemBillRequest, callback func(response *QuerySplitItemBillResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySplitItemBillResponse
		var err error
		defer close(result)
		response, err = client.QuerySplitItemBill(request)
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

// QuerySplitItemBillRequest is the request struct for api QuerySplitItemBill
type QuerySplitItemBillRequest struct {
	*requests.RpcRequest
	ProductCode      string           `position:"Query" name:"ProductCode"`
	SubscriptionType string           `position:"Query" name:"SubscriptionType"`
	BillingCycle     string           `position:"Query" name:"BillingCycle"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	PageNum          requests.Integer `position:"Query" name:"PageNum"`
	ProductType      string           `position:"Query" name:"ProductType"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
}

// QuerySplitItemBillResponse is the response struct for api QuerySplitItemBill
type QuerySplitItemBillResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQuerySplitItemBillRequest creates a request to invoke QuerySplitItemBill API
func CreateQuerySplitItemBillRequest() (request *QuerySplitItemBillRequest) {
	request = &QuerySplitItemBillRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QuerySplitItemBill", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySplitItemBillResponse creates a response to parse from QuerySplitItemBill response
func CreateQuerySplitItemBillResponse() (response *QuerySplitItemBillResponse) {
	response = &QuerySplitItemBillResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
