package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesPodEvidence struct {
	Containers               *[]SecurityContainerEvidence                    `json:"containers,omitempty"`
	Controller               *SecurityKubernetesControllerEvidence           `json:"controller,omitempty"`
	CreatedDateTime          *string                                         `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                       `json:"detailedRoles,omitempty"`
	EphemeralContainers      *[]SecurityContainerEvidence                    `json:"ephemeralContainers,omitempty"`
	InitContainers           *[]SecurityContainerEvidence                    `json:"initContainers,omitempty"`
	Labels                   *SecurityDictionary                             `json:"labels,omitempty"`
	Name                     *string                                         `json:"name,omitempty"`
	Namespace                *SecurityKubernetesNamespaceEvidence            `json:"namespace,omitempty"`
	ODataType                *string                                         `json:"@odata.type,omitempty"`
	PodIp                    *SecurityIpEvidence                             `json:"podIp,omitempty"`
	RemediationStatus        *SecurityKubernetesPodEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                         `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityKubernetesPodEvidenceRoles             `json:"roles,omitempty"`
	ServiceAccount           *SecurityKubernetesServiceAccountEvidence       `json:"serviceAccount,omitempty"`
	Tags                     *[]string                                       `json:"tags,omitempty"`
	Verdict                  *SecurityKubernetesPodEvidenceVerdict           `json:"verdict,omitempty"`
}
