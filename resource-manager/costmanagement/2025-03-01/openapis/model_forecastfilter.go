package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForecastFilter struct {
	And        *[]ForecastFilter             `json:"and,omitempty"`
	Dimensions *ForecastComparisonExpression `json:"dimensions,omitempty"`
	Or         *[]ForecastFilter             `json:"or,omitempty"`
	Tags       *ForecastComparisonExpression `json:"tags,omitempty"`
}
