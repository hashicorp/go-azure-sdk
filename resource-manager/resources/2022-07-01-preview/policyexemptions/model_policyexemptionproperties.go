package policyexemptions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyExemptionProperties struct {
	AssignmentScopeValidation    *AssignmentScopeValidation `json:"assignmentScopeValidation,omitempty"`
	Description                  *string                    `json:"description,omitempty"`
	DisplayName                  *string                    `json:"displayName,omitempty"`
	ExemptionCategory            ExemptionCategory          `json:"exemptionCategory"`
	ExpiresOn                    *string                    `json:"expiresOn,omitempty"`
	Metadata                     *interface{}               `json:"metadata,omitempty"`
	PolicyAssignmentId           string                     `json:"policyAssignmentId"`
	PolicyDefinitionReferenceIds *[]string                  `json:"policyDefinitionReferenceIds,omitempty"`
	ResourceSelectors            *[]ResourceSelector        `json:"resourceSelectors,omitempty"`
}

func (o *PolicyExemptionProperties) GetExpiresOnAsTime() (*time.Time, error) {
	if o.ExpiresOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpiresOn, "2006-01-02T15:04:05Z07:00")
}

func (o *PolicyExemptionProperties) SetExpiresOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiresOn = &formatted
}
