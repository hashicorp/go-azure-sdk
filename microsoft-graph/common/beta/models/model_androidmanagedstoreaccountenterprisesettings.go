package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAccountEnterpriseSettings struct {
	AndroidDeviceOwnerFullyManagedEnrollmentEnabled *bool                                                          `json:"androidDeviceOwnerFullyManagedEnrollmentEnabled,omitempty"`
	BindStatus                                      *AndroidManagedStoreAccountEnterpriseSettingsBindStatus        `json:"bindStatus,omitempty"`
	CompanyCodes                                    *[]AndroidEnrollmentCompanyCode                                `json:"companyCodes,omitempty"`
	DeviceOwnerManagementEnabled                    *bool                                                          `json:"deviceOwnerManagementEnabled,omitempty"`
	EnrollmentTarget                                *AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget  `json:"enrollmentTarget,omitempty"`
	Id                                              *string                                                        `json:"id,omitempty"`
	LastAppSyncDateTime                             *string                                                        `json:"lastAppSyncDateTime,omitempty"`
	LastAppSyncStatus                               *AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus `json:"lastAppSyncStatus,omitempty"`
	LastModifiedDateTime                            *string                                                        `json:"lastModifiedDateTime,omitempty"`
	ManagedGooglePlayInitialScopeTagIds             *[]string                                                      `json:"managedGooglePlayInitialScopeTagIds,omitempty"`
	ODataType                                       *string                                                        `json:"@odata.type,omitempty"`
	OwnerOrganizationName                           *string                                                        `json:"ownerOrganizationName,omitempty"`
	OwnerUserPrincipalName                          *string                                                        `json:"ownerUserPrincipalName,omitempty"`
	TargetGroupIds                                  *[]string                                                      `json:"targetGroupIds,omitempty"`
}
