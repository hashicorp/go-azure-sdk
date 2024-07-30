package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingIntegerConstraint struct {
	MaximumValue *int64  `json:"maximumValue,omitempty"`
	MinimumValue *int64  `json:"minimumValue,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
}
