package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesControllerEvidence struct {
	CreatedDateTime          *string                                                `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                              `json:"detailedRoles,omitempty"`
	Labels                   *SecurityDictionary                                    `json:"labels,omitempty"`
	Name                     *string                                                `json:"name,omitempty"`
	Namespace                *SecurityKubernetesNamespaceEvidence                   `json:"namespace,omitempty"`
	ODataType                *string                                                `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityKubernetesControllerEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                                `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityKubernetesControllerEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                                              `json:"tags,omitempty"`
	Type                     *string                                                `json:"type,omitempty"`
	Verdict                  *SecurityKubernetesControllerEvidenceVerdict           `json:"verdict,omitempty"`
}
