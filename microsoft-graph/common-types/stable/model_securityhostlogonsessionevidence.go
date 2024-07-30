package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostLogonSessionEvidence struct {
	Account                  *SecurityUserEvidence                              `json:"account,omitempty"`
	CreatedDateTime          *string                                            `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                          `json:"detailedRoles,omitempty"`
	EndUtcDateTime           *string                                            `json:"endUtcDateTime,omitempty"`
	Host                     *SecurityDeviceEvidence                            `json:"host,omitempty"`
	ODataType                *string                                            `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityHostLogonSessionEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                            `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityHostLogonSessionEvidenceRoles             `json:"roles,omitempty"`
	SessionId                *string                                            `json:"sessionId,omitempty"`
	StartUtcDateTime         *string                                            `json:"startUtcDateTime,omitempty"`
	Tags                     *[]string                                          `json:"tags,omitempty"`
	Verdict                  *SecurityHostLogonSessionEvidenceVerdict           `json:"verdict,omitempty"`
}
