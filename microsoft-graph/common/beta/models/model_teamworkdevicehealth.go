package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceHealth struct {
	Connection           *TeamworkConnection           `json:"connection,omitempty"`
	CreatedBy            *IdentitySet                  `json:"createdBy,omitempty"`
	CreatedDateTime      *string                       `json:"createdDateTime,omitempty"`
	HardwareHealth       *TeamworkHardwareHealth       `json:"hardwareHealth,omitempty"`
	Id                   *string                       `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet                  `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                       `json:"lastModifiedDateTime,omitempty"`
	LoginStatus          *TeamworkLoginStatus          `json:"loginStatus,omitempty"`
	ODataType            *string                       `json:"@odata.type,omitempty"`
	PeripheralsHealth    *TeamworkPeripheralsHealth    `json:"peripheralsHealth,omitempty"`
	SoftwareUpdateHealth *TeamworkSoftwareUpdateHealth `json:"softwareUpdateHealth,omitempty"`
}
