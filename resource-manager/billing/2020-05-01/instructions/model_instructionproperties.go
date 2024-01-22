package instructions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstructionProperties struct {
	Amount       float64 `json:"amount"`
	CreationDate *string `json:"creationDate,omitempty"`
	EndDate      string  `json:"endDate"`
	StartDate    string  `json:"startDate"`
}

func (o *InstructionProperties) GetCreationDateAsTime() (*time.Time, error) {
	if o.CreationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *InstructionProperties) SetCreationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationDate = &formatted
}

func (o *InstructionProperties) GetEndDateAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.EndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *InstructionProperties) SetEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndDate = formatted
}

func (o *InstructionProperties) GetStartDateAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *InstructionProperties) SetStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartDate = formatted
}
