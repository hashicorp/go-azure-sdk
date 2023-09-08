package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType string

const (
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedefaultLimit                                          Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "DefaultLimit"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions                           Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "DefaultPlatformRestrictions"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "DefaultWindows10EnrollmentCompletionPageConfiguration"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness                        Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "DefaultWindowsHelloForBusiness"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration              Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "DeviceComanagementAuthorityConfiguration"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration                  Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "EnrollmentNotificationsConfiguration"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypelimit                                                 Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "Limit"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions                                  Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "PlatformRestrictions"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction                             Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "SinglePlatformRestriction"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypeunknown                                               Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "Unknown"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration        Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "Windows10EnrollmentCompletionPageConfiguration"
	Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness                               Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType = "WindowsHelloForBusiness"
)

func PossibleValuesForWindows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedefaultLimit),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypelimit),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypeunknown),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration),
		string(Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness),
	}
}

func parseWindows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType(input string) (*Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedefaultLimit,
		"defaultplatformrestrictions": Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypedeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypeenrollmentNotificationsConfiguration,
		"limit":                                                 Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypelimit,
		"platformrestrictions":                                  Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypeplatformRestrictions,
		"singleplatformrestriction":                             Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypesinglePlatformRestriction,
		"unknown":                                               Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypeunknown,
		"windows10enrollmentcompletionpageconfiguration":        Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypewindows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationTypewindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EnrollmentCompletionPageConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}
