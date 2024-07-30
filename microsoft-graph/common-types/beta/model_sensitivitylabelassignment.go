package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabelAssignment struct {
	AssignmentMethod   *SensitivityLabelAssignmentAssignmentMethod `json:"assignmentMethod,omitempty"`
	ODataType          *string                                     `json:"@odata.type,omitempty"`
	SensitivityLabelId *string                                     `json:"sensitivityLabelId,omitempty"`
	TenantId           *string                                     `json:"tenantId,omitempty"`
}
