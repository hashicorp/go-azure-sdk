package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicySettingMapping struct {
	AdmxSettingDefinitionId   *string                                     `json:"admxSettingDefinitionId,omitempty"`
	ChildIdList               *[]string                                   `json:"childIdList,omitempty"`
	Id                        *string                                     `json:"id,omitempty"`
	IntuneSettingDefinitionId *string                                     `json:"intuneSettingDefinitionId,omitempty"`
	IntuneSettingUriList      *[]string                                   `json:"intuneSettingUriList,omitempty"`
	IsMdmSupported            *bool                                       `json:"isMdmSupported,omitempty"`
	MdmCspName                *string                                     `json:"mdmCspName,omitempty"`
	MdmMinimumOSVersion       *int64                                      `json:"mdmMinimumOSVersion,omitempty"`
	MdmSettingUri             *string                                     `json:"mdmSettingUri,omitempty"`
	MdmSupportedState         *GroupPolicySettingMappingMdmSupportedState `json:"mdmSupportedState,omitempty"`
	ODataType                 *string                                     `json:"@odata.type,omitempty"`
	ParentId                  *string                                     `json:"parentId,omitempty"`
	SettingCategory           *string                                     `json:"settingCategory,omitempty"`
	SettingDisplayName        *string                                     `json:"settingDisplayName,omitempty"`
	SettingDisplayValue       *string                                     `json:"settingDisplayValue,omitempty"`
	SettingDisplayValueType   *string                                     `json:"settingDisplayValueType,omitempty"`
	SettingName               *string                                     `json:"settingName,omitempty"`
	SettingScope              *GroupPolicySettingMappingSettingScope      `json:"settingScope,omitempty"`
	SettingType               *GroupPolicySettingMappingSettingType       `json:"settingType,omitempty"`
	SettingValue              *string                                     `json:"settingValue,omitempty"`
	SettingValueDisplayUnits  *string                                     `json:"settingValueDisplayUnits,omitempty"`
	SettingValueType          *string                                     `json:"settingValueType,omitempty"`
}
