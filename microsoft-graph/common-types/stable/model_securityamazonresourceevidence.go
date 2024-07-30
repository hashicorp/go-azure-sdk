package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAmazonResourceEvidence struct {
	AmazonAccountId          *string                                          `json:"amazonAccountId,omitempty"`
	AmazonResourceId         *string                                          `json:"amazonResourceId,omitempty"`
	CreatedDateTime          *string                                          `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                        `json:"detailedRoles,omitempty"`
	ODataType                *string                                          `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityAmazonResourceEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                          `json:"remediationStatusDetails,omitempty"`
	ResourceName             *string                                          `json:"resourceName,omitempty"`
	ResourceType             *string                                          `json:"resourceType,omitempty"`
	Roles                    *SecurityAmazonResourceEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                                        `json:"tags,omitempty"`
	Verdict                  *SecurityAmazonResourceEvidenceVerdict           `json:"verdict,omitempty"`
}
