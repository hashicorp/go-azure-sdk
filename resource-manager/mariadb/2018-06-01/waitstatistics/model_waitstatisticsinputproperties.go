package waitstatistics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WaitStatisticsInputProperties struct {
	AggregationWindow    string `json:"aggregationWindow"`
	ObservationEndTime   string `json:"observationEndTime"`
	ObservationStartTime string `json:"observationStartTime"`
}
