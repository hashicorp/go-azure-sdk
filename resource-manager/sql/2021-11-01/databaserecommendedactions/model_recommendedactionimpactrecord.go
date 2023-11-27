package databaserecommendedactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedActionImpactRecord struct {
	AbsoluteValue       *float64 `json:"absoluteValue,omitempty"`
	ChangeValueAbsolute *float64 `json:"changeValueAbsolute,omitempty"`
	ChangeValueRelative *float64 `json:"changeValueRelative,omitempty"`
	DimensionName       *string  `json:"dimensionName,omitempty"`
	Unit                *string  `json:"unit,omitempty"`
}
