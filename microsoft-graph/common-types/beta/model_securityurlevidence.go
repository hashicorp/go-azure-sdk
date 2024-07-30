package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlEvidence struct {
	CreatedDateTime          *string                               `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                             `json:"detailedRoles,omitempty"`
	ODataType                *string                               `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityUrlEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                               `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityUrlEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                             `json:"tags,omitempty"`
	Url                      *string                               `json:"url,omitempty"`
	Verdict                  *SecurityUrlEvidenceVerdict           `json:"verdict,omitempty"`
}
