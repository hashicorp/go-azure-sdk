package securityinformationprotectionsensitivitylabel

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabelingOptions struct {
	AssignmentMethod       *LabelingOptionsAssignmentMethod `json:"assignmentMethod,omitempty"`
	DowngradeJustification *DowngradeJustification          `json:"downgradeJustification,omitempty"`
	ExtendedProperties     *[]KeyValuePair                  `json:"extendedProperties,omitempty"`
	LabelId                *string                          `json:"labelId,omitempty"`
	ODataType              *string                          `json:"@odata.type,omitempty"`
}
