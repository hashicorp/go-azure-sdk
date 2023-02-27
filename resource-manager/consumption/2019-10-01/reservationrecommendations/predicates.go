// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package reservationrecommendations

type ReservationRecommendationOperationPredicate struct {
}

func (p ReservationRecommendationOperationPredicate) Matches(input ReservationRecommendation) bool {

	return true
}
