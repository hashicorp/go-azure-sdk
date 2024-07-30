package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily string

const (
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_AppQuietTime                                 DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "appQuietTime"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_Baseline                                     DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "baseline"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_CompanyPortal                                DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "companyPortal"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_DeviceConfigurationPolicies                  DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "deviceConfigurationPolicies"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_DeviceConfigurationScripts                   DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "deviceConfigurationScripts"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityAccountProtection            DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "endpointSecurityAccountProtection"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityAntivirus                    DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "endpointSecurityAntivirus"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityApplicationControl           DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "endpointSecurityApplicationControl"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityAttackSurfaceReduction       DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "endpointSecurityAttackSurfaceReduction"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityDiskEncryption               DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "endpointSecurityDiskEncryption"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityEndpointDetectionAndResponse DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "endpointSecurityEndpointDetectionAndResponse"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityEndpointPrivilegeManagement  DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "endpointSecurityEndpointPrivilegeManagement"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityFirewall                     DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "endpointSecurityFirewall"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EnrollmentConfiguration                      DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "enrollmentConfiguration"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_None                                         DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "none"
)

func PossibleValuesForDeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_AppQuietTime),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_Baseline),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_CompanyPortal),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_DeviceConfigurationPolicies),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_DeviceConfigurationScripts),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityAccountProtection),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityAntivirus),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityApplicationControl),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityAttackSurfaceReduction),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityDiskEncryption),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityEndpointDetectionAndResponse),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityEndpointPrivilegeManagement),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityFirewall),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EnrollmentConfiguration),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_None),
	}
}

func (s *DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily(input string) (*DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily, error) {
	vals := map[string]DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily{
		"appquiettime":                                 DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_AppQuietTime,
		"baseline":                                     DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_Baseline,
		"companyportal":                                DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_CompanyPortal,
		"deviceconfigurationpolicies":                  DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_DeviceConfigurationPolicies,
		"deviceconfigurationscripts":                   DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_DeviceConfigurationScripts,
		"endpointsecurityaccountprotection":            DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityAccountProtection,
		"endpointsecurityantivirus":                    DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityAntivirus,
		"endpointsecurityapplicationcontrol":           DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityApplicationControl,
		"endpointsecurityattacksurfacereduction":       DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityAttackSurfaceReduction,
		"endpointsecuritydiskencryption":               DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityDiskEncryption,
		"endpointsecurityendpointdetectionandresponse": DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityEndpointDetectionAndResponse,
		"endpointsecurityendpointprivilegemanagement":  DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityEndpointPrivilegeManagement,
		"endpointsecurityfirewall":                     DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EndpointSecurityFirewall,
		"enrollmentconfiguration":                      DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_EnrollmentConfiguration,
		"none":                                         DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily(input)
	return &out, nil
}
