package quotatier

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaTierProperties struct {
	AssignmentDate             *string                          `json:"assignmentDate,omitempty"`
	CurrentTierName            *string                          `json:"currentTierName,omitempty"`
	TierUpgradeEligibilityInfo *QuotaTierUpgradeEligibilityInfo `json:"tierUpgradeEligibilityInfo,omitempty"`
	TierUpgradePolicy          *TierUpgradePolicy               `json:"tierUpgradePolicy,omitempty"`
}

func (o *QuotaTierProperties) GetAssignmentDateAsTime() (*time.Time, error) {
	if o.AssignmentDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.AssignmentDate, "2006-01-02T15:04:05Z07:00")
}

func (o *QuotaTierProperties) SetAssignmentDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.AssignmentDate = &formatted
}
