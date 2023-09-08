package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationExternalResource struct {
	CreatedBy            *IdentitySet `json:"createdBy,omitempty"`
	CreatedDateTime      *string      `json:"createdDateTime,omitempty"`
	DisplayName          *string      `json:"displayName,omitempty"`
	LastModifiedBy       *IdentitySet `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string      `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string      `json:"@odata.type,omitempty"`
	WebUrl               *string      `json:"webUrl,omitempty"`
}
