package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosDeviceFeaturesConfigurationOperationPredicate struct {
	AssetTagTemplate     *string
	CreatedDateTime      *string
	Description          *string
	DisplayName          *string
	Id                   *string
	LastModifiedDateTime *string
	LockScreenFootnote   *string
	ODataType            *string
	Version              *int64
}

func (p IosDeviceFeaturesConfigurationOperationPredicate) Matches(input IosDeviceFeaturesConfiguration) bool {

	if p.AssetTagTemplate != nil && (input.AssetTagTemplate == nil || *p.AssetTagTemplate != *input.AssetTagTemplate) {
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

	if p.LockScreenFootnote != nil && (input.LockScreenFootnote == nil || *p.LockScreenFootnote != *input.LockScreenFootnote) {
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
