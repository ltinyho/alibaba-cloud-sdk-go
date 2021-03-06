package green

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

// Bill is a nested struct in green response
type Bill struct {
	Region       string `json:"Region" xml:"Region"`
	Scene        string `json:"Scene" xml:"Scene"`
	TotalCount   int64  `json:"TotalCount" xml:"TotalCount"`
	ConfirmCount int64  `json:"ConfirmCount" xml:"ConfirmCount"`
	ReviewCount  int64  `json:"ReviewCount" xml:"ReviewCount"`
	FreeCount    int64  `json:"FreeCount" xml:"FreeCount"`
	SubUid       string `json:"SubUid" xml:"SubUid"`
	BizType      string `json:"BizType" xml:"BizType"`
	Day          string `json:"Day" xml:"Day"`
}
