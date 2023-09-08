package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaArtifactOperationPredicate struct {
	BoardSupportPackageVersion *string
	Description                *string
	DeviceModel                *string
	Id                         *string
	ODataType                  *string
	OsVersion                  *string
	PatchVersion               *string
	ReleaseNotesUrl            *string
}

func (p ZebraFotaArtifactOperationPredicate) Matches(input ZebraFotaArtifact) bool {

	if p.BoardSupportPackageVersion != nil && (input.BoardSupportPackageVersion == nil || *p.BoardSupportPackageVersion != *input.BoardSupportPackageVersion) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DeviceModel != nil && (input.DeviceModel == nil || *p.DeviceModel != *input.DeviceModel) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OsVersion != nil && (input.OsVersion == nil || *p.OsVersion != *input.OsVersion) {
		return false
	}

	if p.PatchVersion != nil && (input.PatchVersion == nil || *p.PatchVersion != *input.PatchVersion) {
		return false
	}

	if p.ReleaseNotesUrl != nil && (input.ReleaseNotesUrl == nil || *p.ReleaseNotesUrl != *input.ReleaseNotesUrl) {
		return false
	}

	return true
}
