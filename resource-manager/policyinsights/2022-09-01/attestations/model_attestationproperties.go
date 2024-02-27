package attestations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *AttestationProperties) GetAssessmentDateAsTime() (*time.Time, error) {
	if o.AssessmentDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.AssessmentDate, "2006-01-02T15:04:05Z07:00")
}

func (o *AttestationProperties) SetAssessmentDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.AssessmentDate = &formatted
}

func (o *AttestationProperties) GetExpiresOnAsTime() (*time.Time, error) {
	if o.ExpiresOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpiresOn, "2006-01-02T15:04:05Z07:00")
}

func (o *AttestationProperties) SetExpiresOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiresOn = &formatted
}

func (o *AttestationProperties) GetLastComplianceStateChangeAtAsTime() (*time.Time, error) {
	if o.LastComplianceStateChangeAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastComplianceStateChangeAt, "2006-01-02T15:04:05Z07:00")
}
