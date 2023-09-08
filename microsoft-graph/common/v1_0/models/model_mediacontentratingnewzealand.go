package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingNewZealand struct {
	MovieRating *MediaContentRatingNewZealandMovieRating `json:"movieRating,omitempty"`
	ODataType   *string                                  `json:"@odata.type,omitempty"`
	TvRating    *MediaContentRatingNewZealandTvRating    `json:"tvRating,omitempty"`
}
