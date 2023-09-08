package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureScoreControlStateUpdate struct {
	AssignedTo      *string `json:"assignedTo,omitempty"`
	Comment         *string `json:"comment,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	State           *string `json:"state,omitempty"`
	UpdatedBy       *string `json:"updatedBy,omitempty"`
	UpdatedDateTime *string `json:"updatedDateTime,omitempty"`
}
