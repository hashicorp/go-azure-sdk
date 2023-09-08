package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RubricQuality struct {
	Criteria    *[]RubricCriterion `json:"criteria,omitempty"`
	Description *EducationItemBody `json:"description,omitempty"`
	DisplayName *string            `json:"displayName,omitempty"`
	ODataType   *string            `json:"@odata.type,omitempty"`
	QualityId   *string            `json:"qualityId,omitempty"`
}
