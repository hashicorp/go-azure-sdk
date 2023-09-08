package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessBranchSiteOperationPredicate struct {
	BandwidthCapacity    *int64
	Country              *string
	Id                   *string
	LastModifiedDateTime *string
	Name                 *string
	ODataType            *string
	Version              *string
}

func (p NetworkaccessBranchSiteOperationPredicate) Matches(input NetworkaccessBranchSite) bool {

	if p.BandwidthCapacity != nil && (input.BandwidthCapacity == nil || *p.BandwidthCapacity != *input.BandwidthCapacity) {
		return false
	}

	if p.Country != nil && (input.Country == nil || *p.Country != *input.Country) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
