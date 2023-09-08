package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppConfigurationOperationPredicate struct {
	AppSupportsOemConfig *bool
	ConnectedAppsEnabled *bool
	CreatedDateTime      *string
	Description          *string
	DisplayName          *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	PackageId            *string
	PayloadJson          *string
	Version              *int64
}

func (p AndroidManagedStoreAppConfigurationOperationPredicate) Matches(input AndroidManagedStoreAppConfiguration) bool {

	if p.AppSupportsOemConfig != nil && (input.AppSupportsOemConfig == nil || *p.AppSupportsOemConfig != *input.AppSupportsOemConfig) {
		return false
	}

	if p.ConnectedAppsEnabled != nil && (input.ConnectedAppsEnabled == nil || *p.ConnectedAppsEnabled != *input.ConnectedAppsEnabled) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PackageId != nil && (input.PackageId == nil || *p.PackageId != *input.PackageId) {
		return false
	}

	if p.PayloadJson != nil && (input.PayloadJson == nil || *p.PayloadJson != *input.PayloadJson) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
