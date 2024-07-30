package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingJapan struct {
	MovieRating *MediaContentRatingJapanMovieRating `json:"movieRating,omitempty"`
	ODataType   *string                             `json:"@odata.type,omitempty"`
	TvRating    *MediaContentRatingJapanTvRating    `json:"tvRating,omitempty"`
}
