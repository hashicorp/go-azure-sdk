package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesFeatureUpdateCatalogEntryOperationPredicate struct {
	BuildNumber             *string
	DeployableUntilDateTime *string
	DisplayName             *string
	Id                      *string
	ODataType               *string
	ReleaseDateTime         *string
	Version                 *string
}

func (p WindowsUpdatesFeatureUpdateCatalogEntryOperationPredicate) Matches(input WindowsUpdatesFeatureUpdateCatalogEntry) bool {

	if p.BuildNumber != nil && (input.BuildNumber == nil || *p.BuildNumber != *input.BuildNumber) {
		return false
	}

	if p.DeployableUntilDateTime != nil && (input.DeployableUntilDateTime == nil || *p.DeployableUntilDateTime != *input.DeployableUntilDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReleaseDateTime != nil && (input.ReleaseDateTime == nil || *p.ReleaseDateTime != *input.ReleaseDateTime) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
