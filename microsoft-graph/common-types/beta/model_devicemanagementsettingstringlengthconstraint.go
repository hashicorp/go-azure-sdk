package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingStringLengthConstraint struct {
	MaximumLength *int64  `json:"maximumLength,omitempty"`
	MinimumLength *int64  `json:"minimumLength,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
}
