package dbs

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

// BackupPlanDetail is a nested struct in dbs response
type BackupPlanDetail struct {
	BackupPlanId                      string `json:"BackupPlanId" xml:"BackupPlanId"`
	SourceEndpointInstanceType        string `json:"SourceEndpointInstanceType" xml:"SourceEndpointInstanceType"`
	SourceEndpointRegion              string `json:"SourceEndpointRegion" xml:"SourceEndpointRegion"`
	SourceEndpointInstanceID          string `json:"SourceEndpointInstanceID" xml:"SourceEndpointInstanceID"`
	SourceEndpointIpPort              string `json:"SourceEndpointIpPort" xml:"SourceEndpointIpPort"`
	SourceEndpointDatabaseName        string `json:"SourceEndpointDatabaseName" xml:"SourceEndpointDatabaseName"`
	SourceEndpointUserName            string `json:"SourceEndpointUserName" xml:"SourceEndpointUserName"`
	BackupObjects                     string `json:"BackupObjects" xml:"BackupObjects"`
	BackupGatewayId                   int    `json:"BackupGatewayId" xml:"BackupGatewayId"`
	OSSBucketRegion                   string `json:"OSSBucketRegion" xml:"OSSBucketRegion"`
	OSSBucketName                     string `json:"OSSBucketName" xml:"OSSBucketName"`
	BackupPeriod                      string `json:"BackupPeriod" xml:"BackupPeriod"`
	BackupStartTime                   string `json:"BackupStartTime" xml:"BackupStartTime"`
	EnableBackupLog                   bool   `json:"EnableBackupLog" xml:"EnableBackupLog"`
	BackupRetentionPeriod             int    `json:"BackupRetentionPeriod" xml:"BackupRetentionPeriod"`
	DuplicationInfrequentAccessPeriod int    `json:"DuplicationInfrequentAccessPeriod" xml:"DuplicationInfrequentAccessPeriod"`
	DuplicationArchivePeriod          int    `json:"DuplicationArchivePeriod" xml:"DuplicationArchivePeriod"`
	BackupPlanName                    string `json:"BackupPlanName" xml:"BackupPlanName"`
	SourceEndpointOracleSID           string `json:"SourceEndpointOracleSID" xml:"SourceEndpointOracleSID"`
	InstanceClass                     string `json:"InstanceClass" xml:"InstanceClass"`
	BackupMethod                      string `json:"BackupMethod" xml:"BackupMethod"`
	BackupPlanCreateTime              int    `json:"BackupPlanCreateTime" xml:"BackupPlanCreateTime"`
	BackupPlanStatus                  string `json:"BackupPlanStatus" xml:"BackupPlanStatus"`
}
