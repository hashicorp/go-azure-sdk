package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendationDetailsProperties struct {
	Currency      *string                                             `json:"currency,omitempty"`
	Resource      *ReservationRecommendationDetailsResourceProperties `json:"resource,omitempty"`
	ResourceGroup *string                                             `json:"resourceGroup,omitempty"`
	Savings       *ReservationRecommendationDetailsSavingsProperties  `json:"savings,omitempty"`
	Scope         *string                                             `json:"scope,omitempty"`
	Usage         *ReservationRecommendationDetailsUsageProperties    `json:"usage,omitempty"`
}
