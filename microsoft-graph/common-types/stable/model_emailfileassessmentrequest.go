package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailFileAssessmentRequest struct {
	Category                 *EmailFileAssessmentRequestCategory                 `json:"category,omitempty"`
	ContentData              *string                                             `json:"contentData,omitempty"`
	ContentType              *EmailFileAssessmentRequestContentType              `json:"contentType,omitempty"`
	CreatedBy                *IdentitySet                                        `json:"createdBy,omitempty"`
	CreatedDateTime          *string                                             `json:"createdDateTime,omitempty"`
	DestinationRoutingReason *EmailFileAssessmentRequestDestinationRoutingReason `json:"destinationRoutingReason,omitempty"`
	ExpectedAssessment       *EmailFileAssessmentRequestExpectedAssessment       `json:"expectedAssessment,omitempty"`
	Id                       *string                                             `json:"id,omitempty"`
	ODataType                *string                                             `json:"@odata.type,omitempty"`
	RecipientEmail           *string                                             `json:"recipientEmail,omitempty"`
	RequestSource            *EmailFileAssessmentRequestRequestSource            `json:"requestSource,omitempty"`
	Results                  *[]ThreatAssessmentResult                           `json:"results,omitempty"`
	Status                   *EmailFileAssessmentRequestStatus                   `json:"status,omitempty"`
}
