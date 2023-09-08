package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerEvidence struct {
	Args                     *[]string                                   `json:"args,omitempty"`
	Command                  *[]string                                   `json:"command,omitempty"`
	ContainerId              *string                                     `json:"containerId,omitempty"`
	CreatedDateTime          *string                                     `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                   `json:"detailedRoles,omitempty"`
	Image                    *SecurityContainerImageEvidence             `json:"image,omitempty"`
	IsPrivileged             *bool                                       `json:"isPrivileged,omitempty"`
	Name                     *string                                     `json:"name,omitempty"`
	ODataType                *string                                     `json:"@odata.type,omitempty"`
	Pod                      *SecurityKubernetesPodEvidence              `json:"pod,omitempty"`
	RemediationStatus        *SecurityContainerEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                     `json:"remediationStatusDetails,omitempty"`
	Roles                    *[]SecurityContainerEvidenceRoles           `json:"roles,omitempty"`
	Tags                     *[]string                                   `json:"tags,omitempty"`
	Verdict                  *SecurityContainerEvidenceVerdict           `json:"verdict,omitempty"`
}
