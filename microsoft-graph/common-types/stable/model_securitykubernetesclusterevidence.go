package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesClusterEvidence struct {
	CloudResource            *SecurityAlertEvidence                              `json:"cloudResource,omitempty"`
	CreatedDateTime          *string                                             `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                           `json:"detailedRoles,omitempty"`
	Distribution             *string                                             `json:"distribution,omitempty"`
	Name                     *string                                             `json:"name,omitempty"`
	ODataType                *string                                             `json:"@odata.type,omitempty"`
	Platform                 *SecurityKubernetesClusterEvidencePlatform          `json:"platform,omitempty"`
	RemediationStatus        *SecurityKubernetesClusterEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                             `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityKubernetesClusterEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                                           `json:"tags,omitempty"`
	Verdict                  *SecurityKubernetesClusterEvidenceVerdict           `json:"verdict,omitempty"`
	Version                  *string                                             `json:"version,omitempty"`
}
