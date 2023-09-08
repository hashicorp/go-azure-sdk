package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType string

const (
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedefaultLimit                                          DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "DefaultLimit"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions                           DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "DefaultPlatformRestrictions"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "DefaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness                        DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "DefaultWindowsHelloForBusiness"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration              DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "DeviceComanagementAuthorityConfiguration"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration                  DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "EnrollmentNotificationsConfiguration"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypelimit                                                 DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "Limit"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions                                  DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "PlatformRestrictions"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction                             DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "SinglePlatformRestriction"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypeunknown                                               DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "Unknown"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration        DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "Windows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness                               DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType = "WindowsHelloForBusiness"
)

func PossibleValuesForDeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedefaultLimit),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypelimit),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypeunknown),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness),
	}
}

func parseDeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType(input string) (*DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedefaultLimit,
		"defaultplatformrestrictions": DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration,
		"limit":                                                 DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypelimit,
		"platformrestrictions":                                  DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions,
		"singleplatformrestriction":                             DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction,
		"unknown":                                               DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypeunknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}
