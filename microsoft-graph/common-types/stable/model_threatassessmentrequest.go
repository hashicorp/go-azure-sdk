package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentRequest struct {
	Category           *ThreatAssessmentRequestCategory           `json:"category,omitempty"`
	ContentType        *ThreatAssessmentRequestContentType        `json:"contentType,omitempty"`
	CreatedBy          *IdentitySet                               `json:"createdBy,omitempty"`
	CreatedDateTime    *string                                    `json:"createdDateTime,omitempty"`
	ExpectedAssessment *ThreatAssessmentRequestExpectedAssessment `json:"expectedAssessment,omitempty"`
	Id                 *string                                    `json:"id,omitempty"`
	ODataType          *string                                    `json:"@odata.type,omitempty"`
	RequestSource      *ThreatAssessmentRequestRequestSource      `json:"requestSource,omitempty"`
	Results            *[]ThreatAssessmentResult                  `json:"results,omitempty"`
	Status             *ThreatAssessmentRequestStatus             `json:"status,omitempty"`
}
