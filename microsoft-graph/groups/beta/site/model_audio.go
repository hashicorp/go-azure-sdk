package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Audio struct {
	Album             *string `json:"album,omitempty"`
	AlbumArtist       *string `json:"albumArtist,omitempty"`
	Artist            *string `json:"artist,omitempty"`
	Bitrate           *int64  `json:"bitrate,omitempty"`
	Composers         *string `json:"composers,omitempty"`
	Copyright         *string `json:"copyright,omitempty"`
	Disc              *int64  `json:"disc,omitempty"`
	DiscCount         *int64  `json:"discCount,omitempty"`
	Duration          *int64  `json:"duration,omitempty"`
	Genre             *string `json:"genre,omitempty"`
	HasDrm            *bool   `json:"hasDrm,omitempty"`
	IsVariableBitrate *bool   `json:"isVariableBitrate,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	Title             *string `json:"title,omitempty"`
	Track             *int64  `json:"track,omitempty"`
	TrackCount        *int64  `json:"trackCount,omitempty"`
	Year              *int64  `json:"year,omitempty"`
}
