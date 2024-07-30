package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationshipOperation struct {
	CreatedDateTime      *string                                           `json:"createdDateTime,omitempty"`
	Data                 *string                                           `json:"data,omitempty"`
	Id                   *string                                           `json:"id,omitempty"`
	LastModifiedDateTime *string                                           `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                           `json:"@odata.type,omitempty"`
	OperationType        *DelegatedAdminRelationshipOperationOperationType `json:"operationType,omitempty"`
	Status               *DelegatedAdminRelationshipOperationStatus        `json:"status,omitempty"`
}
