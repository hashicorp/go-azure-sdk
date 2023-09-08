package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSGeneralDeviceConfiguration struct {
	Assignments                                    *[]DeviceConfigurationAssignment                     `json:"assignments,omitempty"`
	CompliantAppListType                           *MacOSGeneralDeviceConfigurationCompliantAppListType `json:"compliantAppListType,omitempty"`
	CompliantAppsList                              *[]AppListItem                                       `json:"compliantAppsList,omitempty"`
	CreatedDateTime                                *string                                              `json:"createdDateTime,omitempty"`
	Description                                    *string                                              `json:"description,omitempty"`
	DeviceSettingStateSummaries                    *[]SettingStateDeviceSummary                         `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                           *DeviceConfigurationDeviceOverview                   `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                                 *[]DeviceConfigurationDeviceStatus                   `json:"deviceStatuses,omitempty"`
	DisplayName                                    *string                                              `json:"displayName,omitempty"`
	EmailInDomainSuffixes                          *[]string                                            `json:"emailInDomainSuffixes,omitempty"`
	Id                                             *string                                              `json:"id,omitempty"`
	LastModifiedDateTime                           *string                                              `json:"lastModifiedDateTime,omitempty"`
	ODataType                                      *string                                              `json:"@odata.type,omitempty"`
	PasswordBlockSimple                            *bool                                                `json:"passwordBlockSimple,omitempty"`
	PasswordExpirationDays                         *int64                                               `json:"passwordExpirationDays,omitempty"`
	PasswordMinimumCharacterSetCount               *int64                                               `json:"passwordMinimumCharacterSetCount,omitempty"`
	PasswordMinimumLength                          *int64                                               `json:"passwordMinimumLength,omitempty"`
	PasswordMinutesOfInactivityBeforeLock          *int64                                               `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	PasswordMinutesOfInactivityBeforeScreenTimeout *int64                                               `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	PasswordPreviousPasswordBlockCount             *int64                                               `json:"passwordPreviousPasswordBlockCount,omitempty"`
	PasswordRequired                               *bool                                                `json:"passwordRequired,omitempty"`
	PasswordRequiredType                           *MacOSGeneralDeviceConfigurationPasswordRequiredType `json:"passwordRequiredType,omitempty"`
	UserStatusOverview                             *DeviceConfigurationUserOverview                     `json:"userStatusOverview,omitempty"`
	UserStatuses                                   *[]DeviceConfigurationUserStatus                     `json:"userStatuses,omitempty"`
	Version                                        *int64                                               `json:"version,omitempty"`
}
