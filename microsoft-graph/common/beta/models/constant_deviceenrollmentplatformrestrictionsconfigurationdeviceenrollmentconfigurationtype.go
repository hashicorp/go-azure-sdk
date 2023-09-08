package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType string

const (
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedefaultLimit                                          DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "DefaultLimit"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions                           DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "DefaultPlatformRestrictions"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "DefaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness                        DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "DefaultWindowsHelloForBusiness"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration              DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "DeviceComanagementAuthorityConfiguration"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration                  DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "EnrollmentNotificationsConfiguration"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypelimit                                                 DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "Limit"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions                                  DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "PlatformRestrictions"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction                             DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "SinglePlatformRestriction"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypeunknown                                               DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "Unknown"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration        DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "Windows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness                               DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "WindowsHelloForBusiness"
)

func PossibleValuesForDeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedefaultLimit),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypelimit),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypeunknown),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness),
	}
}

func parseDeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType(input string) (*DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedefaultLimit,
		"defaultplatformrestrictions": DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration,
		"limit":                                                 DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypelimit,
		"platformrestrictions":                                  DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions,
		"singleplatformrestriction":                             DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction,
		"unknown":                                               DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypeunknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}
