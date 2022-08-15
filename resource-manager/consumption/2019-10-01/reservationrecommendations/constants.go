package reservationrecommendations

import "strings"

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

func parseReservationRecommendationKind(input string) (*ReservationRecommendationKind, error) {
	vals := map[string]ReservationRecommendationKind{
		"legacy": ReservationRecommendationKindLegacy,
		"modern": ReservationRecommendationKindModern,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReservationRecommendationKind(input)
	return &out, nil
}
