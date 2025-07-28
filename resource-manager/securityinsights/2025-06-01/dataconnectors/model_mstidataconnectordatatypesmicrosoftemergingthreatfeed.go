package dataconnectors

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed struct {
	LookbackPeriod string         `json:"lookbackPeriod"`
	State          *DataTypeState `json:"state,omitempty"`
}

func (o *MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed) GetLookbackPeriodAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.LookbackPeriod, "2006-01-02T15:04:05Z07:00")
}

func (o *MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed) SetLookbackPeriodAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LookbackPeriod = formatted
}
