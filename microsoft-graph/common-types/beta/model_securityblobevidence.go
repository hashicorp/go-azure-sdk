package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlobEvidence struct {
	BlobContainer            *SecurityBlobContainerEvidence         `json:"blobContainer,omitempty"`
	CreatedDateTime          *string                                `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                              `json:"detailedRoles,omitempty"`
	Etag                     *string                                `json:"etag,omitempty"`
	FileHashes               *[]SecurityFileHash                    `json:"fileHashes,omitempty"`
	Name                     *string                                `json:"name,omitempty"`
	ODataType                *string                                `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityBlobEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityBlobEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                              `json:"tags,omitempty"`
	Url                      *string                                `json:"url,omitempty"`
	Verdict                  *SecurityBlobEvidenceVerdict           `json:"verdict,omitempty"`
}
