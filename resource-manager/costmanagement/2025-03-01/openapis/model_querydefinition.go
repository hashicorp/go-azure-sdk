package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryDefinition struct {
	Dataset    QueryDataset     `json:"dataset"`
	TimePeriod *QueryTimePeriod `json:"timePeriod,omitempty"`
	Timeframe  TimeframeType    `json:"timeframe"`
	Type       ExportType       `json:"type"`
}
