package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppOperationPredicate struct {
	CommittedContentVersion        *string
	CreatedDateTime                *string
	Description                    *string
	Developer                      *string
	DisplayName                    *string
	FileName                       *string
	Id                             *string
	InformationUrl                 *string
	InstallCommandLine             *string
	IsFeatured                     *bool
	LastModifiedDateTime           *string
	MinimumCpuSpeedInMHz           *int64
	MinimumFreeDiskSpaceInMB       *int64
	MinimumMemoryInMB              *int64
	MinimumNumberOfProcessors      *int64
	MinimumSupportedWindowsRelease *string
	Notes                          *string
	ODataType                      *string
	Owner                          *string
	PrivacyInformationUrl          *string
	Publisher                      *string
	SetupFilePath                  *string
	Size                           *int64
	UninstallCommandLine           *string
}

func (p Win32LobAppOperationPredicate) Matches(input Win32LobApp) bool {

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

	if p.InstallCommandLine != nil && (input.InstallCommandLine == nil || *p.InstallCommandLine != *input.InstallCommandLine) {
		return false
	}

	if p.IsFeatured != nil && (input.IsFeatured == nil || *p.IsFeatured != *input.IsFeatured) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.MinimumCpuSpeedInMHz != nil && (input.MinimumCpuSpeedInMHz == nil || *p.MinimumCpuSpeedInMHz != *input.MinimumCpuSpeedInMHz) {
		return false
	}

	if p.MinimumFreeDiskSpaceInMB != nil && (input.MinimumFreeDiskSpaceInMB == nil || *p.MinimumFreeDiskSpaceInMB != *input.MinimumFreeDiskSpaceInMB) {
		return false
	}

	if p.MinimumMemoryInMB != nil && (input.MinimumMemoryInMB == nil || *p.MinimumMemoryInMB != *input.MinimumMemoryInMB) {
		return false
	}

	if p.MinimumNumberOfProcessors != nil && (input.MinimumNumberOfProcessors == nil || *p.MinimumNumberOfProcessors != *input.MinimumNumberOfProcessors) {
		return false
	}

	if p.MinimumSupportedWindowsRelease != nil && (input.MinimumSupportedWindowsRelease == nil || *p.MinimumSupportedWindowsRelease != *input.MinimumSupportedWindowsRelease) {
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

	if p.PrivacyInformationUrl != nil && (input.PrivacyInformationUrl == nil || *p.PrivacyInformationUrl != *input.PrivacyInformationUrl) {
		return false
	}

	if p.Publisher != nil && (input.Publisher == nil || *p.Publisher != *input.Publisher) {
		return false
	}

	if p.SetupFilePath != nil && (input.SetupFilePath == nil || *p.SetupFilePath != *input.SetupFilePath) {
		return false
	}

	if p.Size != nil && (input.Size == nil || *p.Size != *input.Size) {
		return false
	}

	if p.UninstallCommandLine != nil && (input.UninstallCommandLine == nil || *p.UninstallCommandLine != *input.UninstallCommandLine) {
		return false
	}

	return true
}
