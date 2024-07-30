package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRegistryValueEvidence struct {
	CreatedDateTime          *string                                         `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                       `json:"detailedRoles,omitempty"`
	MdeDeviceId              *string                                         `json:"mdeDeviceId,omitempty"`
	ODataType                *string                                         `json:"@odata.type,omitempty"`
	RegistryHive             *string                                         `json:"registryHive,omitempty"`
	RegistryKey              *string                                         `json:"registryKey,omitempty"`
	RegistryValue            *string                                         `json:"registryValue,omitempty"`
	RegistryValueName        *string                                         `json:"registryValueName,omitempty"`
	RegistryValueType        *string                                         `json:"registryValueType,omitempty"`
	RemediationStatus        *SecurityRegistryValueEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                         `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityRegistryValueEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                                       `json:"tags,omitempty"`
	Verdict                  *SecurityRegistryValueEvidenceVerdict           `json:"verdict,omitempty"`
}
