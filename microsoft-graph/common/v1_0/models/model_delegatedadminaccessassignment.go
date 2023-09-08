package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminAccessAssignment struct {
	AccessContainer      *DelegatedAdminAccessContainer        `json:"accessContainer,omitempty"`
	AccessDetails        *DelegatedAdminAccessDetails          `json:"accessDetails,omitempty"`
	CreatedDateTime      *string                               `json:"createdDateTime,omitempty"`
	Id                   *string                               `json:"id,omitempty"`
	LastModifiedDateTime *string                               `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                               `json:"@odata.type,omitempty"`
	Status               *DelegatedAdminAccessAssignmentStatus `json:"status,omitempty"`
}
