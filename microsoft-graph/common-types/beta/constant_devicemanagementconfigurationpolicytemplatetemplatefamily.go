package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyTemplateTemplateFamily string

const (
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_AppQuietTime                                 DeviceManagementConfigurationPolicyTemplateTemplateFamily = "appQuietTime"
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_Baseline                                     DeviceManagementConfigurationPolicyTemplateTemplateFamily = "baseline"
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_CompanyPortal                                DeviceManagementConfigurationPolicyTemplateTemplateFamily = "companyPortal"
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_DeviceConfigurationPolicies                  DeviceManagementConfigurationPolicyTemplateTemplateFamily = "deviceConfigurationPolicies"
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_DeviceConfigurationScripts                   DeviceManagementConfigurationPolicyTemplateTemplateFamily = "deviceConfigurationScripts"
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityAccountProtection            DeviceManagementConfigurationPolicyTemplateTemplateFamily = "endpointSecurityAccountProtection"
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityAntivirus                    DeviceManagementConfigurationPolicyTemplateTemplateFamily = "endpointSecurityAntivirus"
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityApplicationControl           DeviceManagementConfigurationPolicyTemplateTemplateFamily = "endpointSecurityApplicationControl"
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityAttackSurfaceReduction       DeviceManagementConfigurationPolicyTemplateTemplateFamily = "endpointSecurityAttackSurfaceReduction"
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityDiskEncryption               DeviceManagementConfigurationPolicyTemplateTemplateFamily = "endpointSecurityDiskEncryption"
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityEndpointDetectionAndResponse DeviceManagementConfigurationPolicyTemplateTemplateFamily = "endpointSecurityEndpointDetectionAndResponse"
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityEndpointPrivilegeManagement  DeviceManagementConfigurationPolicyTemplateTemplateFamily = "endpointSecurityEndpointPrivilegeManagement"
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityFirewall                     DeviceManagementConfigurationPolicyTemplateTemplateFamily = "endpointSecurityFirewall"
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_EnrollmentConfiguration                      DeviceManagementConfigurationPolicyTemplateTemplateFamily = "enrollmentConfiguration"
	DeviceManagementConfigurationPolicyTemplateTemplateFamily_None                                         DeviceManagementConfigurationPolicyTemplateTemplateFamily = "none"
)

func PossibleValuesForDeviceManagementConfigurationPolicyTemplateTemplateFamily() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_AppQuietTime),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_Baseline),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_CompanyPortal),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_DeviceConfigurationPolicies),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_DeviceConfigurationScripts),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityAccountProtection),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityAntivirus),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityApplicationControl),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityAttackSurfaceReduction),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityDiskEncryption),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityEndpointDetectionAndResponse),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityEndpointPrivilegeManagement),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityFirewall),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_EnrollmentConfiguration),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamily_None),
	}
}

func (s *DeviceManagementConfigurationPolicyTemplateTemplateFamily) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationPolicyTemplateTemplateFamily(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationPolicyTemplateTemplateFamily(input string) (*DeviceManagementConfigurationPolicyTemplateTemplateFamily, error) {
	vals := map[string]DeviceManagementConfigurationPolicyTemplateTemplateFamily{
		"appquiettime":                                 DeviceManagementConfigurationPolicyTemplateTemplateFamily_AppQuietTime,
		"baseline":                                     DeviceManagementConfigurationPolicyTemplateTemplateFamily_Baseline,
		"companyportal":                                DeviceManagementConfigurationPolicyTemplateTemplateFamily_CompanyPortal,
		"deviceconfigurationpolicies":                  DeviceManagementConfigurationPolicyTemplateTemplateFamily_DeviceConfigurationPolicies,
		"deviceconfigurationscripts":                   DeviceManagementConfigurationPolicyTemplateTemplateFamily_DeviceConfigurationScripts,
		"endpointsecurityaccountprotection":            DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityAccountProtection,
		"endpointsecurityantivirus":                    DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityAntivirus,
		"endpointsecurityapplicationcontrol":           DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityApplicationControl,
		"endpointsecurityattacksurfacereduction":       DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityAttackSurfaceReduction,
		"endpointsecuritydiskencryption":               DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityDiskEncryption,
		"endpointsecurityendpointdetectionandresponse": DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityEndpointDetectionAndResponse,
		"endpointsecurityendpointprivilegemanagement":  DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityEndpointPrivilegeManagement,
		"endpointsecurityfirewall":                     DeviceManagementConfigurationPolicyTemplateTemplateFamily_EndpointSecurityFirewall,
		"enrollmentconfiguration":                      DeviceManagementConfigurationPolicyTemplateTemplateFamily_EnrollmentConfiguration,
		"none":                                         DeviceManagementConfigurationPolicyTemplateTemplateFamily_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyTemplateTemplateFamily(input)
	return &out, nil
}
