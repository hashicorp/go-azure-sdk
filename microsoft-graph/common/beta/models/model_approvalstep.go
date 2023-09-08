package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalStep struct {
	AssignedToMe     *bool     `json:"assignedToMe,omitempty"`
	DisplayName      *string   `json:"displayName,omitempty"`
	Id               *string   `json:"id,omitempty"`
	Justification    *string   `json:"justification,omitempty"`
	ODataType        *string   `json:"@odata.type,omitempty"`
	ReviewResult     *string   `json:"reviewResult,omitempty"`
	ReviewedBy       *Identity `json:"reviewedBy,omitempty"`
	ReviewedDateTime *string   `json:"reviewedDateTime,omitempty"`
	Status           *string   `json:"status,omitempty"`
}
