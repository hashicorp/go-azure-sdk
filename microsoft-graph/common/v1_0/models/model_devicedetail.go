package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceDetail struct {
	Browser         *string `json:"browser,omitempty"`
	DeviceId        *string `json:"deviceId,omitempty"`
	DisplayName     *string `json:"displayName,omitempty"`
	IsCompliant     *bool   `json:"isCompliant,omitempty"`
	IsManaged       *bool   `json:"isManaged,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	OperatingSystem *string `json:"operatingSystem,omitempty"`
	TrustType       *string `json:"trustType,omitempty"`
}
