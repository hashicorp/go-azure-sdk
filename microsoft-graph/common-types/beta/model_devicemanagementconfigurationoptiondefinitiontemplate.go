package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationOptionDefinitionTemplate struct {
	Children  *[]DeviceManagementConfigurationSettingInstanceTemplate `json:"children,omitempty"`
	ItemId    *string                                                 `json:"itemId,omitempty"`
	ODataType *string                                                 `json:"@odata.type,omitempty"`
}
