package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSubmission struct {
	Id                  *string                        `json:"id,omitempty"`
	ODataType           *string                        `json:"@odata.type,omitempty"`
	Outcomes            *[]EducationOutcome            `json:"outcomes,omitempty"`
	ReassignedBy        *IdentitySet                   `json:"reassignedBy,omitempty"`
	ReassignedDateTime  *string                        `json:"reassignedDateTime,omitempty"`
	Recipient           *EducationSubmissionRecipient  `json:"recipient,omitempty"`
	Resources           *[]EducationSubmissionResource `json:"resources,omitempty"`
	ResourcesFolderUrl  *string                        `json:"resourcesFolderUrl,omitempty"`
	ReturnedBy          *IdentitySet                   `json:"returnedBy,omitempty"`
	ReturnedDateTime    *string                        `json:"returnedDateTime,omitempty"`
	Status              *EducationSubmissionStatus     `json:"status,omitempty"`
	SubmittedBy         *IdentitySet                   `json:"submittedBy,omitempty"`
	SubmittedDateTime   *string                        `json:"submittedDateTime,omitempty"`
	SubmittedResources  *[]EducationSubmissionResource `json:"submittedResources,omitempty"`
	UnsubmittedBy       *IdentitySet                   `json:"unsubmittedBy,omitempty"`
	UnsubmittedDateTime *string                        `json:"unsubmittedDateTime,omitempty"`
	WebUrl              *string                        `json:"webUrl,omitempty"`
}
