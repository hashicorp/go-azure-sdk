package reservationrecommendationdetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendationDetailsModel struct {
	Etag       *string                                     `json:"etag,omitempty"`
	Id         *string                                     `json:"id,omitempty"`
	Location   *string                                     `json:"location,omitempty"`
	Name       *string                                     `json:"name,omitempty"`
	Properties *ReservationRecommendationDetailsProperties `json:"properties,omitempty"`
	Sku        *string                                     `json:"sku,omitempty"`
	Tags       *map[string]string                          `json:"tags,omitempty"`
	Type       *string                                     `json:"type,omitempty"`
}
