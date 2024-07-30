package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGitHubRepoEvidence struct {
	BaseUrl                  *string                                      `json:"baseUrl,omitempty"`
	CreatedDateTime          *string                                      `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                    `json:"detailedRoles,omitempty"`
	Login                    *string                                      `json:"login,omitempty"`
	ODataType                *string                                      `json:"@odata.type,omitempty"`
	Owner                    *string                                      `json:"owner,omitempty"`
	OwnerType                *string                                      `json:"ownerType,omitempty"`
	RemediationStatus        *SecurityGitHubRepoEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                      `json:"remediationStatusDetails,omitempty"`
	RepoId                   *string                                      `json:"repoId,omitempty"`
	Roles                    *SecurityGitHubRepoEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                                    `json:"tags,omitempty"`
	Verdict                  *SecurityGitHubRepoEvidenceVerdict           `json:"verdict,omitempty"`
}
