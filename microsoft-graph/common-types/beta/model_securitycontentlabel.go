package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContentLabel struct {
	AssignmentMethod   *SecurityContentLabelAssignmentMethod `json:"assignmentMethod,omitempty"`
	CreatedDateTime    *string                               `json:"createdDateTime,omitempty"`
	ODataType          *string                               `json:"@odata.type,omitempty"`
	SensitivityLabelId *string                               `json:"sensitivityLabelId,omitempty"`
}
