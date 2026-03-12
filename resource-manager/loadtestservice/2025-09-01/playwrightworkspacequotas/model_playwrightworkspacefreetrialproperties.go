package playwrightworkspacequotas

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlaywrightWorkspaceFreeTrialProperties struct {
	AllocatedValue int64   `json:"allocatedValue"`
	CreatedAt      string  `json:"createdAt"`
	ExpiryAt       string  `json:"expiryAt"`
	PercentageUsed float64 `json:"percentageUsed"`
	UsedValue      float64 `json:"usedValue"`
}

func (o *PlaywrightWorkspaceFreeTrialProperties) GetCreatedAtAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.CreatedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *PlaywrightWorkspaceFreeTrialProperties) SetCreatedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedAt = formatted
}

func (o *PlaywrightWorkspaceFreeTrialProperties) GetExpiryAtAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.ExpiryAt, "2006-01-02T15:04:05Z07:00")
}

func (o *PlaywrightWorkspaceFreeTrialProperties) SetExpiryAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiryAt = formatted
}
