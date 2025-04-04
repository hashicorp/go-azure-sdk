package advisorscore

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScoreEntity struct {
	CategoryCount          *float64 `json:"categoryCount,omitempty"`
	ConsumptionUnits       *float64 `json:"consumptionUnits,omitempty"`
	Date                   *string  `json:"date,omitempty"`
	ImpactedResourceCount  *float64 `json:"impactedResourceCount,omitempty"`
	PotentialScoreIncrease *float64 `json:"potentialScoreIncrease,omitempty"`
	Score                  *float64 `json:"score,omitempty"`
}
