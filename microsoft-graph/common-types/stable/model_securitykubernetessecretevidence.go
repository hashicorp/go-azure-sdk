package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesSecretEvidence struct {
	CreatedDateTime          *string                                            `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                          `json:"detailedRoles,omitempty"`
	Name                     *string                                            `json:"name,omitempty"`
	Namespace                *SecurityKubernetesNamespaceEvidence               `json:"namespace,omitempty"`
	ODataType                *string                                            `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityKubernetesSecretEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                            `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityKubernetesSecretEvidenceRoles             `json:"roles,omitempty"`
	SecretType               *string                                            `json:"secretType,omitempty"`
	Tags                     *[]string                                          `json:"tags,omitempty"`
	Verdict                  *SecurityKubernetesSecretEvidenceVerdict           `json:"verdict,omitempty"`
}
