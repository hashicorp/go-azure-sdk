package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Approval struct {
	Id        *string          `json:"id,omitempty"`
	ODataType *string          `json:"@odata.type,omitempty"`
	Stages    *[]ApprovalStage `json:"stages,omitempty"`
}
