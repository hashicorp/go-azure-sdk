package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadByFilterPayloadType string

const (
	PayloadByFilterPayloadType_AndroidEnterpriseApp                            PayloadByFilterPayloadType = "androidEnterpriseApp"
	PayloadByFilterPayloadType_AndroidEnterpriseConfiguration                  PayloadByFilterPayloadType = "androidEnterpriseConfiguration"
	PayloadByFilterPayloadType_Application                                     PayloadByFilterPayloadType = "application"
	PayloadByFilterPayloadType_DeviceConfigurationAndCompliance                PayloadByFilterPayloadType = "deviceConfigurationAndCompliance"
	PayloadByFilterPayloadType_DeviceFirmwareConfigurationInterfacePolicy      PayloadByFilterPayloadType = "deviceFirmwareConfigurationInterfacePolicy"
	PayloadByFilterPayloadType_DeviceManagmentConfigurationAndCompliancePolicy PayloadByFilterPayloadType = "deviceManagmentConfigurationAndCompliancePolicy"
	PayloadByFilterPayloadType_EnrollmentConfiguration                         PayloadByFilterPayloadType = "enrollmentConfiguration"
	PayloadByFilterPayloadType_GroupPolicyConfiguration                        PayloadByFilterPayloadType = "groupPolicyConfiguration"
	PayloadByFilterPayloadType_ResourceAccessPolicy                            PayloadByFilterPayloadType = "resourceAccessPolicy"
	PayloadByFilterPayloadType_Unknown                                         PayloadByFilterPayloadType = "unknown"
	PayloadByFilterPayloadType_Win32app                                        PayloadByFilterPayloadType = "win32app"
	PayloadByFilterPayloadType_ZeroTouchDeploymentDeviceConfigProfile          PayloadByFilterPayloadType = "zeroTouchDeploymentDeviceConfigProfile"
)

func PossibleValuesForPayloadByFilterPayloadType() []string {
	return []string{
		string(PayloadByFilterPayloadType_AndroidEnterpriseApp),
		string(PayloadByFilterPayloadType_AndroidEnterpriseConfiguration),
		string(PayloadByFilterPayloadType_Application),
		string(PayloadByFilterPayloadType_DeviceConfigurationAndCompliance),
		string(PayloadByFilterPayloadType_DeviceFirmwareConfigurationInterfacePolicy),
		string(PayloadByFilterPayloadType_DeviceManagmentConfigurationAndCompliancePolicy),
		string(PayloadByFilterPayloadType_EnrollmentConfiguration),
		string(PayloadByFilterPayloadType_GroupPolicyConfiguration),
		string(PayloadByFilterPayloadType_ResourceAccessPolicy),
		string(PayloadByFilterPayloadType_Unknown),
		string(PayloadByFilterPayloadType_Win32app),
		string(PayloadByFilterPayloadType_ZeroTouchDeploymentDeviceConfigProfile),
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
		"androidenterpriseapp":                            PayloadByFilterPayloadType_AndroidEnterpriseApp,
		"androidenterpriseconfiguration":                  PayloadByFilterPayloadType_AndroidEnterpriseConfiguration,
		"application":                                     PayloadByFilterPayloadType_Application,
		"deviceconfigurationandcompliance":                PayloadByFilterPayloadType_DeviceConfigurationAndCompliance,
		"devicefirmwareconfigurationinterfacepolicy":      PayloadByFilterPayloadType_DeviceFirmwareConfigurationInterfacePolicy,
		"devicemanagmentconfigurationandcompliancepolicy": PayloadByFilterPayloadType_DeviceManagmentConfigurationAndCompliancePolicy,
		"enrollmentconfiguration":                         PayloadByFilterPayloadType_EnrollmentConfiguration,
		"grouppolicyconfiguration":                        PayloadByFilterPayloadType_GroupPolicyConfiguration,
		"resourceaccesspolicy":                            PayloadByFilterPayloadType_ResourceAccessPolicy,
		"unknown":                                         PayloadByFilterPayloadType_Unknown,
		"win32app":                                        PayloadByFilterPayloadType_Win32app,
		"zerotouchdeploymentdeviceconfigprofile":          PayloadByFilterPayloadType_ZeroTouchDeploymentDeviceConfigProfile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadByFilterPayloadType(input)
	return &out, nil
}
