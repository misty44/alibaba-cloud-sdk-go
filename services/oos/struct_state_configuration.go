package oos

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

// StateConfiguration is a nested struct in oos response
type StateConfiguration struct {
	TemplateId           string                 `json:"TemplateId" xml:"TemplateId"`
	ConfigureMode        string                 `json:"ConfigureMode" xml:"ConfigureMode"`
	TemplateName         string                 `json:"TemplateName" xml:"TemplateName"`
	UpdateTime           string                 `json:"UpdateTime" xml:"UpdateTime"`
	CreateTime           string                 `json:"CreateTime" xml:"CreateTime"`
	ScheduleType         string                 `json:"ScheduleType" xml:"ScheduleType"`
	StateConfigurationId string                 `json:"StateConfigurationId" xml:"StateConfigurationId"`
	Targets              string                 `json:"Targets" xml:"Targets"`
	TemplateVersion      string                 `json:"TemplateVersion" xml:"TemplateVersion"`
	Parameters           string                 `json:"Parameters" xml:"Parameters"`
	ScheduleExpression   string                 `json:"ScheduleExpression" xml:"ScheduleExpression"`
	Description          string                 `json:"Description" xml:"Description"`
	Tags                 map[string]interface{} `json:"Tags" xml:"Tags"`
}
