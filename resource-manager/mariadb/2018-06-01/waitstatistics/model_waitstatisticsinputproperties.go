package waitstatistics

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WaitStatisticsInputProperties struct {
	AggregationWindow    string `json:"aggregationWindow"`
	ObservationEndTime   string `json:"observationEndTime"`
	ObservationStartTime string `json:"observationStartTime"`
}

func (o *WaitStatisticsInputProperties) GetObservationEndTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.ObservationEndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WaitStatisticsInputProperties) SetObservationEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ObservationEndTime = formatted
}

func (o *WaitStatisticsInputProperties) GetObservationStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.ObservationStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WaitStatisticsInputProperties) SetObservationStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ObservationStartTime = formatted
}
