package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendationDetailsUsageProperties struct {
	FirstConsumptionDate *string    `json:"firstConsumptionDate,omitempty"`
	LastConsumptionDate  *string    `json:"lastConsumptionDate,omitempty"`
	LookBackUnitType     *string    `json:"lookBackUnitType,omitempty"`
	UsageData            *[]float64 `json:"usageData,omitempty"`
	UsageGrain           *string    `json:"usageGrain,omitempty"`
}
