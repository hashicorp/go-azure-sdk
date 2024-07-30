package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProgramControl struct {
	ControlId       *string          `json:"controlId,omitempty"`
	ControlTypeId   *string          `json:"controlTypeId,omitempty"`
	CreatedDateTime *string          `json:"createdDateTime,omitempty"`
	DisplayName     *string          `json:"displayName,omitempty"`
	Id              *string          `json:"id,omitempty"`
	ODataType       *string          `json:"@odata.type,omitempty"`
	Owner           *UserIdentity    `json:"owner,omitempty"`
	Program         *Program         `json:"program,omitempty"`
	ProgramId       *string          `json:"programId,omitempty"`
	Resource        *ProgramResource `json:"resource,omitempty"`
	Status          *string          `json:"status,omitempty"`
}
