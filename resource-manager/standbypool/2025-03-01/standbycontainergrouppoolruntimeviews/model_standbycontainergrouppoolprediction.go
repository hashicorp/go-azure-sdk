package standbycontainergrouppoolruntimeviews

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandbyContainerGroupPoolPrediction struct {
	ForecastInfo      string                                  `json:"forecastInfo"`
	ForecastStartTime string                                  `json:"forecastStartTime"`
	ForecastValues    StandbyContainerGroupPoolForecastValues `json:"forecastValues"`
}

func (o *StandbyContainerGroupPoolPrediction) GetForecastStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.ForecastStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *StandbyContainerGroupPoolPrediction) SetForecastStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ForecastStartTime = formatted
}
