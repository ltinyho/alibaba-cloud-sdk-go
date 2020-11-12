package openanalytics_open

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

// Table is a nested struct in openanalytics_open response
type Table struct {
	TableName         string                 `json:"TableName" xml:"TableName"`
	ViewOriginalText  string                 `json:"ViewOriginalText" xml:"ViewOriginalText"`
	Owner             string                 `json:"Owner" xml:"Owner"`
	TableType         string                 `json:"TableType" xml:"TableType"`
	Parameters        map[string]interface{} `json:"Parameters" xml:"Parameters"`
	ViewExpandedText  string                 `json:"ViewExpandedText" xml:"ViewExpandedText"`
	CreateTime        int64                  `json:"CreateTime" xml:"CreateTime"`
	DbName            string                 `json:"DbName" xml:"DbName"`
	LastAccessTime    int64                  `json:"LastAccessTime" xml:"LastAccessTime"`
	StorageDescriptor StorageDescriptor      `json:"StorageDescriptor" xml:"StorageDescriptor"`
	PartitionKeys     []PartitionKeysItem    `json:"PartitionKeys" xml:"PartitionKeys"`
}