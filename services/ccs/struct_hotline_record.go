package ccs

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

// HotlineRecord is a nested struct in ccs response
type HotlineRecord struct {
	Id              string `json:"Id" xml:"Id"`
	VisitorId       string `json:"VisitorId" xml:"VisitorId"`
	VisitorPhone    string `json:"VisitorPhone" xml:"VisitorPhone"`
	VisitorProvince string `json:"VisitorProvince" xml:"VisitorProvince"`
	CallType        string `json:"CallType" xml:"CallType"`
	AgentId         string `json:"AgentId" xml:"AgentId"`
	AgentName       string `json:"AgentName" xml:"AgentName"`
	GroupId         string `json:"GroupId" xml:"GroupId"`
	GroupName       string `json:"GroupName" xml:"GroupName"`
	CreateTime      string `json:"CreateTime" xml:"CreateTime"`
	FinishTime      string `json:"FinishTime" xml:"FinishTime"`
	Status          string `json:"Status" xml:"Status"`
	Memo            string `json:"Memo" xml:"Memo"`
	HangupType      string `json:"HangupType" xml:"HangupType"`
	Satisfaction    string `json:"Satisfaction" xml:"Satisfaction"`
	OutboundTaskId  string `json:"OutboundTaskId" xml:"OutboundTaskId"`
	Categories      string `json:"Categories" xml:"Categories"`
	CcsInstanceId   string `json:"CcsInstanceId" xml:"CcsInstanceId"`
	TalkDuration    int    `json:"TalkDuration" xml:"TalkDuration"`
}
