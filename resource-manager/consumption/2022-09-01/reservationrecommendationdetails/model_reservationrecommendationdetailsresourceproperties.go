package reservationrecommendationdetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendationDetailsResourceProperties struct {
	AppliedScopes   *[]string `json:"appliedScopes,omitempty"`
	OnDemandRate    *float64  `json:"onDemandRate,omitempty"`
	Product         *string   `json:"product,omitempty"`
	Region          *string   `json:"region,omitempty"`
	ReservationRate *float64  `json:"reservationRate,omitempty"`
	ResourceType    *string   `json:"resourceType,omitempty"`
}
