package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityLabelingOptions struct {
	AssignmentMethod       *SecurityLabelingOptionsAssignmentMethod `json:"assignmentMethod,omitempty"`
	DowngradeJustification *SecurityDowngradeJustification          `json:"downgradeJustification,omitempty"`
	ExtendedProperties     *[]SecurityKeyValuePair                  `json:"extendedProperties,omitempty"`
	LabelId                *string                                  `json:"labelId,omitempty"`
	ODataType              *string                                  `json:"@odata.type,omitempty"`
}
