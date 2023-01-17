package costs

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetCostProperties struct {
	CostThresholds     *[]CostThresholdProperties `json:"costThresholds,omitempty"`
	CycleEndDateTime   *string                    `json:"cycleEndDateTime,omitempty"`
	CycleStartDateTime *string                    `json:"cycleStartDateTime,omitempty"`
	CycleType          *ReportingCycleType        `json:"cycleType,omitempty"`
	Status             *TargetCostStatus          `json:"status,omitempty"`
	Target             *int64                     `json:"target,omitempty"`
}

func (o *TargetCostProperties) GetCycleEndDateTimeAsTime() (*time.Time, error) {
	if o.CycleEndDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CycleEndDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *TargetCostProperties) SetCycleEndDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CycleEndDateTime = &formatted
}

func (o *TargetCostProperties) GetCycleStartDateTimeAsTime() (*time.Time, error) {
	if o.CycleStartDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CycleStartDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *TargetCostProperties) SetCycleStartDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CycleStartDateTime = &formatted
}
