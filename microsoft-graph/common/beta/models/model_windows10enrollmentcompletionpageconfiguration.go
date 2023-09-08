package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EnrollmentCompletionPageConfiguration struct {
	AllowDeviceResetOnInstallFailure        *bool                                                                            `json:"allowDeviceResetOnInstallFailure,omitempty"`
	AllowDeviceUseOnInstallFailure          *bool                                                                            `json:"allowDeviceUseOnInstallFailure,omitempty"`
	AllowLogCollectionOnInstallFailure      *bool                                                                            `json:"allowLogCollectionOnInstallFailure,omitempty"`
	AllowNonBlockingAppInstallation         *bool                                                                            `json:"allowNonBlockingAppInstallation,omitempty"`
	Assignments                             *[]EnrollmentConfigurationAssignment                                             `json:"assignments,omitempty"`
	BlockDeviceSetupRetryByUser             *bool                                                                            `json:"blockDeviceSetupRetryByUser,omitempty"`
	CreatedDateTime                         *string                                                                          `json:"createdDateTime,omitempty"`
	CustomErrorMessage                      *string                                                                          `json:"customErrorMessage,omitempty"`
	Description                             *string                                                                          `json:"description,omitempty"`
	DeviceEnrollmentConfigurationType       *Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType `json:"deviceEnrollmentConfigurationType,omitempty"`
	DisableUserStatusTrackingAfterFirstUser *bool                                                                            `json:"disableUserStatusTrackingAfterFirstUser,omitempty"`
	DisplayName                             *string                                                                          `json:"displayName,omitempty"`
	Id                                      *string                                                                          `json:"id,omitempty"`
	InstallProgressTimeoutInMinutes         *int64                                                                           `json:"installProgressTimeoutInMinutes,omitempty"`
	InstallQualityUpdates                   *bool                                                                            `json:"installQualityUpdates,omitempty"`
	LastModifiedDateTime                    *string                                                                          `json:"lastModifiedDateTime,omitempty"`
	ODataType                               *string                                                                          `json:"@odata.type,omitempty"`
	Priority                                *int64                                                                           `json:"priority,omitempty"`
	RoleScopeTagIds                         *[]string                                                                        `json:"roleScopeTagIds,omitempty"`
	SelectedMobileAppIds                    *[]string                                                                        `json:"selectedMobileAppIds,omitempty"`
	ShowInstallationProgress                *bool                                                                            `json:"showInstallationProgress,omitempty"`
	TrackInstallProgressForAutopilotOnly    *bool                                                                            `json:"trackInstallProgressForAutopilotOnly,omitempty"`
	Version                                 *int64                                                                           `json:"version,omitempty"`
}
