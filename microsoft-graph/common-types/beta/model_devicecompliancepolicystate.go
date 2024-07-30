package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyState struct {
	DisplayName       *string                                  `json:"displayName,omitempty"`
	Id                *string                                  `json:"id,omitempty"`
	ODataType         *string                                  `json:"@odata.type,omitempty"`
	PlatformType      *DeviceCompliancePolicyStatePlatformType `json:"platformType,omitempty"`
	SettingCount      *int64                                   `json:"settingCount,omitempty"`
	SettingStates     *[]DeviceCompliancePolicySettingState    `json:"settingStates,omitempty"`
	State             *DeviceCompliancePolicyStateState        `json:"state,omitempty"`
	UserId            *string                                  `json:"userId,omitempty"`
	UserPrincipalName *string                                  `json:"userPrincipalName,omitempty"`
	Version           *int64                                   `json:"version,omitempty"`
}
