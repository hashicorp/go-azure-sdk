package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadByFilterPayloadType string

const (
	PayloadByFilterPayloadTypeandroidEnterpriseApp                            PayloadByFilterPayloadType = "AndroidEnterpriseApp"
	PayloadByFilterPayloadTypeandroidEnterpriseConfiguration                  PayloadByFilterPayloadType = "AndroidEnterpriseConfiguration"
	PayloadByFilterPayloadTypeapplication                                     PayloadByFilterPayloadType = "Application"
	PayloadByFilterPayloadTypedeviceConfigurationAndCompliance                PayloadByFilterPayloadType = "DeviceConfigurationAndCompliance"
	PayloadByFilterPayloadTypedeviceFirmwareConfigurationInterfacePolicy      PayloadByFilterPayloadType = "DeviceFirmwareConfigurationInterfacePolicy"
	PayloadByFilterPayloadTypedeviceManagmentConfigurationAndCompliancePolicy PayloadByFilterPayloadType = "DeviceManagmentConfigurationAndCompliancePolicy"
	PayloadByFilterPayloadTypeenrollmentConfiguration                         PayloadByFilterPayloadType = "EnrollmentConfiguration"
	PayloadByFilterPayloadTypegroupPolicyConfiguration                        PayloadByFilterPayloadType = "GroupPolicyConfiguration"
	PayloadByFilterPayloadTyperesourceAccessPolicy                            PayloadByFilterPayloadType = "ResourceAccessPolicy"
	PayloadByFilterPayloadTypeunknown                                         PayloadByFilterPayloadType = "Unknown"
	PayloadByFilterPayloadTypewin32app                                        PayloadByFilterPayloadType = "Win32app"
	PayloadByFilterPayloadTypezeroTouchDeploymentDeviceConfigProfile          PayloadByFilterPayloadType = "ZeroTouchDeploymentDeviceConfigProfile"
)

func PossibleValuesForPayloadByFilterPayloadType() []string {
	return []string{
		string(PayloadByFilterPayloadTypeandroidEnterpriseApp),
		string(PayloadByFilterPayloadTypeandroidEnterpriseConfiguration),
		string(PayloadByFilterPayloadTypeapplication),
		string(PayloadByFilterPayloadTypedeviceConfigurationAndCompliance),
		string(PayloadByFilterPayloadTypedeviceFirmwareConfigurationInterfacePolicy),
		string(PayloadByFilterPayloadTypedeviceManagmentConfigurationAndCompliancePolicy),
		string(PayloadByFilterPayloadTypeenrollmentConfiguration),
		string(PayloadByFilterPayloadTypegroupPolicyConfiguration),
		string(PayloadByFilterPayloadTyperesourceAccessPolicy),
		string(PayloadByFilterPayloadTypeunknown),
		string(PayloadByFilterPayloadTypewin32app),
		string(PayloadByFilterPayloadTypezeroTouchDeploymentDeviceConfigProfile),
	}
}

func parsePayloadByFilterPayloadType(input string) (*PayloadByFilterPayloadType, error) {
	vals := map[string]PayloadByFilterPayloadType{
		"androidenterpriseapp":                            PayloadByFilterPayloadTypeandroidEnterpriseApp,
		"androidenterpriseconfiguration":                  PayloadByFilterPayloadTypeandroidEnterpriseConfiguration,
		"application":                                     PayloadByFilterPayloadTypeapplication,
		"deviceconfigurationandcompliance":                PayloadByFilterPayloadTypedeviceConfigurationAndCompliance,
		"devicefirmwareconfigurationinterfacepolicy":      PayloadByFilterPayloadTypedeviceFirmwareConfigurationInterfacePolicy,
		"devicemanagmentconfigurationandcompliancepolicy": PayloadByFilterPayloadTypedeviceManagmentConfigurationAndCompliancePolicy,
		"enrollmentconfiguration":                         PayloadByFilterPayloadTypeenrollmentConfiguration,
		"grouppolicyconfiguration":                        PayloadByFilterPayloadTypegroupPolicyConfiguration,
		"resourceaccesspolicy":                            PayloadByFilterPayloadTyperesourceAccessPolicy,
		"unknown":                                         PayloadByFilterPayloadTypeunknown,
		"win32app":                                        PayloadByFilterPayloadTypewin32app,
		"zerotouchdeploymentdeviceconfigprofile":          PayloadByFilterPayloadTypezeroTouchDeploymentDeviceConfigProfile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadByFilterPayloadType(input)
	return &out, nil
}
