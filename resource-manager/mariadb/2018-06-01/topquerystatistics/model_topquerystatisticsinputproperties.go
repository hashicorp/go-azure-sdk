package topquerystatistics

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
