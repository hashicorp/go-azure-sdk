package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementReusablePolicySetting struct {
	CreatedDateTime                     *string                                       `json:"createdDateTime,omitempty"`
	Description                         *string                                       `json:"description,omitempty"`
	DisplayName                         *string                                       `json:"displayName,omitempty"`
	Id                                  *string                                       `json:"id,omitempty"`
	LastModifiedDateTime                *string                                       `json:"lastModifiedDateTime,omitempty"`
	ODataType                           *string                                       `json:"@odata.type,omitempty"`
	ReferencingConfigurationPolicies    *[]DeviceManagementConfigurationPolicy        `json:"referencingConfigurationPolicies,omitempty"`
	ReferencingConfigurationPolicyCount *int64                                        `json:"referencingConfigurationPolicyCount,omitempty"`
	SettingDefinitionId                 *string                                       `json:"settingDefinitionId,omitempty"`
	SettingInstance                     *DeviceManagementConfigurationSettingInstance `json:"settingInstance,omitempty"`
	Version                             *int64                                        `json:"version,omitempty"`
}
