package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosiPadOSWebClipOperationPredicate struct {
	AppUrl                            *string
	CreatedDateTime                   *string
	DependentAppCount                 *int64
	Description                       *string
	Developer                         *string
	DisplayName                       *string
	FullScreenEnabled                 *bool
	Id                                *string
	IgnoreManifestScope               *bool
	InformationUrl                    *string
	IsAssigned                        *bool
	IsFeatured                        *bool
	LastModifiedDateTime              *string
	Notes                             *string
	ODataType                         *string
	Owner                             *string
	PreComposedIconEnabled            *bool
	PrivacyInformationUrl             *string
	Publisher                         *string
	SupersededAppCount                *int64
	SupersedingAppCount               *int64
	TargetApplicationBundleIdentifier *string
	UploadState                       *int64
	UseManagedBrowser                 *bool
}

func (p IosiPadOSWebClipOperationPredicate) Matches(input IosiPadOSWebClip) bool {

	if p.AppUrl != nil && (input.AppUrl == nil || *p.AppUrl != *input.AppUrl) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DependentAppCount != nil && (input.DependentAppCount == nil || *p.DependentAppCount != *input.DependentAppCount) {
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

	if p.FullScreenEnabled != nil && (input.FullScreenEnabled == nil || *p.FullScreenEnabled != *input.FullScreenEnabled) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IgnoreManifestScope != nil && (input.IgnoreManifestScope == nil || *p.IgnoreManifestScope != *input.IgnoreManifestScope) {
		return false
	}

	if p.InformationUrl != nil && (input.InformationUrl == nil || *p.InformationUrl != *input.InformationUrl) {
		return false
	}

	if p.IsAssigned != nil && (input.IsAssigned == nil || *p.IsAssigned != *input.IsAssigned) {
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

	if p.PreComposedIconEnabled != nil && (input.PreComposedIconEnabled == nil || *p.PreComposedIconEnabled != *input.PreComposedIconEnabled) {
		return false
	}

	if p.PrivacyInformationUrl != nil && (input.PrivacyInformationUrl == nil || *p.PrivacyInformationUrl != *input.PrivacyInformationUrl) {
		return false
	}

	if p.Publisher != nil && (input.Publisher == nil || *p.Publisher != *input.Publisher) {
		return false
	}

	if p.SupersededAppCount != nil && (input.SupersededAppCount == nil || *p.SupersededAppCount != *input.SupersededAppCount) {
		return false
	}

	if p.SupersedingAppCount != nil && (input.SupersedingAppCount == nil || *p.SupersedingAppCount != *input.SupersedingAppCount) {
		return false
	}

	if p.TargetApplicationBundleIdentifier != nil && (input.TargetApplicationBundleIdentifier == nil || *p.TargetApplicationBundleIdentifier != *input.TargetApplicationBundleIdentifier) {
		return false
	}

	if p.UploadState != nil && (input.UploadState == nil || *p.UploadState != *input.UploadState) {
		return false
	}

	if p.UseManagedBrowser != nil && (input.UseManagedBrowser == nil || *p.UseManagedBrowser != *input.UseManagedBrowser) {
		return false
	}

	return true
}
