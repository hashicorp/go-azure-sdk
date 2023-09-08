package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionContentLabel struct {
	AssignmentMethod *InformationProtectionContentLabelAssignmentMethod `json:"assignmentMethod,omitempty"`
	CreationDateTime *string                                            `json:"creationDateTime,omitempty"`
	Label            *LabelDetails                                      `json:"label,omitempty"`
	ODataType        *string                                            `json:"@odata.type,omitempty"`
}
