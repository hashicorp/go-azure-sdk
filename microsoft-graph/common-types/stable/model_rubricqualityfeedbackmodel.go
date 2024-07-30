package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RubricQualityFeedbackModel struct {
	Feedback  *EducationItemBody `json:"feedback,omitempty"`
	ODataType *string            `json:"@odata.type,omitempty"`
	QualityId *string            `json:"qualityId,omitempty"`
}
