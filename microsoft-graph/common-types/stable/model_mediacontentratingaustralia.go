package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingAustralia struct {
	MovieRating *MediaContentRatingAustraliaMovieRating `json:"movieRating,omitempty"`
	ODataType   *string                                 `json:"@odata.type,omitempty"`
	TvRating    *MediaContentRatingAustraliaTvRating    `json:"tvRating,omitempty"`
}
