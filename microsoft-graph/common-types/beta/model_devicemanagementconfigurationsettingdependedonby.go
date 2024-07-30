package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingDependedOnBy struct {
	DependedOnBy *string `json:"dependedOnBy,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	Required     *bool   `json:"required,omitempty"`
}
