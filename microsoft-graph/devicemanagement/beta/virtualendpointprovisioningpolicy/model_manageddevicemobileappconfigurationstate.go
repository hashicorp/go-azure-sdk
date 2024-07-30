package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationState struct {
	DisplayName       *string                                               `json:"displayName,omitempty"`
	Id                *string                                               `json:"id,omitempty"`
	ODataType         *string                                               `json:"@odata.type,omitempty"`
	PlatformType      *ManagedDeviceMobileAppConfigurationStatePlatformType `json:"platformType,omitempty"`
	SettingCount      *int64                                                `json:"settingCount,omitempty"`
	SettingStates     *[]ManagedDeviceMobileAppConfigurationSettingState    `json:"settingStates,omitempty"`
	State             *ManagedDeviceMobileAppConfigurationStateState        `json:"state,omitempty"`
	UserId            *string                                               `json:"userId,omitempty"`
	UserPrincipalName *string                                               `json:"userPrincipalName,omitempty"`
	Version           *int64                                                `json:"version,omitempty"`
}
