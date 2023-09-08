package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceEvidence struct {
	ClusterIP                *SecurityIpEvidence                                 `json:"clusterIP,omitempty"`
	CreatedDateTime          *string                                             `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                           `json:"detailedRoles,omitempty"`
	ExternalIPs              *[]SecurityIpEvidence                               `json:"externalIPs,omitempty"`
	Labels                   *SecurityDictionary                                 `json:"labels,omitempty"`
	Name                     *string                                             `json:"name,omitempty"`
	Namespace                *SecurityKubernetesNamespaceEvidence                `json:"namespace,omitempty"`
	ODataType                *string                                             `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityKubernetesServiceEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                             `json:"remediationStatusDetails,omitempty"`
	Roles                    *[]SecurityKubernetesServiceEvidenceRoles           `json:"roles,omitempty"`
	Selector                 *SecurityDictionary                                 `json:"selector,omitempty"`
	ServicePorts             *[]SecurityKubernetesServicePort                    `json:"servicePorts,omitempty"`
	ServiceType              *SecurityKubernetesServiceEvidenceServiceType       `json:"serviceType,omitempty"`
	Tags                     *[]string                                           `json:"tags,omitempty"`
	Verdict                  *SecurityKubernetesServiceEvidenceVerdict           `json:"verdict,omitempty"`
}
