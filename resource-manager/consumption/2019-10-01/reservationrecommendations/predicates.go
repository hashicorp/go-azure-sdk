package reservationrecommendations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendationOperationPredicate struct {
}

func (p ReservationRecommendationOperationPredicate) Matches(input ReservationRecommendation) bool {

	return true
}
