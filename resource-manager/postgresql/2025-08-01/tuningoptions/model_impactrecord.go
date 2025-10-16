package tuningoptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImpactRecord struct {
	AbsoluteValue *float64 `json:"absoluteValue,omitempty"`
	DimensionName *string  `json:"dimensionName,omitempty"`
	QueryId       *int64   `json:"queryId,omitempty"`
	Unit          *string  `json:"unit,omitempty"`
}
