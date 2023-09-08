package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAndroidLobAppOperationPredicate struct {
	CommittedContentVersion *string
	CreatedDateTime         *string
	Description             *string
	Developer               *string
	DisplayName             *string
	FileName                *string
	Id                      *string
	InformationUrl          *string
	IsFeatured              *bool
	LastModifiedDateTime    *string
	Notes                   *string
	ODataType               *string
	Owner                   *string
	PackageId               *string
	PrivacyInformationUrl   *string
	Publisher               *string
	Size                    *int64
	Version                 *string
	VersionCode             *string
	VersionName             *string
}

func (p ManagedAndroidLobAppOperationPredicate) Matches(input ManagedAndroidLobApp) bool {

	if p.CommittedContentVersion != nil && (input.CommittedContentVersion == nil || *p.CommittedContentVersion != *input.CommittedContentVersion) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Developer != nil && (input.Developer == nil || *p.Developer != *input.Developer) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.FileName != nil && (input.FileName == nil || *p.FileName != *input.FileName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InformationUrl != nil && (input.InformationUrl == nil || *p.InformationUrl != *input.InformationUrl) {
		return false
	}

	if p.IsFeatured != nil && (input.IsFeatured == nil || *p.IsFeatured != *input.IsFeatured) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.Notes != nil && (input.Notes == nil || *p.Notes != *input.Notes) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Owner != nil && (input.Owner == nil || *p.Owner != *input.Owner) {
		return false
	}

	if p.PackageId != nil && (input.PackageId == nil || *p.PackageId != *input.PackageId) {
		return false
	}

	if p.PrivacyInformationUrl != nil && (input.PrivacyInformationUrl == nil || *p.PrivacyInformationUrl != *input.PrivacyInformationUrl) {
		return false
	}

	if p.Publisher != nil && (input.Publisher == nil || *p.Publisher != *input.Publisher) {
		return false
	}

	if p.Size != nil && (input.Size == nil || *p.Size != *input.Size) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	if p.VersionCode != nil && (input.VersionCode == nil || *p.VersionCode != *input.VersionCode) {
		return false
	}

	if p.VersionName != nil && (input.VersionName == nil || *p.VersionName != *input.VersionName) {
		return false
	}

	return true
}
