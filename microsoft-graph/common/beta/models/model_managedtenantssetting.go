package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsSetting struct {
	DisplayName      *string                         `json:"displayName,omitempty"`
	JsonValue        *string                         `json:"jsonValue,omitempty"`
	ODataType        *string                         `json:"@odata.type,omitempty"`
	OverwriteAllowed *bool                           `json:"overwriteAllowed,omitempty"`
	SettingId        *string                         `json:"settingId,omitempty"`
	ValueType        *ManagedTenantsSettingValueType `json:"valueType,omitempty"`
}
