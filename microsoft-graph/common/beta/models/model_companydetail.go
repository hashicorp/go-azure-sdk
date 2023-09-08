package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyDetail struct {
	Address        *PhysicalAddress `json:"address,omitempty"`
	Department     *string          `json:"department,omitempty"`
	DisplayName    *string          `json:"displayName,omitempty"`
	ODataType      *string          `json:"@odata.type,omitempty"`
	OfficeLocation *string          `json:"officeLocation,omitempty"`
	Pronunciation  *string          `json:"pronunciation,omitempty"`
	WebUrl         *string          `json:"webUrl,omitempty"`
}
