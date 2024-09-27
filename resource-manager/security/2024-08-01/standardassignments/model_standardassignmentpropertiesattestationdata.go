package standardassignments

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardAssignmentPropertiesAttestationData struct {
	AssignedAssessment *AssignedAssessmentItem     `json:"assignedAssessment,omitempty"`
	ComplianceDate     *string                     `json:"complianceDate,omitempty"`
	ComplianceState    *AttestationComplianceState `json:"complianceState,omitempty"`
	Evidence           *[]AttestationEvidence      `json:"evidence,omitempty"`
}

func (o *StandardAssignmentPropertiesAttestationData) GetComplianceDateAsTime() (*time.Time, error) {
	if o.ComplianceDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ComplianceDate, "2006-01-02T15:04:05Z07:00")
}

func (o *StandardAssignmentPropertiesAttestationData) SetComplianceDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ComplianceDate = &formatted
}
