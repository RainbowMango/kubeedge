/*
Copyright 2023 The KubeEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import "encoding/json"

// DeviceProfile is structure to store in configMap.
type DeviceProfile struct {
	DeviceInstances []DeviceInstance `json:"deviceInstances,omitempty"`
	DeviceModels    []DeviceModel    `json:"deviceModels,omitempty"`
	Protocols       []ProtocolConfig `json:"protocols,omitempty"`
}

// DeviceInstance is structure to store device in deviceProfile.json in configmap.
type DeviceInstance struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	ProtocolName string `json:"protocol,omitempty"`
	PProtocol    ProtocolConfig
	Model        string           `json:"model,omitempty"`
	Twins        []Twin           `json:"twins,omitempty"`
	Properties   []DeviceProperty `json:"properties,omitempty"`
}

// DeviceModel is structure to store deviceModel in deviceProfile.json in configmap.
type DeviceModel struct {
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Properties  []ModelProperty `json:"properties,omitempty"`
}

// ModelProperty is structure to store deviceModel property.
type ModelProperty struct {
	Name        string `json:"name,omitempty"`
	DataType    string `json:"dataType,omitempty"`
	Description string `json:"description,omitempty"`
	AccessMode  string `json:"accessMode,omitempty"`
	//DefaultValue interface{} `json:"defaultValue,omitempty"`
	Minimum string `json:"minimum,omitempty"` //todo todo why the type is int64
	Maximum string `json:"maximum,omitempty"`
	Unit    string `json:"unit,omitempty"`
}

// Protocol is structure to store protocol in deviceProfile.json in configmap.

type ProtocolConfig struct {
	// Unique protocol name
	// Required.
	ProtocolName string `json:"protocolName,omitempty"`
	// Any config data
	// +optional
	// +kubebuilder:validation:XPreserveUnknownFields
	ConfigData json.RawMessage `json:"configData,omitempty"`
}

// DeviceProperty is structure to store propertyVisitor in deviceProfile.json in configmap.
type DeviceProperty struct {
	Name         string          `json:"name,omitempty"`
	PropertyName string          `json:"propertyName,omitempty"`
	ModelName    string          `json:"modelName,omitempty"`
	Protocol     string          `json:"protocol,omitempty"`
	Visitors     json.RawMessage `json:"visitorConfig"`

	// whether be reported to the cloud
	ReportToCloud bool             `json:"reportToCloud,omitempty"`
	CollectCycle  int64            `json:"collectCycle"`
	ReportCycle   int64            `json:"reportCycle,omitempty"`
	PushMethod    PushMethodConfig `json:"pushMethod,omitempty"`
	DBProvider    DBProviderConfig `json:"dbProvider,omitempty"`

	PProperty ModelProperty
}

// PushMethodConfig is structure to store push config
type PushMethodConfig struct {
	MethodName   string          `json:"MethodName"`
	MethodConfig json.RawMessage `json:"MethodConfig"`
}

type DBProviderConfig struct {
	DBProviderName string         `json:"dbProviderName"`
	ProviderConfig ProviderConfig `json:"providerConfig"`
}

type ProviderConfig struct {
	ConfigData      json.RawMessage `json:"configData"`
	DataStandard    json.RawMessage `json:"dataStandard"`
	RedisConfigData json.RawMessage `json:" redisConfigData"`
}

// Metadata is the metadata for data.
type Metadata struct {
	Timestamp string `json:"timestamp,omitempty"`
	Type      string `json:"type,omitempty"`
}

// Twin is the set/get pair to one register.
type Twin struct {
	PropertyName    string `json:"propertyName,omitempty"`
	Property        *DeviceProperty
	ObservedDesired TwinProperty `json:"observedDesired,omitempty"`
	Reported        TwinProperty `json:"reported,omitempty"`
}

//// DesiredData is the desired data.
//type DesiredData struct {
//	Value     string   `json:"value,omitempty"`
//	Metadatas Metadata `json:"metadata,omitempty"`
//}
//
//// ReportedData is the reported data.
//type ReportedData struct {
//	Value     string   `json:"value,omitempty"`
//	Metadatas Metadata `json:"metadata,omitempty"`
//}

type TwinProperty struct {
	// Required: The value for this property.
	Value string `json:"value,"`
	// Additional metadata like timestamp when the value was reported etc.
	// +optional
	Metadata Metadata `json:"metadata,omitempty"`
}
