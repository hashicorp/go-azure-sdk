package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AudioOperationPredicate struct {
	Album             *string
	AlbumArtist       *string
	Artist            *string
	Bitrate           *int64
	Composers         *string
	Copyright         *string
	Disc              *int64
	DiscCount         *int64
	Duration          *int64
	Genre             *string
	HasDrm            *bool
	IsVariableBitrate *bool
	ODataType         *string
	Title             *string
	Track             *int64
	TrackCount        *int64
	Year              *int64
}

func (p AudioOperationPredicate) Matches(input Audio) bool {

	if p.Album != nil && (input.Album == nil || *p.Album != *input.Album) {
		return false
	}

	if p.AlbumArtist != nil && (input.AlbumArtist == nil || *p.AlbumArtist != *input.AlbumArtist) {
		return false
	}

	if p.Artist != nil && (input.Artist == nil || *p.Artist != *input.Artist) {
		return false
	}

	if p.Bitrate != nil && (input.Bitrate == nil || *p.Bitrate != *input.Bitrate) {
		return false
	}

	if p.Composers != nil && (input.Composers == nil || *p.Composers != *input.Composers) {
		return false
	}

	if p.Copyright != nil && (input.Copyright == nil || *p.Copyright != *input.Copyright) {
		return false
	}

	if p.Disc != nil && (input.Disc == nil || *p.Disc != *input.Disc) {
		return false
	}

	if p.DiscCount != nil && (input.DiscCount == nil || *p.DiscCount != *input.DiscCount) {
		return false
	}

	if p.Duration != nil && (input.Duration == nil || *p.Duration != *input.Duration) {
		return false
	}

	if p.Genre != nil && (input.Genre == nil || *p.Genre != *input.Genre) {
		return false
	}

	if p.HasDrm != nil && (input.HasDrm == nil || *p.HasDrm != *input.HasDrm) {
		return false
	}

	if p.IsVariableBitrate != nil && (input.IsVariableBitrate == nil || *p.IsVariableBitrate != *input.IsVariableBitrate) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	if p.Track != nil && (input.Track == nil || *p.Track != *input.Track) {
		return false
	}

	if p.TrackCount != nil && (input.TrackCount == nil || *p.TrackCount != *input.TrackCount) {
		return false
	}

	if p.Year != nil && (input.Year == nil || *p.Year != *input.Year) {
		return false
	}

	return true
}
