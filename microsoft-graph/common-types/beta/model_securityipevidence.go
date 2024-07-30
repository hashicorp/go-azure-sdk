package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIpEvidence struct {
	CountryLetterCode        *string                              `json:"countryLetterCode,omitempty"`
	CreatedDateTime          *string                              `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                            `json:"detailedRoles,omitempty"`
	IpAddress                *string                              `json:"ipAddress,omitempty"`
	ODataType                *string                              `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityIpEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                              `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityIpEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                            `json:"tags,omitempty"`
	Verdict                  *SecurityIpEvidenceVerdict           `json:"verdict,omitempty"`
}
