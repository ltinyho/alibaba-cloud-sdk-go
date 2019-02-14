package cloudapi

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

// DescribePluginsByApi invokes the cloudapi.DescribePluginsByApi API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describepluginsbyapi.html
func (client *Client) DescribePluginsByApi(request *DescribePluginsByApiRequest) (response *DescribePluginsByApiResponse, err error) {
	response = CreateDescribePluginsByApiResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePluginsByApiWithChan invokes the cloudapi.DescribePluginsByApi API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describepluginsbyapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePluginsByApiWithChan(request *DescribePluginsByApiRequest) (<-chan *DescribePluginsByApiResponse, <-chan error) {
	responseChan := make(chan *DescribePluginsByApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePluginsByApi(request)
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

// DescribePluginsByApiWithCallback invokes the cloudapi.DescribePluginsByApi API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describepluginsbyapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePluginsByApiWithCallback(request *DescribePluginsByApiRequest, callback func(response *DescribePluginsByApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePluginsByApiResponse
		var err error
		defer close(result)
		response, err = client.DescribePluginsByApi(request)
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

// DescribePluginsByApiRequest is the request struct for api DescribePluginsByApi
type DescribePluginsByApiRequest struct {
	*requests.RpcRequest
	StageName     string `position:"Query" name:"StageName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GroupId       string `position:"Query" name:"GroupId"`
	ApiId         string `position:"Query" name:"ApiId"`
}

// DescribePluginsByApiResponse is the response struct for api DescribePluginsByApi
type DescribePluginsByApiResponse struct {
	*responses.BaseResponse
	RequestId  string                        `json:"RequestId" xml:"RequestId"`
	TotalCount int                           `json:"TotalCount" xml:"TotalCount"`
	PageSize   int                           `json:"PageSize" xml:"PageSize"`
	PageNumber int                           `json:"PageNumber" xml:"PageNumber"`
	Plugins    PluginsInDescribePluginsByApi `json:"Plugins" xml:"Plugins"`
}

// CreateDescribePluginsByApiRequest creates a request to invoke DescribePluginsByApi API
func CreateDescribePluginsByApiRequest() (request *DescribePluginsByApiRequest) {
	request = &DescribePluginsByApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribePluginsByApi", "apigateway", "openAPI")
	return
}

// CreateDescribePluginsByApiResponse creates a response to parse from DescribePluginsByApi response
func CreateDescribePluginsByApiResponse() (response *DescribePluginsByApiResponse) {
	response = &DescribePluginsByApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}