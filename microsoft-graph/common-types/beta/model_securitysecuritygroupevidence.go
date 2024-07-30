package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySecurityGroupEvidence struct {
	CreatedDateTime          *string                                         `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                       `json:"detailedRoles,omitempty"`
	DisplayName              *string                                         `json:"displayName,omitempty"`
	ODataType                *string                                         `json:"@odata.type,omitempty"`
	RemediationStatus        *SecuritySecurityGroupEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                         `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecuritySecurityGroupEvidenceRoles             `json:"roles,omitempty"`
	SecurityGroupId          *string                                         `json:"securityGroupId,omitempty"`
	Tags                     *[]string                                       `json:"tags,omitempty"`
	Verdict                  *SecuritySecurityGroupEvidenceVerdict           `json:"verdict,omitempty"`
}
