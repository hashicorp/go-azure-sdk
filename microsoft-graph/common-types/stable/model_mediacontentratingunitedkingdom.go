package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingUnitedKingdom struct {
	MovieRating *MediaContentRatingUnitedKingdomMovieRating `json:"movieRating,omitempty"`
	ODataType   *string                                     `json:"@odata.type,omitempty"`
	TvRating    *MediaContentRatingUnitedKingdomTvRating    `json:"tvRating,omitempty"`
}
