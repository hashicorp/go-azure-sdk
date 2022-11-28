package forecast

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForecastDataset struct {
	Aggregation   map[string]ForecastAggregation `json:"aggregation"`
	Configuration *ForecastDatasetConfiguration  `json:"configuration"`
	Filter        *ForecastFilter                `json:"filter"`
	Granularity   *GranularityType               `json:"granularity,omitempty"`
}
