package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepOnboardingSettingOperationPredicate struct {
	AppleIdentifier                     *string
	DataSharingConsentGranted           *bool
	Id                                  *string
	LastModifiedDateTime                *string
	LastSuccessfulSyncDateTime          *string
	LastSyncErrorCode                   *int64
	LastSyncTriggeredDateTime           *string
	ODataType                           *string
	ShareTokenWithSchoolDataSyncService *bool
	SyncedDeviceCount                   *int64
	TokenExpirationDateTime             *string
	TokenName                           *string
}

func (p DepOnboardingSettingOperationPredicate) Matches(input DepOnboardingSetting) bool {

	if p.AppleIdentifier != nil && (input.AppleIdentifier == nil || *p.AppleIdentifier != *input.AppleIdentifier) {
		return false
	}

	if p.DataSharingConsentGranted != nil && (input.DataSharingConsentGranted == nil || *p.DataSharingConsentGranted != *input.DataSharingConsentGranted) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.LastSuccessfulSyncDateTime != nil && (input.LastSuccessfulSyncDateTime == nil || *p.LastSuccessfulSyncDateTime != *input.LastSuccessfulSyncDateTime) {
		return false
	}

	if p.LastSyncErrorCode != nil && (input.LastSyncErrorCode == nil || *p.LastSyncErrorCode != *input.LastSyncErrorCode) {
		return false
	}

	if p.LastSyncTriggeredDateTime != nil && (input.LastSyncTriggeredDateTime == nil || *p.LastSyncTriggeredDateTime != *input.LastSyncTriggeredDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ShareTokenWithSchoolDataSyncService != nil && (input.ShareTokenWithSchoolDataSyncService == nil || *p.ShareTokenWithSchoolDataSyncService != *input.ShareTokenWithSchoolDataSyncService) {
		return false
	}

	if p.SyncedDeviceCount != nil && (input.SyncedDeviceCount == nil || *p.SyncedDeviceCount != *input.SyncedDeviceCount) {
		return false
	}

	if p.TokenExpirationDateTime != nil && (input.TokenExpirationDateTime == nil || *p.TokenExpirationDateTime != *input.TokenExpirationDateTime) {
		return false
	}

	if p.TokenName != nil && (input.TokenName == nil || *p.TokenName != *input.TokenName) {
		return false
	}

	return true
}
