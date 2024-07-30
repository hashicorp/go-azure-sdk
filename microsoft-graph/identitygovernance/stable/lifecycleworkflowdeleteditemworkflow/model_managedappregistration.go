package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppRegistration struct {
	AppIdentifier        *MobileAppIdentifier                  `json:"appIdentifier,omitempty"`
	ApplicationVersion   *string                               `json:"applicationVersion,omitempty"`
	AppliedPolicies      *[]ManagedAppPolicy                   `json:"appliedPolicies,omitempty"`
	CreatedDateTime      *string                               `json:"createdDateTime,omitempty"`
	DeviceName           *string                               `json:"deviceName,omitempty"`
	DeviceTag            *string                               `json:"deviceTag,omitempty"`
	DeviceType           *string                               `json:"deviceType,omitempty"`
	FlaggedReasons       *ManagedAppRegistrationFlaggedReasons `json:"flaggedReasons,omitempty"`
	Id                   *string                               `json:"id,omitempty"`
	IntendedPolicies     *[]ManagedAppPolicy                   `json:"intendedPolicies,omitempty"`
	LastSyncDateTime     *string                               `json:"lastSyncDateTime,omitempty"`
	ManagementSdkVersion *string                               `json:"managementSdkVersion,omitempty"`
	ODataType            *string                               `json:"@odata.type,omitempty"`
	Operations           *[]ManagedAppOperation                `json:"operations,omitempty"`
	PlatformVersion      *string                               `json:"platformVersion,omitempty"`
	UserId               *string                               `json:"userId,omitempty"`
	Version              *string                               `json:"version,omitempty"`
}
