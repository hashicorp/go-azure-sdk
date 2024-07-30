package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionMailEvidence struct {
	CreatedDateTime          *string                                          `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                        `json:"detailedRoles,omitempty"`
	NetworkMessageId         *string                                          `json:"networkMessageId,omitempty"`
	ODataType                *string                                          `json:"@odata.type,omitempty"`
	Recipient                *string                                          `json:"recipient,omitempty"`
	RemediationStatus        *SecuritySubmissionMailEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                          `json:"remediationStatusDetails,omitempty"`
	ReportType               *string                                          `json:"reportType,omitempty"`
	Roles                    *SecuritySubmissionMailEvidenceRoles             `json:"roles,omitempty"`
	Sender                   *string                                          `json:"sender,omitempty"`
	SenderIp                 *string                                          `json:"senderIp,omitempty"`
	Subject                  *string                                          `json:"subject,omitempty"`
	SubmissionDateTime       *string                                          `json:"submissionDateTime,omitempty"`
	SubmissionId             *string                                          `json:"submissionId,omitempty"`
	Submitter                *string                                          `json:"submitter,omitempty"`
	Tags                     *[]string                                        `json:"tags,omitempty"`
	Verdict                  *SecuritySubmissionMailEvidenceVerdict           `json:"verdict,omitempty"`
}
