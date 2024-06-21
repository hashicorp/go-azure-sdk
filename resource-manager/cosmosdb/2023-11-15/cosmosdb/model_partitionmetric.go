package cosmosdb

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartitionMetric struct {
	EndTime             *string        `json:"endTime,omitempty"`
	MetricValues        *[]MetricValue `json:"metricValues,omitempty"`
	Name                *MetricName    `json:"name,omitempty"`
	PartitionId         *string        `json:"partitionId,omitempty"`
	PartitionKeyRangeId *string        `json:"partitionKeyRangeId,omitempty"`
	StartTime           *string        `json:"startTime,omitempty"`
	TimeGrain           *string        `json:"timeGrain,omitempty"`
	Unit                *UnitType      `json:"unit,omitempty"`
}
