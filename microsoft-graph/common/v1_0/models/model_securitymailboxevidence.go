package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailboxEvidence struct {
	CreatedDateTime          *string                                   `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                 `json:"detailedRoles,omitempty"`
	DisplayName              *string                                   `json:"displayName,omitempty"`
	ODataType                *string                                   `json:"@odata.type,omitempty"`
	PrimaryAddress           *string                                   `json:"primaryAddress,omitempty"`
	RemediationStatus        *SecurityMailboxEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                   `json:"remediationStatusDetails,omitempty"`
	Roles                    *[]SecurityMailboxEvidenceRoles           `json:"roles,omitempty"`
	Tags                     *[]string                                 `json:"tags,omitempty"`
	UserAccount              *SecurityUserAccount                      `json:"userAccount,omitempty"`
	Verdict                  *SecurityMailboxEvidenceVerdict           `json:"verdict,omitempty"`
}
