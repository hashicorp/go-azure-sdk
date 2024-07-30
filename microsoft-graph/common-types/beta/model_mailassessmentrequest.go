package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailAssessmentRequest struct {
	Category                 *MailAssessmentRequestCategory                 `json:"category,omitempty"`
	ContentType              *MailAssessmentRequestContentType              `json:"contentType,omitempty"`
	CreatedBy                *IdentitySet                                   `json:"createdBy,omitempty"`
	CreatedDateTime          *string                                        `json:"createdDateTime,omitempty"`
	DestinationRoutingReason *MailAssessmentRequestDestinationRoutingReason `json:"destinationRoutingReason,omitempty"`
	ExpectedAssessment       *MailAssessmentRequestExpectedAssessment       `json:"expectedAssessment,omitempty"`
	Id                       *string                                        `json:"id,omitempty"`
	MessageUri               *string                                        `json:"messageUri,omitempty"`
	ODataType                *string                                        `json:"@odata.type,omitempty"`
	RecipientEmail           *string                                        `json:"recipientEmail,omitempty"`
	RequestSource            *MailAssessmentRequestRequestSource            `json:"requestSource,omitempty"`
	Results                  *[]ThreatAssessmentResult                      `json:"results,omitempty"`
	Status                   *MailAssessmentRequestStatus                   `json:"status,omitempty"`
}
