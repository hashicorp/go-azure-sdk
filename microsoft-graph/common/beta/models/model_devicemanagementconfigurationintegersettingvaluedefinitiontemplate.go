package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate struct {
	MaxValue  *int64  `json:"maxValue,omitempty"`
	MinValue  *int64  `json:"minValue,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
