package reservationrecommendations

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *ReservationRecommendationKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReservationRecommendationKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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
