package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingCanada struct {
	MovieRating *MediaContentRatingCanadaMovieRating `json:"movieRating,omitempty"`
	ODataType   *string                              `json:"@odata.type,omitempty"`
	TvRating    *MediaContentRatingCanadaTvRating    `json:"tvRating,omitempty"`
}
