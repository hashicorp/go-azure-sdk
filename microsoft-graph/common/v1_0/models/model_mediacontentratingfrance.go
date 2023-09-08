package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingFrance struct {
	MovieRating *MediaContentRatingFranceMovieRating `json:"movieRating,omitempty"`
	ODataType   *string                              `json:"@odata.type,omitempty"`
	TvRating    *MediaContentRatingFranceTvRating    `json:"tvRating,omitempty"`
}
