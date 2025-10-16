package dataconnectors

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PremiumMdtiDataConnectorProperties struct {
	DataTypes           PremiumMdtiDataConnectorDataTypes `json:"dataTypes"`
	LookbackPeriod      string                            `json:"lookbackPeriod"`
	RequiredSKUsPresent *bool                             `json:"requiredSKUsPresent,omitempty"`
	TenantId            *string                           `json:"tenantId,omitempty"`
}

func (o *PremiumMdtiDataConnectorProperties) GetLookbackPeriodAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.LookbackPeriod, "2006-01-02T15:04:05Z07:00")
}

func (o *PremiumMdtiDataConnectorProperties) SetLookbackPeriodAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LookbackPeriod = formatted
}
