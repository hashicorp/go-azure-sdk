package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationRubricOutcome struct {
	Id                                   *string                             `json:"id,omitempty"`
	LastModifiedBy                       *IdentitySet                        `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime                 *string                             `json:"lastModifiedDateTime,omitempty"`
	ODataType                            *string                             `json:"@odata.type,omitempty"`
	PublishedRubricQualityFeedback       *[]RubricQualityFeedbackModel       `json:"publishedRubricQualityFeedback,omitempty"`
	PublishedRubricQualitySelectedLevels *[]RubricQualitySelectedColumnModel `json:"publishedRubricQualitySelectedLevels,omitempty"`
	RubricQualityFeedback                *[]RubricQualityFeedbackModel       `json:"rubricQualityFeedback,omitempty"`
	RubricQualitySelectedLevels          *[]RubricQualitySelectedColumnModel `json:"rubricQualitySelectedLevels,omitempty"`
}
