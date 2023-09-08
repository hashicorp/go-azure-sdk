package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingRequiredConstraint struct {
	NotConfiguredValue *string `json:"notConfiguredValue,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
}
