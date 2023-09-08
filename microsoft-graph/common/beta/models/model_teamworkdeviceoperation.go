package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceOperation struct {
	CompletedDateTime  *string                               `json:"completedDateTime,omitempty"`
	CreatedBy          *IdentitySet                          `json:"createdBy,omitempty"`
	CreatedDateTime    *string                               `json:"createdDateTime,omitempty"`
	Error              *OperationError                       `json:"error,omitempty"`
	Id                 *string                               `json:"id,omitempty"`
	LastActionBy       *IdentitySet                          `json:"lastActionBy,omitempty"`
	LastActionDateTime *string                               `json:"lastActionDateTime,omitempty"`
	ODataType          *string                               `json:"@odata.type,omitempty"`
	OperationType      *TeamworkDeviceOperationOperationType `json:"operationType,omitempty"`
	StartedDateTime    *string                               `json:"startedDateTime,omitempty"`
	Status             *string                               `json:"status,omitempty"`
}
