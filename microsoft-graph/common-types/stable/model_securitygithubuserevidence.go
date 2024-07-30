package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGitHubUserEvidence struct {
	CreatedDateTime          *string                                      `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                    `json:"detailedRoles,omitempty"`
	Email                    *string                                      `json:"email,omitempty"`
	Login                    *string                                      `json:"login,omitempty"`
	Name                     *string                                      `json:"name,omitempty"`
	ODataType                *string                                      `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityGitHubUserEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                      `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityGitHubUserEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                                    `json:"tags,omitempty"`
	UserId                   *string                                      `json:"userId,omitempty"`
	Verdict                  *SecurityGitHubUserEvidenceVerdict           `json:"verdict,omitempty"`
	WebUrl                   *string                                      `json:"webUrl,omitempty"`
}
