package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerImageEvidence struct {
	CreatedDateTime          *string                                          `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                        `json:"detailedRoles,omitempty"`
	DigestImage              *SecurityContainerImageEvidence                  `json:"digestImage,omitempty"`
	ImageId                  *string                                          `json:"imageId,omitempty"`
	ODataType                *string                                          `json:"@odata.type,omitempty"`
	Registry                 *SecurityContainerRegistryEvidence               `json:"registry,omitempty"`
	RemediationStatus        *SecurityContainerImageEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                          `json:"remediationStatusDetails,omitempty"`
	Roles                    *[]SecurityContainerImageEvidenceRoles           `json:"roles,omitempty"`
	Tags                     *[]string                                        `json:"tags,omitempty"`
	Verdict                  *SecurityContainerImageEvidenceVerdict           `json:"verdict,omitempty"`
}
