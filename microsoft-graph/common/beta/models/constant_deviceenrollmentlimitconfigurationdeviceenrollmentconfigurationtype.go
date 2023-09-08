package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType string

const (
	DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedefaultLimit                                          DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType = "DefaultLimit"
	DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions                           DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType = "DefaultPlatformRestrictions"
	DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType = "DefaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness                        DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType = "DefaultWindowsHelloForBusiness"
	DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration              DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType = "DeviceComanagementAuthorityConfiguration"
	DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration                  DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType = "EnrollmentNotificationsConfiguration"
	DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypelimit                                                 DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType = "Limit"
	DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions                                  DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType = "PlatformRestrictions"
	DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction                             DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType = "SinglePlatformRestriction"
	DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypeunknown                                               DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType = "Unknown"
	DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration        DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType = "Windows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness                               DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType = "WindowsHelloForBusiness"
)

func PossibleValuesForDeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedefaultLimit),
		string(DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions),
		string(DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness),
		string(DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration),
		string(DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration),
		string(DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypelimit),
		string(DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions),
		string(DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction),
		string(DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypeunknown),
		string(DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness),
	}
}

func parseDeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType(input string) (*DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedefaultLimit,
		"defaultplatformrestrictions": DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration,
		"limit":                                                 DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypelimit,
		"platformrestrictions":                                  DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions,
		"singleplatformrestriction":                             DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction,
		"unknown":                                               DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypeunknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentLimitConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}
