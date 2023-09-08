package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType string

const (
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedefaultLimit                                          DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "DefaultLimit"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions                           DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "DefaultPlatformRestrictions"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "DefaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness                        DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "DefaultWindowsHelloForBusiness"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration              DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "DeviceComanagementAuthorityConfiguration"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration                  DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "EnrollmentNotificationsConfiguration"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypelimit                                                 DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "Limit"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions                                  DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "PlatformRestrictions"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction                             DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "SinglePlatformRestriction"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypeunknown                                               DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "Unknown"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration        DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "Windows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness                               DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "WindowsHelloForBusiness"
)

func PossibleValuesForDeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedefaultLimit),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypelimit),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypeunknown),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness),
	}
}

func parseDeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType(input string) (*DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedefaultLimit,
		"defaultplatformrestrictions": DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration,
		"limit":                                                 DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypelimit,
		"platformrestrictions":                                  DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions,
		"singleplatformrestriction":                             DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction,
		"unknown":                                               DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypeunknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}
