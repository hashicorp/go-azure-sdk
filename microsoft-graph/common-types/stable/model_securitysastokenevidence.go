package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySasTokenEvidence struct {
	AllowedIpAddresses       *string                                    `json:"allowedIpAddresses,omitempty"`
	AllowedResourceTypes     *[]string                                  `json:"allowedResourceTypes,omitempty"`
	AllowedServices          *[]string                                  `json:"allowedServices,omitempty"`
	CreatedDateTime          *string                                    `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                  `json:"detailedRoles,omitempty"`
	ExpiryDateTime           *string                                    `json:"expiryDateTime,omitempty"`
	ODataType                *string                                    `json:"@odata.type,omitempty"`
	Permissions              *[]string                                  `json:"permissions,omitempty"`
	Protocol                 *string                                    `json:"protocol,omitempty"`
	RemediationStatus        *SecuritySasTokenEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                    `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecuritySasTokenEvidenceRoles             `json:"roles,omitempty"`
	SignatureHash            *string                                    `json:"signatureHash,omitempty"`
	SignedWith               *string                                    `json:"signedWith,omitempty"`
	StartDateTime            *string                                    `json:"startDateTime,omitempty"`
	StorageResource          *SecurityAzureResourceEvidence             `json:"storageResource,omitempty"`
	Tags                     *[]string                                  `json:"tags,omitempty"`
	Verdict                  *SecuritySasTokenEvidenceVerdict           `json:"verdict,omitempty"`
}
