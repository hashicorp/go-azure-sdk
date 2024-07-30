package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationSettingState struct {
	CurrentValue        *string                               `json:"currentValue,omitempty"`
	ErrorCode           *int64                                `json:"errorCode,omitempty"`
	ErrorDescription    *string                               `json:"errorDescription,omitempty"`
	InstanceDisplayName *string                               `json:"instanceDisplayName,omitempty"`
	ODataType           *string                               `json:"@odata.type,omitempty"`
	Setting             *string                               `json:"setting,omitempty"`
	SettingName         *string                               `json:"settingName,omitempty"`
	Sources             *[]SettingSource                      `json:"sources,omitempty"`
	State               *DeviceConfigurationSettingStateState `json:"state,omitempty"`
	UserEmail           *string                               `json:"userEmail,omitempty"`
	UserId              *string                               `json:"userId,omitempty"`
	UserName            *string                               `json:"userName,omitempty"`
	UserPrincipalName   *string                               `json:"userPrincipalName,omitempty"`
}
