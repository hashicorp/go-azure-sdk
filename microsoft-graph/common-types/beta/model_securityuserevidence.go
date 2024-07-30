package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserEvidence struct {
	CreatedDateTime          *string                                `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                              `json:"detailedRoles,omitempty"`
	ODataType                *string                                `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityUserEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityUserEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                              `json:"tags,omitempty"`
	UserAccount              *SecurityUserAccount                   `json:"userAccount,omitempty"`
	Verdict                  *SecurityUserEvidenceVerdict           `json:"verdict,omitempty"`
}
