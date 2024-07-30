package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAzureResourceEvidence struct {
	CreatedDateTime          *string                                         `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                       `json:"detailedRoles,omitempty"`
	ODataType                *string                                         `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityAzureResourceEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                         `json:"remediationStatusDetails,omitempty"`
	ResourceId               *string                                         `json:"resourceId,omitempty"`
	ResourceName             *string                                         `json:"resourceName,omitempty"`
	ResourceType             *string                                         `json:"resourceType,omitempty"`
	Roles                    *SecurityAzureResourceEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                                       `json:"tags,omitempty"`
	Verdict                  *SecurityAzureResourceEvidenceVerdict           `json:"verdict,omitempty"`
}
