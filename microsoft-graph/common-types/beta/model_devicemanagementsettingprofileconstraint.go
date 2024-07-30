package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingProfileConstraint struct {
	ODataType *string   `json:"@odata.type,omitempty"`
	Source    *string   `json:"source,omitempty"`
	Types     *[]string `json:"types,omitempty"`
}
