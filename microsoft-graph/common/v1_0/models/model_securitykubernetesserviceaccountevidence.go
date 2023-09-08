package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceAccountEvidence struct {
	CreatedDateTime          *string                                                    `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                                  `json:"detailedRoles,omitempty"`
	Name                     *string                                                    `json:"name,omitempty"`
	Namespace                *SecurityKubernetesNamespaceEvidence                       `json:"namespace,omitempty"`
	ODataType                *string                                                    `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityKubernetesServiceAccountEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                                    `json:"remediationStatusDetails,omitempty"`
	Roles                    *[]SecurityKubernetesServiceAccountEvidenceRoles           `json:"roles,omitempty"`
	Tags                     *[]string                                                  `json:"tags,omitempty"`
	Verdict                  *SecurityKubernetesServiceAccountEvidenceVerdict           `json:"verdict,omitempty"`
}
