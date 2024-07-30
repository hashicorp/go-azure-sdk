package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingGermany struct {
	MovieRating *MediaContentRatingGermanyMovieRating `json:"movieRating,omitempty"`
	ODataType   *string                               `json:"@odata.type,omitempty"`
	TvRating    *MediaContentRatingGermanyTvRating    `json:"tvRating,omitempty"`
}
