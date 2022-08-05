package topquerystatistics

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TopQueryStatisticsInputProperties struct {
	AggregationFunction  string `json:"aggregationFunction"`
	AggregationWindow    string `json:"aggregationWindow"`
	NumberOfTopQueries   int64  `json:"numberOfTopQueries"`
	ObservationEndTime   string `json:"observationEndTime"`
	ObservationStartTime string `json:"observationStartTime"`
	ObservedMetric       string `json:"observedMetric"`
}

func (o *TopQueryStatisticsInputProperties) GetObservationEndTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.ObservationEndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *TopQueryStatisticsInputProperties) SetObservationEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ObservationEndTime = formatted
}

func (o *TopQueryStatisticsInputProperties) GetObservationStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.ObservationStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *TopQueryStatisticsInputProperties) SetObservationStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ObservationStartTime = formatted
}
