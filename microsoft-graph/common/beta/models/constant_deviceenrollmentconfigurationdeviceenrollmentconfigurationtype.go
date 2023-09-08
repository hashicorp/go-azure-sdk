package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType string

const (
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedefaultLimit                                          DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "DefaultLimit"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions                           DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "DefaultPlatformRestrictions"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "DefaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness                        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "DefaultWindowsHelloForBusiness"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration              DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "DeviceComanagementAuthorityConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "EnrollmentNotificationsConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypelimit                                                 DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "Limit"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions                                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "PlatformRestrictions"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction                             DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "SinglePlatformRestriction"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeunknown                                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "Unknown"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "Windows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "WindowsHelloForBusiness"
)

func PossibleValuesForDeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedefaultLimit),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypelimit),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeunknown),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness),
	}
}

func parseDeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType(input string) (*DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedefaultLimit,
		"defaultplatformrestrictions": DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration,
		"limit":                                                 DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypelimit,
		"platformrestrictions":                                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions,
		"singleplatformrestriction":                             DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction,
		"unknown":                                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeunknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}
