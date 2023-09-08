package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType string

const (
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedefaultLimit                                          DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "DefaultLimit"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions                           DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "DefaultPlatformRestrictions"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "DefaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness                        DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "DefaultWindowsHelloForBusiness"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration              DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "DeviceComanagementAuthorityConfiguration"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration                  DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "EnrollmentNotificationsConfiguration"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypelimit                                                 DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "Limit"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions                                  DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "PlatformRestrictions"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction                             DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "SinglePlatformRestriction"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypeunknown                                               DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "Unknown"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration        DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "Windows10EnrollmentCompletionPageConfiguration"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness                               DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "WindowsHelloForBusiness"
)

func PossibleValuesForDeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedefaultLimit),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypelimit),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypeunknown),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness),
	}
}

func parseDeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType(input string) (*DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedefaultLimit,
		"defaultplatformrestrictions": DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration,
		"limit":                                                 DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypelimit,
		"platformrestrictions":                                  DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions,
		"singleplatformrestriction":                             DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction,
		"unknown":                                               DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypeunknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}
