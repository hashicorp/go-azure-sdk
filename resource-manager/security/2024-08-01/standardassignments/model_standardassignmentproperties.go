package standardassignments

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardAssignmentProperties struct {
	AssignedStandard *AssignedStandardItem                        `json:"assignedStandard,omitempty"`
	AttestationData  *StandardAssignmentPropertiesAttestationData `json:"attestationData,omitempty"`
	Description      *string                                      `json:"description,omitempty"`
	DisplayName      *string                                      `json:"displayName,omitempty"`
	Effect           *Effect                                      `json:"effect,omitempty"`
	ExcludedScopes   *[]string                                    `json:"excludedScopes,omitempty"`
	ExemptionData    *StandardAssignmentPropertiesExemptionData   `json:"exemptionData,omitempty"`
	ExpiresOn        *string                                      `json:"expiresOn,omitempty"`
	Metadata         *StandardAssignmentMetadata                  `json:"metadata,omitempty"`
}

func (o *StandardAssignmentProperties) GetExpiresOnAsTime() (*time.Time, error) {
	if o.ExpiresOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpiresOn, "2006-01-02T15:04:05Z07:00")
}

func (o *StandardAssignmentProperties) SetExpiresOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiresOn = &formatted
}
