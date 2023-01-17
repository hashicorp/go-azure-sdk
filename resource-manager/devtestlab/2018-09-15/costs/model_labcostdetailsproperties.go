package costs

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabCostDetailsProperties struct {
	Cost     *float64  `json:"cost,omitempty"`
	CostType *CostType `json:"costType,omitempty"`
	Date     *string   `json:"date,omitempty"`
}

func (o *LabCostDetailsProperties) GetDateAsTime() (*time.Time, error) {
	if o.Date == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Date, "2006-01-02T15:04:05Z07:00")
}

func (o *LabCostDetailsProperties) SetDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Date = &formatted
}
