package codeup

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

// CreateMergeRequest invokes the codeup.CreateMergeRequest API synchronously
func (client *Client) CreateMergeRequest(request *CreateMergeRequestRequest) (response *CreateMergeRequestResponse, err error) {
	response = CreateCreateMergeRequestResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMergeRequestWithChan invokes the codeup.CreateMergeRequest API asynchronously
func (client *Client) CreateMergeRequestWithChan(request *CreateMergeRequestRequest) (<-chan *CreateMergeRequestResponse, <-chan error) {
	responseChan := make(chan *CreateMergeRequestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMergeRequest(request)
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

// CreateMergeRequestWithCallback invokes the codeup.CreateMergeRequest API asynchronously
func (client *Client) CreateMergeRequestWithCallback(request *CreateMergeRequestRequest, callback func(response *CreateMergeRequestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMergeRequestResponse
		var err error
		defer close(result)
		response, err = client.CreateMergeRequest(request)
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

// CreateMergeRequestRequest is the request struct for api CreateMergeRequest
type CreateMergeRequestRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	SubUserId      string           `position:"Query" name:"SubUserId"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
	ProjectId      requests.Integer `position:"Path" name:"ProjectId"`
}

// CreateMergeRequestResponse is the response struct for api CreateMergeRequest
type CreateMergeRequestResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateCreateMergeRequestRequest creates a request to invoke CreateMergeRequest API
func CreateCreateMergeRequestRequest() (request *CreateMergeRequestRequest) {
	request = &CreateMergeRequestRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "CreateMergeRequest", "/api/v4/projects/[ProjectId]/merge_requests", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateMergeRequestResponse creates a response to parse from CreateMergeRequest response
func CreateCreateMergeRequestResponse() (response *CreateMergeRequestResponse) {
	response = &CreateMergeRequestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}