package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamFunSettings struct {
	AllowCustomMemes      *bool                              `json:"allowCustomMemes,omitempty"`
	AllowGiphy            *bool                              `json:"allowGiphy,omitempty"`
	AllowStickersAndMemes *bool                              `json:"allowStickersAndMemes,omitempty"`
	GiphyContentRating    *TeamFunSettingsGiphyContentRating `json:"giphyContentRating,omitempty"`
	ODataType             *string                            `json:"@odata.type,omitempty"`
}
