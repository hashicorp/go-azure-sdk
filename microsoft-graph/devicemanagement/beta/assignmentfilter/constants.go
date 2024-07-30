package assignmentfilter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType string

const (
	DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementTypeApps    DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType = "apps"
	DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementTypeDevices DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType = "devices"
)

func PossibleValuesForDeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType() []string {
	return []string{
		string(DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementTypeApps),
		string(DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementTypeDevices),
	}
}

func (s *DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType(input string) (*DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType, error) {
	vals := map[string]DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType{
		"apps":    DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementTypeApps,
		"devices": DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementTypeDevices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType(input)
	return &out, nil
}

type DeviceAndAppManagementAssignmentFilterPlatform string

const (
	DeviceAndAppManagementAssignmentFilterPlatformAndroid                            DeviceAndAppManagementAssignmentFilterPlatform = "android"
	DeviceAndAppManagementAssignmentFilterPlatformAndroidAOSP                        DeviceAndAppManagementAssignmentFilterPlatform = "androidAOSP"
	DeviceAndAppManagementAssignmentFilterPlatformAndroidForWork                     DeviceAndAppManagementAssignmentFilterPlatform = "androidForWork"
	DeviceAndAppManagementAssignmentFilterPlatformAndroidMobileApplicationManagement DeviceAndAppManagementAssignmentFilterPlatform = "androidMobileApplicationManagement"
	DeviceAndAppManagementAssignmentFilterPlatformAndroidWorkProfile                 DeviceAndAppManagementAssignmentFilterPlatform = "androidWorkProfile"
	DeviceAndAppManagementAssignmentFilterPlatformIOS                                DeviceAndAppManagementAssignmentFilterPlatform = "iOS"
	DeviceAndAppManagementAssignmentFilterPlatformIOSMobileApplicationManagement     DeviceAndAppManagementAssignmentFilterPlatform = "iOSMobileApplicationManagement"
	DeviceAndAppManagementAssignmentFilterPlatformMacOS                              DeviceAndAppManagementAssignmentFilterPlatform = "macOS"
	DeviceAndAppManagementAssignmentFilterPlatformUnknown                            DeviceAndAppManagementAssignmentFilterPlatform = "unknown"
	DeviceAndAppManagementAssignmentFilterPlatformWindows10AndLater                  DeviceAndAppManagementAssignmentFilterPlatform = "windows10AndLater"
	DeviceAndAppManagementAssignmentFilterPlatformWindows81AndLater                  DeviceAndAppManagementAssignmentFilterPlatform = "windows81AndLater"
	DeviceAndAppManagementAssignmentFilterPlatformWindowsPhone81                     DeviceAndAppManagementAssignmentFilterPlatform = "windowsPhone81"
)

func PossibleValuesForDeviceAndAppManagementAssignmentFilterPlatform() []string {
	return []string{
		string(DeviceAndAppManagementAssignmentFilterPlatformAndroid),
		string(DeviceAndAppManagementAssignmentFilterPlatformAndroidAOSP),
		string(DeviceAndAppManagementAssignmentFilterPlatformAndroidForWork),
		string(DeviceAndAppManagementAssignmentFilterPlatformAndroidMobileApplicationManagement),
		string(DeviceAndAppManagementAssignmentFilterPlatformAndroidWorkProfile),
		string(DeviceAndAppManagementAssignmentFilterPlatformIOS),
		string(DeviceAndAppManagementAssignmentFilterPlatformIOSMobileApplicationManagement),
		string(DeviceAndAppManagementAssignmentFilterPlatformMacOS),
		string(DeviceAndAppManagementAssignmentFilterPlatformUnknown),
		string(DeviceAndAppManagementAssignmentFilterPlatformWindows10AndLater),
		string(DeviceAndAppManagementAssignmentFilterPlatformWindows81AndLater),
		string(DeviceAndAppManagementAssignmentFilterPlatformWindowsPhone81),
	}
}

func (s *DeviceAndAppManagementAssignmentFilterPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAndAppManagementAssignmentFilterPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAndAppManagementAssignmentFilterPlatform(input string) (*DeviceAndAppManagementAssignmentFilterPlatform, error) {
	vals := map[string]DeviceAndAppManagementAssignmentFilterPlatform{
		"android":                            DeviceAndAppManagementAssignmentFilterPlatformAndroid,
		"androidaosp":                        DeviceAndAppManagementAssignmentFilterPlatformAndroidAOSP,
		"androidforwork":                     DeviceAndAppManagementAssignmentFilterPlatformAndroidForWork,
		"androidmobileapplicationmanagement": DeviceAndAppManagementAssignmentFilterPlatformAndroidMobileApplicationManagement,
		"androidworkprofile":                 DeviceAndAppManagementAssignmentFilterPlatformAndroidWorkProfile,
		"ios":                                DeviceAndAppManagementAssignmentFilterPlatformIOS,
		"iosmobileapplicationmanagement":     DeviceAndAppManagementAssignmentFilterPlatformIOSMobileApplicationManagement,
		"macos":                              DeviceAndAppManagementAssignmentFilterPlatformMacOS,
		"unknown":                            DeviceAndAppManagementAssignmentFilterPlatformUnknown,
		"windows10andlater":                  DeviceAndAppManagementAssignmentFilterPlatformWindows10AndLater,
		"windows81andlater":                  DeviceAndAppManagementAssignmentFilterPlatformWindows81AndLater,
		"windowsphone81":                     DeviceAndAppManagementAssignmentFilterPlatformWindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementAssignmentFilterPlatform(input)
	return &out, nil
}

type PayloadByFilterAssignmentFilterType string

const (
	PayloadByFilterAssignmentFilterTypeExclude PayloadByFilterAssignmentFilterType = "exclude"
	PayloadByFilterAssignmentFilterTypeInclude PayloadByFilterAssignmentFilterType = "include"
	PayloadByFilterAssignmentFilterTypeNone    PayloadByFilterAssignmentFilterType = "none"
)

func PossibleValuesForPayloadByFilterAssignmentFilterType() []string {
	return []string{
		string(PayloadByFilterAssignmentFilterTypeExclude),
		string(PayloadByFilterAssignmentFilterTypeInclude),
		string(PayloadByFilterAssignmentFilterTypeNone),
	}
}

func (s *PayloadByFilterAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadByFilterAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadByFilterAssignmentFilterType(input string) (*PayloadByFilterAssignmentFilterType, error) {
	vals := map[string]PayloadByFilterAssignmentFilterType{
		"exclude": PayloadByFilterAssignmentFilterTypeExclude,
		"include": PayloadByFilterAssignmentFilterTypeInclude,
		"none":    PayloadByFilterAssignmentFilterTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadByFilterAssignmentFilterType(input)
	return &out, nil
}

type PayloadByFilterPayloadType string

const (
	PayloadByFilterPayloadTypeAndroidEnterpriseApp                            PayloadByFilterPayloadType = "androidEnterpriseApp"
	PayloadByFilterPayloadTypeAndroidEnterpriseConfiguration                  PayloadByFilterPayloadType = "androidEnterpriseConfiguration"
	PayloadByFilterPayloadTypeApplication                                     PayloadByFilterPayloadType = "application"
	PayloadByFilterPayloadTypeDeviceConfigurationAndCompliance                PayloadByFilterPayloadType = "deviceConfigurationAndCompliance"
	PayloadByFilterPayloadTypeDeviceFirmwareConfigurationInterfacePolicy      PayloadByFilterPayloadType = "deviceFirmwareConfigurationInterfacePolicy"
	PayloadByFilterPayloadTypeDeviceManagmentConfigurationAndCompliancePolicy PayloadByFilterPayloadType = "deviceManagmentConfigurationAndCompliancePolicy"
	PayloadByFilterPayloadTypeEnrollmentConfiguration                         PayloadByFilterPayloadType = "enrollmentConfiguration"
	PayloadByFilterPayloadTypeGroupPolicyConfiguration                        PayloadByFilterPayloadType = "groupPolicyConfiguration"
	PayloadByFilterPayloadTypeResourceAccessPolicy                            PayloadByFilterPayloadType = "resourceAccessPolicy"
	PayloadByFilterPayloadTypeUnknown                                         PayloadByFilterPayloadType = "unknown"
	PayloadByFilterPayloadTypeWin32app                                        PayloadByFilterPayloadType = "win32app"
	PayloadByFilterPayloadTypeZeroTouchDeploymentDeviceConfigProfile          PayloadByFilterPayloadType = "zeroTouchDeploymentDeviceConfigProfile"
)

func PossibleValuesForPayloadByFilterPayloadType() []string {
	return []string{
		string(PayloadByFilterPayloadTypeAndroidEnterpriseApp),
		string(PayloadByFilterPayloadTypeAndroidEnterpriseConfiguration),
		string(PayloadByFilterPayloadTypeApplication),
		string(PayloadByFilterPayloadTypeDeviceConfigurationAndCompliance),
		string(PayloadByFilterPayloadTypeDeviceFirmwareConfigurationInterfacePolicy),
		string(PayloadByFilterPayloadTypeDeviceManagmentConfigurationAndCompliancePolicy),
		string(PayloadByFilterPayloadTypeEnrollmentConfiguration),
		string(PayloadByFilterPayloadTypeGroupPolicyConfiguration),
		string(PayloadByFilterPayloadTypeResourceAccessPolicy),
		string(PayloadByFilterPayloadTypeUnknown),
		string(PayloadByFilterPayloadTypeWin32app),
		string(PayloadByFilterPayloadTypeZeroTouchDeploymentDeviceConfigProfile),
	}
}

func (s *PayloadByFilterPayloadType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadByFilterPayloadType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadByFilterPayloadType(input string) (*PayloadByFilterPayloadType, error) {
	vals := map[string]PayloadByFilterPayloadType{
		"androidenterpriseapp":                            PayloadByFilterPayloadTypeAndroidEnterpriseApp,
		"androidenterpriseconfiguration":                  PayloadByFilterPayloadTypeAndroidEnterpriseConfiguration,
		"application":                                     PayloadByFilterPayloadTypeApplication,
		"deviceconfigurationandcompliance":                PayloadByFilterPayloadTypeDeviceConfigurationAndCompliance,
		"devicefirmwareconfigurationinterfacepolicy":      PayloadByFilterPayloadTypeDeviceFirmwareConfigurationInterfacePolicy,
		"devicemanagmentconfigurationandcompliancepolicy": PayloadByFilterPayloadTypeDeviceManagmentConfigurationAndCompliancePolicy,
		"enrollmentconfiguration":                         PayloadByFilterPayloadTypeEnrollmentConfiguration,
		"grouppolicyconfiguration":                        PayloadByFilterPayloadTypeGroupPolicyConfiguration,
		"resourceaccesspolicy":                            PayloadByFilterPayloadTypeResourceAccessPolicy,
		"unknown":                                         PayloadByFilterPayloadTypeUnknown,
		"win32app":                                        PayloadByFilterPayloadTypeWin32app,
		"zerotouchdeploymentdeviceconfigprofile":          PayloadByFilterPayloadTypeZeroTouchDeploymentDeviceConfigProfile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadByFilterPayloadType(input)
	return &out, nil
}
