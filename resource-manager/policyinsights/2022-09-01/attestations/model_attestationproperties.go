package attestations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationProperties struct {
	AssessmentDate              *string                `json:"assessmentDate,omitempty"`
	Comments                    *string                `json:"comments,omitempty"`
	ComplianceState             *ComplianceState       `json:"complianceState,omitempty"`
	Evidence                    *[]AttestationEvidence `json:"evidence,omitempty"`
	ExpiresOn                   *string                `json:"expiresOn,omitempty"`
	LastComplianceStateChangeAt *string                `json:"lastComplianceStateChangeAt,omitempty"`
	Metadata                    *interface{}           `json:"metadata,omitempty"`
	Owner                       *string                `json:"owner,omitempty"`
	PolicyAssignmentId          string                 `json:"policyAssignmentId"`
	PolicyDefinitionReferenceId *string                `json:"policyDefinitionReferenceId,omitempty"`
	ProvisioningState           *string                `json:"provisioningState,omitempty"`
}
