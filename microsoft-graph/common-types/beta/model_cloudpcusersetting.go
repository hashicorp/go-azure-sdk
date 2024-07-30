package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCUserSetting struct {
	Assignments          *[]CloudPCUserSettingAssignment `json:"assignments,omitempty"`
	CreatedDateTime      *string                         `json:"createdDateTime,omitempty"`
	DisplayName          *string                         `json:"displayName,omitempty"`
	Id                   *string                         `json:"id,omitempty"`
	LastModifiedDateTime *string                         `json:"lastModifiedDateTime,omitempty"`
	LocalAdminEnabled    *bool                           `json:"localAdminEnabled,omitempty"`
	ODataType            *string                         `json:"@odata.type,omitempty"`
	ResetEnabled         *bool                           `json:"resetEnabled,omitempty"`
	RestorePointSetting  *CloudPCRestorePointSetting     `json:"restorePointSetting,omitempty"`
	SelfServiceEnabled   *bool                           `json:"selfServiceEnabled,omitempty"`
}
