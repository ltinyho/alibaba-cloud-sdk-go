package vod

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

// SetMessageCallback invokes the vod.SetMessageCallback API synchronously
// api document: https://help.aliyun.com/api/vod/setmessagecallback.html
func (client *Client) SetMessageCallback(request *SetMessageCallbackRequest) (response *SetMessageCallbackResponse, err error) {
	response = CreateSetMessageCallbackResponse()
	err = client.DoAction(request, response)
	return
}

// SetMessageCallbackWithChan invokes the vod.SetMessageCallback API asynchronously
// api document: https://help.aliyun.com/api/vod/setmessagecallback.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetMessageCallbackWithChan(request *SetMessageCallbackRequest) (<-chan *SetMessageCallbackResponse, <-chan error) {
	responseChan := make(chan *SetMessageCallbackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetMessageCallback(request)
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

// SetMessageCallbackWithCallback invokes the vod.SetMessageCallback API asynchronously
// api document: https://help.aliyun.com/api/vod/setmessagecallback.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetMessageCallbackWithCallback(request *SetMessageCallbackRequest, callback func(response *SetMessageCallbackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetMessageCallbackResponse
		var err error
		defer close(result)
		response, err = client.SetMessageCallback(request)
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

// SetMessageCallbackRequest is the request struct for api SetMessageCallback
type SetMessageCallbackRequest struct {
	*requests.RpcRequest
	AuthKey              string           `position:"Query" name:"AuthKey"`
	ResourceOwnerId      string           `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EventTypeList        string           `position:"Query" name:"EventTypeList"`
	MnsQueueName         string           `position:"Query" name:"MnsQueueName"`
	ResourceRealOwnerId  requests.Integer `position:"Query" name:"ResourceRealOwnerId"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	CallbackType         string           `position:"Query" name:"CallbackType"`
	CallbackSwitch       string           `position:"Query" name:"CallbackSwitch"`
	MnsEndpoint          string           `position:"Query" name:"MnsEndpoint"`
	AppId                string           `position:"Query" name:"AppId"`
	AuthSwitch           string           `position:"Query" name:"AuthSwitch"`
	CallbackURL          string           `position:"Query" name:"CallbackURL"`
}

// SetMessageCallbackResponse is the response struct for api SetMessageCallback
type SetMessageCallbackResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetMessageCallbackRequest creates a request to invoke SetMessageCallback API
func CreateSetMessageCallbackRequest() (request *SetMessageCallbackRequest) {
	request = &SetMessageCallbackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SetMessageCallback", "vod", "openAPI")
	return
}

// CreateSetMessageCallbackResponse creates a response to parse from SetMessageCallback response
func CreateSetMessageCallbackResponse() (response *SetMessageCallbackResponse) {
	response = &SetMessageCallbackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}