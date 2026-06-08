package summaryrules

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuleDefinition struct {
	BinDelay         *int64            `json:"binDelay,omitempty"`
	BinSize          *int64            `json:"binSize,omitempty"`
	BinStartTime     *string           `json:"binStartTime,omitempty"`
	DestinationTable *string           `json:"destinationTable,omitempty"`
	Query            *string           `json:"query,omitempty"`
	TimeSelector     *TimeSelectorEnum `json:"timeSelector,omitempty"`
}

func (o *RuleDefinition) GetBinStartTimeAsTime() (*time.Time, error) {
	if o.BinStartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BinStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RuleDefinition) SetBinStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BinStartTime = &formatted
}
