package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForecastDefinition struct {
	Dataset                 ForecastDataset     `json:"dataset"`
	IncludeActualCost       *bool               `json:"includeActualCost,omitempty"`
	IncludeFreshPartialCost *bool               `json:"includeFreshPartialCost,omitempty"`
	TimePeriod              *ForecastTimePeriod `json:"timePeriod,omitempty"`
	Timeframe               ForecastTimeframe   `json:"timeframe"`
	Type                    ForecastType        `json:"type"`
}
