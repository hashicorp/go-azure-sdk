package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationState struct {
	DisplayName   *string                               `json:"displayName,omitempty"`
	Id            *string                               `json:"id,omitempty"`
	ODataType     *string                               `json:"@odata.type,omitempty"`
	PlatformType  *DeviceConfigurationStatePlatformType `json:"platformType,omitempty"`
	SettingCount  *int64                                `json:"settingCount,omitempty"`
	SettingStates *[]DeviceConfigurationSettingState    `json:"settingStates,omitempty"`
	State         *DeviceConfigurationStateState        `json:"state,omitempty"`
	Version       *int64                                `json:"version,omitempty"`
}
