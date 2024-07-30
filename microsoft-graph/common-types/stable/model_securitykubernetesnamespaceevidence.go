package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesNamespaceEvidence struct {
	Cluster                  *SecurityKubernetesClusterEvidence                    `json:"cluster,omitempty"`
	CreatedDateTime          *string                                               `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                             `json:"detailedRoles,omitempty"`
	Labels                   *SecurityDictionary                                   `json:"labels,omitempty"`
	Name                     *string                                               `json:"name,omitempty"`
	ODataType                *string                                               `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityKubernetesNamespaceEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                               `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityKubernetesNamespaceEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                                             `json:"tags,omitempty"`
	Verdict                  *SecurityKubernetesNamespaceEvidenceVerdict           `json:"verdict,omitempty"`
}
