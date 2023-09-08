package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAppManagementOperationPredicate struct {
	Id                                                        *string
	IsEnabledForMicrosoftStoreForBusiness                     *bool
	MicrosoftStoreForBusinessLanguage                         *string
	MicrosoftStoreForBusinessLastCompletedApplicationSyncTime *string
	MicrosoftStoreForBusinessLastSuccessfulSyncDateTime       *string
	ODataType                                                 *string
}

func (p DeviceAppManagementOperationPredicate) Matches(input DeviceAppManagement) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsEnabledForMicrosoftStoreForBusiness != nil && (input.IsEnabledForMicrosoftStoreForBusiness == nil || *p.IsEnabledForMicrosoftStoreForBusiness != *input.IsEnabledForMicrosoftStoreForBusiness) {
		return false
	}

	if p.MicrosoftStoreForBusinessLanguage != nil && (input.MicrosoftStoreForBusinessLanguage == nil || *p.MicrosoftStoreForBusinessLanguage != *input.MicrosoftStoreForBusinessLanguage) {
		return false
	}

	if p.MicrosoftStoreForBusinessLastCompletedApplicationSyncTime != nil && (input.MicrosoftStoreForBusinessLastCompletedApplicationSyncTime == nil || *p.MicrosoftStoreForBusinessLastCompletedApplicationSyncTime != *input.MicrosoftStoreForBusinessLastCompletedApplicationSyncTime) {
		return false
	}

	if p.MicrosoftStoreForBusinessLastSuccessfulSyncDateTime != nil && (input.MicrosoftStoreForBusinessLastSuccessfulSyncDateTime == nil || *p.MicrosoftStoreForBusinessLastSuccessfulSyncDateTime != *input.MicrosoftStoreForBusinessLastSuccessfulSyncDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
