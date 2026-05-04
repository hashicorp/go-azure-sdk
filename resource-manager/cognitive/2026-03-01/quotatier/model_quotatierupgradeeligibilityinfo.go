package quotatier

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaTierUpgradeEligibilityInfo struct {
	NextTierName                *string                    `json:"nextTierName,omitempty"`
	UpgradeApplicableDate       *string                    `json:"upgradeApplicableDate,omitempty"`
	UpgradeAvailabilityStatus   *UpgradeAvailabilityStatus `json:"upgradeAvailabilityStatus,omitempty"`
	UpgradeUnavailabilityReason *string                    `json:"upgradeUnavailabilityReason,omitempty"`
}

func (o *QuotaTierUpgradeEligibilityInfo) GetUpgradeApplicableDateAsTime() (*time.Time, error) {
	if o.UpgradeApplicableDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.UpgradeApplicableDate, "2006-01-02T15:04:05Z07:00")
}

func (o *QuotaTierUpgradeEligibilityInfo) SetUpgradeApplicableDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.UpgradeApplicableDate = &formatted
}
