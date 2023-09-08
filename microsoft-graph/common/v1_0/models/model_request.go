package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Request struct {
	ApprovalId        *string      `json:"approvalId,omitempty"`
	CompletedDateTime *string      `json:"completedDateTime,omitempty"`
	CreatedBy         *IdentitySet `json:"createdBy,omitempty"`
	CreatedDateTime   *string      `json:"createdDateTime,omitempty"`
	CustomData        *string      `json:"customData,omitempty"`
	Id                *string      `json:"id,omitempty"`
	ODataType         *string      `json:"@odata.type,omitempty"`
	Status            *string      `json:"status,omitempty"`
}
