package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendationDetailsCalculatedSavingsProperties struct {
	OnDemandCost         *float64 `json:"onDemandCost,omitempty"`
	OverageCost          *float64 `json:"overageCost,omitempty"`
	Quantity             *float64 `json:"quantity,omitempty"`
	ReservationCost      *float64 `json:"reservationCost,omitempty"`
	ReservedUnitCount    *float64 `json:"reservedUnitCount,omitempty"`
	Savings              *float64 `json:"savings,omitempty"`
	TotalReservationCost *float64 `json:"totalReservationCost,omitempty"`
}
