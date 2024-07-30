package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedAppRegistration struct {
	AppIdentifier        *MobileAppIdentifier                         `json:"appIdentifier,omitempty"`
	ApplicationVersion   *string                                      `json:"applicationVersion,omitempty"`
	AppliedPolicies      *[]ManagedAppPolicy                          `json:"appliedPolicies,omitempty"`
	AzureADDeviceId      *string                                      `json:"azureADDeviceId,omitempty"`
	CreatedDateTime      *string                                      `json:"createdDateTime,omitempty"`
	DeviceManufacturer   *string                                      `json:"deviceManufacturer,omitempty"`
	DeviceModel          *string                                      `json:"deviceModel,omitempty"`
	DeviceName           *string                                      `json:"deviceName,omitempty"`
	DeviceTag            *string                                      `json:"deviceTag,omitempty"`
	DeviceType           *string                                      `json:"deviceType,omitempty"`
	FlaggedReasons       *WindowsManagedAppRegistrationFlaggedReasons `json:"flaggedReasons,omitempty"`
	Id                   *string                                      `json:"id,omitempty"`
	IntendedPolicies     *[]ManagedAppPolicy                          `json:"intendedPolicies,omitempty"`
	LastSyncDateTime     *string                                      `json:"lastSyncDateTime,omitempty"`
	ManagedDeviceId      *string                                      `json:"managedDeviceId,omitempty"`
	ManagementSdkVersion *string                                      `json:"managementSdkVersion,omitempty"`
	ODataType            *string                                      `json:"@odata.type,omitempty"`
	Operations           *[]ManagedAppOperation                       `json:"operations,omitempty"`
	PlatformVersion      *string                                      `json:"platformVersion,omitempty"`
	UserId               *string                                      `json:"userId,omitempty"`
	Version              *string                                      `json:"version,omitempty"`
}
