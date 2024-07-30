package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingUnitedStates struct {
	MovieRating *MediaContentRatingUnitedStatesMovieRating `json:"movieRating,omitempty"`
	ODataType   *string                                    `json:"@odata.type,omitempty"`
	TvRating    *MediaContentRatingUnitedStatesTvRating    `json:"tvRating,omitempty"`
}
