package pricesheetresults

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeterDetails struct {
	MeterCategory         *string  `json:"meterCategory,omitempty"`
	MeterLocation         *string  `json:"meterLocation,omitempty"`
	MeterName             *string  `json:"meterName,omitempty"`
	MeterSubCategory      *string  `json:"meterSubCategory,omitempty"`
	PretaxStandardRate    *float64 `json:"pretaxStandardRate,omitempty"`
	ServiceName           *string  `json:"serviceName,omitempty"`
	ServiceTier           *string  `json:"serviceTier,omitempty"`
	TotalIncludedQuantity *float64 `json:"totalIncludedQuantity,omitempty"`
	Unit                  *string  `json:"unit,omitempty"`
}
