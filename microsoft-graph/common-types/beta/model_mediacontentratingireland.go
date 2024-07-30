package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingIreland struct {
	MovieRating *MediaContentRatingIrelandMovieRating `json:"movieRating,omitempty"`
	ODataType   *string                               `json:"@odata.type,omitempty"`
	TvRating    *MediaContentRatingIrelandTvRating    `json:"tvRating,omitempty"`
}
