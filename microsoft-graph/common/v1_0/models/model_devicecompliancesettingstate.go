package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceSettingState struct {
	ComplianceGracePeriodExpirationDateTime *string                            `json:"complianceGracePeriodExpirationDateTime,omitempty"`
	DeviceId                                *string                            `json:"deviceId,omitempty"`
	DeviceModel                             *string                            `json:"deviceModel,omitempty"`
	DeviceName                              *string                            `json:"deviceName,omitempty"`
	Id                                      *string                            `json:"id,omitempty"`
	ODataType                               *string                            `json:"@odata.type,omitempty"`
	Setting                                 *string                            `json:"setting,omitempty"`
	SettingName                             *string                            `json:"settingName,omitempty"`
	State                                   *DeviceComplianceSettingStateState `json:"state,omitempty"`
	UserEmail                               *string                            `json:"userEmail,omitempty"`
	UserId                                  *string                            `json:"userId,omitempty"`
	UserName                                *string                            `json:"userName,omitempty"`
	UserPrincipalName                       *string                            `json:"userPrincipalName,omitempty"`
}
