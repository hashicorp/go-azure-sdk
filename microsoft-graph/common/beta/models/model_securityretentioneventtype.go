package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionEventType struct {
	CreatedBy            *IdentitySet `json:"createdBy,omitempty"`
	CreatedDateTime      *string      `json:"createdDateTime,omitempty"`
	Description          *string      `json:"description,omitempty"`
	DisplayName          *string      `json:"displayName,omitempty"`
	Id                   *string      `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string      `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string      `json:"@odata.type,omitempty"`
}
