package forecast

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForecastTimePeriod struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func (o *ForecastTimePeriod) GetFromAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.From, "2006-01-02T15:04:05Z07:00")
}

func (o *ForecastTimePeriod) SetFromAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.From = formatted
}

func (o *ForecastTimePeriod) GetToAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.To, "2006-01-02T15:04:05Z07:00")
}

func (o *ForecastTimePeriod) SetToAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.To = formatted
}
