package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertEvidence struct {
	CreatedDateTime          *string                                 `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                               `json:"detailedRoles,omitempty"`
	ODataType                *string                                 `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityAlertEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                 `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityAlertEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                               `json:"tags,omitempty"`
	Verdict                  *SecurityAlertEvidenceVerdict           `json:"verdict,omitempty"`
}
