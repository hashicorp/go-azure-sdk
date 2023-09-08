package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstitutionData struct {
	Description *string          `json:"description,omitempty"`
	DisplayName *string          `json:"displayName,omitempty"`
	Location    *PhysicalAddress `json:"location,omitempty"`
	ODataType   *string          `json:"@odata.type,omitempty"`
	WebUrl      *string          `json:"webUrl,omitempty"`
}
