package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationConflictSummary struct {
	ConflictingDeviceConfigurations *[]SettingSource `json:"conflictingDeviceConfigurations,omitempty"`
	ContributingSettings            *[]string        `json:"contributingSettings,omitempty"`
	DeviceCheckinsImpacted          *int64           `json:"deviceCheckinsImpacted,omitempty"`
	Id                              *string          `json:"id,omitempty"`
	ODataType                       *string          `json:"@odata.type,omitempty"`
}
