package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileAssessmentRequest struct {
	Category           *FileAssessmentRequestCategory           `json:"category,omitempty"`
	ContentData        *string                                  `json:"contentData,omitempty"`
	ContentType        *FileAssessmentRequestContentType        `json:"contentType,omitempty"`
	CreatedBy          *IdentitySet                             `json:"createdBy,omitempty"`
	CreatedDateTime    *string                                  `json:"createdDateTime,omitempty"`
	ExpectedAssessment *FileAssessmentRequestExpectedAssessment `json:"expectedAssessment,omitempty"`
	FileName           *string                                  `json:"fileName,omitempty"`
	Id                 *string                                  `json:"id,omitempty"`
	ODataType          *string                                  `json:"@odata.type,omitempty"`
	RequestSource      *FileAssessmentRequestRequestSource      `json:"requestSource,omitempty"`
	Results            *[]ThreatAssessmentResult                `json:"results,omitempty"`
	Status             *FileAssessmentRequestStatus             `json:"status,omitempty"`
}
