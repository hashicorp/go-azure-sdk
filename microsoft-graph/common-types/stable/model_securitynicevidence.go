package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityNicEvidence struct {
	CreatedDateTime          *string                               `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                             `json:"detailedRoles,omitempty"`
	IpAddress                *SecurityIpEvidence                   `json:"ipAddress,omitempty"`
	MacAddress               *string                               `json:"macAddress,omitempty"`
	ODataType                *string                               `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityNicEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                               `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityNicEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                             `json:"tags,omitempty"`
	Verdict                  *SecurityNicEvidenceVerdict           `json:"verdict,omitempty"`
	Vlans                    *[]string                             `json:"vlans,omitempty"`
}
