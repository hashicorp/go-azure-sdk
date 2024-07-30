package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGitHubOrganizationEvidence struct {
	Company                  *string                                              `json:"company,omitempty"`
	CreatedDateTime          *string                                              `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                            `json:"detailedRoles,omitempty"`
	DisplayName              *string                                              `json:"displayName,omitempty"`
	Email                    *string                                              `json:"email,omitempty"`
	Login                    *string                                              `json:"login,omitempty"`
	ODataType                *string                                              `json:"@odata.type,omitempty"`
	OrgId                    *string                                              `json:"orgId,omitempty"`
	RemediationStatus        *SecurityGitHubOrganizationEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                              `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityGitHubOrganizationEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                                            `json:"tags,omitempty"`
	Verdict                  *SecurityGitHubOrganizationEvidenceVerdict           `json:"verdict,omitempty"`
	WebUrl                   *string                                              `json:"webUrl,omitempty"`
}
