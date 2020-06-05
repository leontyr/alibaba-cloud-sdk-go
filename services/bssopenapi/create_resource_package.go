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

// CreateResourcePackage invokes the bssopenapi.CreateResourcePackage API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/createresourcepackage.html
func (client *Client) CreateResourcePackage(request *CreateResourcePackageRequest) (response *CreateResourcePackageResponse, err error) {
	response = CreateCreateResourcePackageResponse()
	err = client.DoAction(request, response)
	return
}

// CreateResourcePackageWithChan invokes the bssopenapi.CreateResourcePackage API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/createresourcepackage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateResourcePackageWithChan(request *CreateResourcePackageRequest) (<-chan *CreateResourcePackageResponse, <-chan error) {
	responseChan := make(chan *CreateResourcePackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateResourcePackage(request)
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

// CreateResourcePackageWithCallback invokes the bssopenapi.CreateResourcePackage API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/createresourcepackage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateResourcePackageWithCallback(request *CreateResourcePackageRequest, callback func(response *CreateResourcePackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateResourcePackageResponse
		var err error
		defer close(result)
		response, err = client.CreateResourcePackage(request)
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

// CreateResourcePackageRequest is the request struct for api CreateResourcePackage
type CreateResourcePackageRequest struct {
	*requests.RpcRequest
	ProductCode   string           `position:"Query" name:"ProductCode"`
	Specification string           `position:"Query" name:"Specification"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	EffectiveDate string           `position:"Query" name:"EffectiveDate"`
	Duration      requests.Integer `position:"Query" name:"Duration"`
	PackageType   string           `position:"Query" name:"PackageType"`
	PricingCycle  string           `position:"Query" name:"PricingCycle"`
}

// CreateResourcePackageResponse is the response struct for api CreateResourcePackage
type CreateResourcePackageResponse struct {
	*responses.BaseResponse
	RequestId string                      `json:"RequestId" xml:"RequestId"`
	OrderId   int64                       `json:"OrderId" xml:"OrderId"`
	Success   bool                        `json:"Success" xml:"Success"`
	Code      string                      `json:"Code" xml:"Code"`
	Message   string                      `json:"Message" xml:"Message"`
	Data      DataInCreateResourcePackage `json:"Data" xml:"Data"`
}

// CreateCreateResourcePackageRequest creates a request to invoke CreateResourcePackage API
func CreateCreateResourcePackageRequest() (request *CreateResourcePackageRequest) {
	request = &CreateResourcePackageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "CreateResourcePackage", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateResourcePackageResponse creates a response to parse from CreateResourcePackage response
func CreateCreateResourcePackageResponse() (response *CreateResourcePackageResponse) {
	response = &CreateResourcePackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
