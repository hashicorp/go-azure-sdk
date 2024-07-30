package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlobContainerEvidence struct {
	CreatedDateTime          *string                                         `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                       `json:"detailedRoles,omitempty"`
	Name                     *string                                         `json:"name,omitempty"`
	ODataType                *string                                         `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityBlobContainerEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                         `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityBlobContainerEvidenceRoles             `json:"roles,omitempty"`
	StorageResource          *SecurityAzureResourceEvidence                  `json:"storageResource,omitempty"`
	Tags                     *[]string                                       `json:"tags,omitempty"`
	Url                      *string                                         `json:"url,omitempty"`
	Verdict                  *SecurityBlobContainerEvidenceVerdict           `json:"verdict,omitempty"`
}
