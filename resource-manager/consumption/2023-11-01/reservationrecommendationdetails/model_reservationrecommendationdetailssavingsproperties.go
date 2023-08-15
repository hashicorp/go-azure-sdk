package reservationrecommendationdetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendationDetailsSavingsProperties struct {
	CalculatedSavings    *[]ReservationRecommendationDetailsCalculatedSavingsProperties `json:"calculatedSavings,omitempty"`
	LookBackPeriod       *int64                                                         `json:"lookBackPeriod,omitempty"`
	RecommendedQuantity  *float64                                                       `json:"recommendedQuantity,omitempty"`
	ReservationOrderTerm *string                                                        `json:"reservationOrderTerm,omitempty"`
	SavingsType          *string                                                        `json:"savingsType,omitempty"`
	UnitOfMeasure        *string                                                        `json:"unitOfMeasure,omitempty"`
}
