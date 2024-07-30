package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRegistryKeyEvidence struct {
	CreatedDateTime          *string                                       `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                     `json:"detailedRoles,omitempty"`
	ODataType                *string                                       `json:"@odata.type,omitempty"`
	RegistryHive             *string                                       `json:"registryHive,omitempty"`
	RegistryKey              *string                                       `json:"registryKey,omitempty"`
	RemediationStatus        *SecurityRegistryKeyEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                       `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityRegistryKeyEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                                     `json:"tags,omitempty"`
	Verdict                  *SecurityRegistryKeyEvidenceVerdict           `json:"verdict,omitempty"`
}
