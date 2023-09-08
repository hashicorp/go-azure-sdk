package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EnrollmentCompletionPageConfigurationOperationPredicate struct {
	AllowDeviceResetOnInstallFailure        *bool
	AllowDeviceUseOnInstallFailure          *bool
	AllowLogCollectionOnInstallFailure      *bool
	AllowNonBlockingAppInstallation         *bool
	BlockDeviceSetupRetryByUser             *bool
	CreatedDateTime                         *string
	CustomErrorMessage                      *string
	Description                             *string
	DisableUserStatusTrackingAfterFirstUser *bool
	DisplayName                             *string
	Id                                      *string
	InstallProgressTimeoutInMinutes         *int64
	InstallQualityUpdates                   *bool
	LastModifiedDateTime                    *string
	ODataType                               *string
	Priority                                *int64
	ShowInstallationProgress                *bool
	TrackInstallProgressForAutopilotOnly    *bool
	Version                                 *int64
}

func (p Windows10EnrollmentCompletionPageConfigurationOperationPredicate) Matches(input Windows10EnrollmentCompletionPageConfiguration) bool {

	if p.AllowDeviceResetOnInstallFailure != nil && (input.AllowDeviceResetOnInstallFailure == nil || *p.AllowDeviceResetOnInstallFailure != *input.AllowDeviceResetOnInstallFailure) {
		return false
	}

	if p.AllowDeviceUseOnInstallFailure != nil && (input.AllowDeviceUseOnInstallFailure == nil || *p.AllowDeviceUseOnInstallFailure != *input.AllowDeviceUseOnInstallFailure) {
		return false
	}

	if p.AllowLogCollectionOnInstallFailure != nil && (input.AllowLogCollectionOnInstallFailure == nil || *p.AllowLogCollectionOnInstallFailure != *input.AllowLogCollectionOnInstallFailure) {
		return false
	}

	if p.AllowNonBlockingAppInstallation != nil && (input.AllowNonBlockingAppInstallation == nil || *p.AllowNonBlockingAppInstallation != *input.AllowNonBlockingAppInstallation) {
		return false
	}

	if p.BlockDeviceSetupRetryByUser != nil && (input.BlockDeviceSetupRetryByUser == nil || *p.BlockDeviceSetupRetryByUser != *input.BlockDeviceSetupRetryByUser) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.CustomErrorMessage != nil && (input.CustomErrorMessage == nil || *p.CustomErrorMessage != *input.CustomErrorMessage) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisableUserStatusTrackingAfterFirstUser != nil && (input.DisableUserStatusTrackingAfterFirstUser == nil || *p.DisableUserStatusTrackingAfterFirstUser != *input.DisableUserStatusTrackingAfterFirstUser) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InstallProgressTimeoutInMinutes != nil && (input.InstallProgressTimeoutInMinutes == nil || *p.InstallProgressTimeoutInMinutes != *input.InstallProgressTimeoutInMinutes) {
		return false
	}

	if p.InstallQualityUpdates != nil && (input.InstallQualityUpdates == nil || *p.InstallQualityUpdates != *input.InstallQualityUpdates) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Priority != nil && (input.Priority == nil || *p.Priority != *input.Priority) {
		return false
	}

	if p.ShowInstallationProgress != nil && (input.ShowInstallationProgress == nil || *p.ShowInstallationProgress != *input.ShowInstallationProgress) {
		return false
	}

	if p.TrackInstallProgressForAutopilotOnly != nil && (input.TrackInstallProgressForAutopilotOnly == nil || *p.TrackInstallProgressForAutopilotOnly != *input.TrackInstallProgressForAutopilotOnly) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
