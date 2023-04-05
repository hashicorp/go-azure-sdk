package reservationrecommendations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendationKind string

const (
	ReservationRecommendationKindLegacy ReservationRecommendationKind = "legacy"
	ReservationRecommendationKindModern ReservationRecommendationKind = "modern"
)

func PossibleValuesForReservationRecommendationKind() []string {
	return []string{
		string(ReservationRecommendationKindLegacy),
		string(ReservationRecommendationKindModern),
	}
}
