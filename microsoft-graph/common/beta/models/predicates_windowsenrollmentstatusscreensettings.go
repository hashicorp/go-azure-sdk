package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsEnrollmentStatusScreenSettingsOperationPredicate struct {
	AllowDeviceUseBeforeProfileAndAppInstallComplete *bool
	AllowDeviceUseOnInstallFailure                   *bool
	AllowLogCollectionOnInstallFailure               *bool
	BlockDeviceSetupRetryByUser                      *bool
	CustomErrorMessage                               *string
	HideInstallationProgress                         *bool
	InstallProgressTimeoutInMinutes                  *int64
	ODataType                                        *string
}

func (p WindowsEnrollmentStatusScreenSettingsOperationPredicate) Matches(input WindowsEnrollmentStatusScreenSettings) bool {

	if p.AllowDeviceUseBeforeProfileAndAppInstallComplete != nil && (input.AllowDeviceUseBeforeProfileAndAppInstallComplete == nil || *p.AllowDeviceUseBeforeProfileAndAppInstallComplete != *input.AllowDeviceUseBeforeProfileAndAppInstallComplete) {
		return false
	}

	if p.AllowDeviceUseOnInstallFailure != nil && (input.AllowDeviceUseOnInstallFailure == nil || *p.AllowDeviceUseOnInstallFailure != *input.AllowDeviceUseOnInstallFailure) {
		return false
	}

	if p.AllowLogCollectionOnInstallFailure != nil && (input.AllowLogCollectionOnInstallFailure == nil || *p.AllowLogCollectionOnInstallFailure != *input.AllowLogCollectionOnInstallFailure) {
		return false
	}

	if p.BlockDeviceSetupRetryByUser != nil && (input.BlockDeviceSetupRetryByUser == nil || *p.BlockDeviceSetupRetryByUser != *input.BlockDeviceSetupRetryByUser) {
		return false
	}

	if p.CustomErrorMessage != nil && (input.CustomErrorMessage == nil || *p.CustomErrorMessage != *input.CustomErrorMessage) {
		return false
	}

	if p.HideInstallationProgress != nil && (input.HideInstallationProgress == nil || *p.HideInstallationProgress != *input.HideInstallationProgress) {
		return false
	}

	if p.InstallProgressTimeoutInMinutes != nil && (input.InstallProgressTimeoutInMinutes == nil || *p.InstallProgressTimeoutInMinutes != *input.InstallProgressTimeoutInMinutes) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
