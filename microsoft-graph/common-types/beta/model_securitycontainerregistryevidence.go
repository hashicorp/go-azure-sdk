package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerRegistryEvidence struct {
	CreatedDateTime          *string                                             `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                           `json:"detailedRoles,omitempty"`
	ODataType                *string                                             `json:"@odata.type,omitempty"`
	Registry                 *string                                             `json:"registry,omitempty"`
	RemediationStatus        *SecurityContainerRegistryEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                             `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityContainerRegistryEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                                           `json:"tags,omitempty"`
	Verdict                  *SecurityContainerRegistryEvidenceVerdict           `json:"verdict,omitempty"`
}
