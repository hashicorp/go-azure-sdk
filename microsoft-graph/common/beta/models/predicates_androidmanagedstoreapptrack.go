package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppTrackOperationPredicate struct {
	ODataType  *string
	TrackAlias *string
	TrackId    *string
}

func (p AndroidManagedStoreAppTrackOperationPredicate) Matches(input AndroidManagedStoreAppTrack) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TrackAlias != nil && (input.TrackAlias == nil || *p.TrackAlias != *input.TrackAlias) {
		return false
	}

	if p.TrackId != nil && (input.TrackId == nil || *p.TrackId != *input.TrackId) {
		return false
	}

	return true
}
