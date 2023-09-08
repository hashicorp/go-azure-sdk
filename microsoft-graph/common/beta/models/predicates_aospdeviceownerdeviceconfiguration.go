package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerDeviceConfigurationOperationPredicate struct {
	AppsBlockInstallFromUnknownSources             *bool
	BluetoothBlockConfiguration                    *bool
	BluetoothBlocked                               *bool
	CameraBlocked                                  *bool
	CreatedDateTime                                *string
	Description                                    *string
	DisplayName                                    *string
	FactoryResetBlocked                            *bool
	Id                                             *string
	LastModifiedDateTime                           *string
	ODataType                                      *string
	PasswordMinimumLength                          *int64
	PasswordMinutesOfInactivityBeforeScreenTimeout *int64
	PasswordSignInFailureCountBeforeFactoryReset   *int64
	ScreenCaptureBlocked                           *bool
	SecurityAllowDebuggingFeatures                 *bool
	StorageBlockExternalMedia                      *bool
	StorageBlockUsbFileTransfer                    *bool
	SupportsScopeTags                              *bool
	Version                                        *int64
	WifiBlockEditConfigurations                    *bool
}

func (p AospDeviceOwnerDeviceConfigurationOperationPredicate) Matches(input AospDeviceOwnerDeviceConfiguration) bool {

	if p.AppsBlockInstallFromUnknownSources != nil && (input.AppsBlockInstallFromUnknownSources == nil || *p.AppsBlockInstallFromUnknownSources != *input.AppsBlockInstallFromUnknownSources) {
		return false
	}

	if p.BluetoothBlockConfiguration != nil && (input.BluetoothBlockConfiguration == nil || *p.BluetoothBlockConfiguration != *input.BluetoothBlockConfiguration) {
		return false
	}

	if p.BluetoothBlocked != nil && (input.BluetoothBlocked == nil || *p.BluetoothBlocked != *input.BluetoothBlocked) {
		return false
	}

	if p.CameraBlocked != nil && (input.CameraBlocked == nil || *p.CameraBlocked != *input.CameraBlocked) {
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

	if p.FactoryResetBlocked != nil && (input.FactoryResetBlocked == nil || *p.FactoryResetBlocked != *input.FactoryResetBlocked) {
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

	if p.PasswordMinimumLength != nil && (input.PasswordMinimumLength == nil || *p.PasswordMinimumLength != *input.PasswordMinimumLength) {
		return false
	}

	if p.PasswordMinutesOfInactivityBeforeScreenTimeout != nil && (input.PasswordMinutesOfInactivityBeforeScreenTimeout == nil || *p.PasswordMinutesOfInactivityBeforeScreenTimeout != *input.PasswordMinutesOfInactivityBeforeScreenTimeout) {
		return false
	}

	if p.PasswordSignInFailureCountBeforeFactoryReset != nil && (input.PasswordSignInFailureCountBeforeFactoryReset == nil || *p.PasswordSignInFailureCountBeforeFactoryReset != *input.PasswordSignInFailureCountBeforeFactoryReset) {
		return false
	}

	if p.ScreenCaptureBlocked != nil && (input.ScreenCaptureBlocked == nil || *p.ScreenCaptureBlocked != *input.ScreenCaptureBlocked) {
		return false
	}

	if p.SecurityAllowDebuggingFeatures != nil && (input.SecurityAllowDebuggingFeatures == nil || *p.SecurityAllowDebuggingFeatures != *input.SecurityAllowDebuggingFeatures) {
		return false
	}

	if p.StorageBlockExternalMedia != nil && (input.StorageBlockExternalMedia == nil || *p.StorageBlockExternalMedia != *input.StorageBlockExternalMedia) {
		return false
	}

	if p.StorageBlockUsbFileTransfer != nil && (input.StorageBlockUsbFileTransfer == nil || *p.StorageBlockUsbFileTransfer != *input.StorageBlockUsbFileTransfer) {
		return false
	}

	if p.SupportsScopeTags != nil && (input.SupportsScopeTags == nil || *p.SupportsScopeTags != *input.SupportsScopeTags) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	if p.WifiBlockEditConfigurations != nil && (input.WifiBlockEditConfigurations == nil || *p.WifiBlockEditConfigurations != *input.WifiBlockEditConfigurations) {
		return false
	}

	return true
}
