package forecasts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForecastProperties struct {
	Charge           *float64                                     `json:"charge,omitempty"`
	ChargeType       *ChargeType                                  `json:"chargeType,omitempty"`
	ConfidenceLevels *[]ForecastPropertiesConfidenceLevelsInlined `json:"confidenceLevels,omitempty"`
	Currency         *string                                      `json:"currency,omitempty"`
	Grain            *Grain                                       `json:"grain,omitempty"`
	UsageDate        *string                                      `json:"usageDate,omitempty"`
}
