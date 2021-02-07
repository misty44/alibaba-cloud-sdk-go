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

// StateConfigurationItemInUpdateStateConfiguration is a nested struct in oos response
type StateConfigurationItemInUpdateStateConfiguration struct {
	CreateTime           string                 `json:"CreateTime" xml:"CreateTime"`
	Description          string                 `json:"Description" xml:"Description"`
	StateConfigurationId string                 `json:"StateConfigurationId" xml:"StateConfigurationId"`
	TemplateId           string                 `json:"TemplateId" xml:"TemplateId"`
	TemplateName         string                 `json:"TemplateName" xml:"TemplateName"`
	TemplateVersion      string                 `json:"TemplateVersion" xml:"TemplateVersion"`
	Parameters           string                 `json:"Parameters" xml:"Parameters"`
	ConfigureMode        string                 `json:"ConfigureMode" xml:"ConfigureMode"`
	ScheduleType         string                 `json:"ScheduleType" xml:"ScheduleType"`
	ScheduleExpression   string                 `json:"ScheduleExpression" xml:"ScheduleExpression"`
	Targets              string                 `json:"Targets" xml:"Targets"`
	UpdateTime           string                 `json:"UpdateTime" xml:"UpdateTime"`
	Tags                 map[string]interface{} `json:"Tags" xml:"Tags"`
}
