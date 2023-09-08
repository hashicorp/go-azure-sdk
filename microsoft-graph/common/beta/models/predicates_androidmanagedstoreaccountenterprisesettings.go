package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAccountEnterpriseSettingsOperationPredicate struct {
	AndroidDeviceOwnerFullyManagedEnrollmentEnabled *bool
	DeviceOwnerManagementEnabled                    *bool
	Id                                              *string
	LastAppSyncDateTime                             *string
	LastModifiedDateTime                            *string
	ODataType                                       *string
	OwnerOrganizationName                           *string
	OwnerUserPrincipalName                          *string
}

func (p AndroidManagedStoreAccountEnterpriseSettingsOperationPredicate) Matches(input AndroidManagedStoreAccountEnterpriseSettings) bool {

	if p.AndroidDeviceOwnerFullyManagedEnrollmentEnabled != nil && (input.AndroidDeviceOwnerFullyManagedEnrollmentEnabled == nil || *p.AndroidDeviceOwnerFullyManagedEnrollmentEnabled != *input.AndroidDeviceOwnerFullyManagedEnrollmentEnabled) {
		return false
	}

	if p.DeviceOwnerManagementEnabled != nil && (input.DeviceOwnerManagementEnabled == nil || *p.DeviceOwnerManagementEnabled != *input.DeviceOwnerManagementEnabled) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastAppSyncDateTime != nil && (input.LastAppSyncDateTime == nil || *p.LastAppSyncDateTime != *input.LastAppSyncDateTime) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OwnerOrganizationName != nil && (input.OwnerOrganizationName == nil || *p.OwnerOrganizationName != *input.OwnerOrganizationName) {
		return false
	}

	if p.OwnerUserPrincipalName != nil && (input.OwnerUserPrincipalName == nil || *p.OwnerUserPrincipalName != *input.OwnerUserPrincipalName) {
		return false
	}

	return true
}
