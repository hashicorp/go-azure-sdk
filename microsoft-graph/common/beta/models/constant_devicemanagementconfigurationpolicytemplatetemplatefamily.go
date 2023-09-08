package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyTemplateTemplateFamily string

const (
	DeviceManagementConfigurationPolicyTemplateTemplateFamilyappQuietTime                                 DeviceManagementConfigurationPolicyTemplateTemplateFamily = "AppQuietTime"
	DeviceManagementConfigurationPolicyTemplateTemplateFamilybaseline                                     DeviceManagementConfigurationPolicyTemplateTemplateFamily = "Baseline"
	DeviceManagementConfigurationPolicyTemplateTemplateFamilydeviceConfigurationPolicies                  DeviceManagementConfigurationPolicyTemplateTemplateFamily = "DeviceConfigurationPolicies"
	DeviceManagementConfigurationPolicyTemplateTemplateFamilydeviceConfigurationScripts                   DeviceManagementConfigurationPolicyTemplateTemplateFamily = "DeviceConfigurationScripts"
	DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityAccountProtection            DeviceManagementConfigurationPolicyTemplateTemplateFamily = "EndpointSecurityAccountProtection"
	DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityAntivirus                    DeviceManagementConfigurationPolicyTemplateTemplateFamily = "EndpointSecurityAntivirus"
	DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityApplicationControl           DeviceManagementConfigurationPolicyTemplateTemplateFamily = "EndpointSecurityApplicationControl"
	DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityAttackSurfaceReduction       DeviceManagementConfigurationPolicyTemplateTemplateFamily = "EndpointSecurityAttackSurfaceReduction"
	DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityDiskEncryption               DeviceManagementConfigurationPolicyTemplateTemplateFamily = "EndpointSecurityDiskEncryption"
	DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityEndpointDetectionAndResponse DeviceManagementConfigurationPolicyTemplateTemplateFamily = "EndpointSecurityEndpointDetectionAndResponse"
	DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityEndpointPrivilegeManagement  DeviceManagementConfigurationPolicyTemplateTemplateFamily = "EndpointSecurityEndpointPrivilegeManagement"
	DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityFirewall                     DeviceManagementConfigurationPolicyTemplateTemplateFamily = "EndpointSecurityFirewall"
	DeviceManagementConfigurationPolicyTemplateTemplateFamilyenrollmentConfiguration                      DeviceManagementConfigurationPolicyTemplateTemplateFamily = "EnrollmentConfiguration"
	DeviceManagementConfigurationPolicyTemplateTemplateFamilynone                                         DeviceManagementConfigurationPolicyTemplateTemplateFamily = "None"
)

func PossibleValuesForDeviceManagementConfigurationPolicyTemplateTemplateFamily() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamilyappQuietTime),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamilybaseline),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamilydeviceConfigurationPolicies),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamilydeviceConfigurationScripts),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityAccountProtection),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityAntivirus),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityApplicationControl),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityAttackSurfaceReduction),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityDiskEncryption),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityEndpointDetectionAndResponse),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityEndpointPrivilegeManagement),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityFirewall),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamilyenrollmentConfiguration),
		string(DeviceManagementConfigurationPolicyTemplateTemplateFamilynone),
	}
}

func parseDeviceManagementConfigurationPolicyTemplateTemplateFamily(input string) (*DeviceManagementConfigurationPolicyTemplateTemplateFamily, error) {
	vals := map[string]DeviceManagementConfigurationPolicyTemplateTemplateFamily{
		"appquiettime":                                 DeviceManagementConfigurationPolicyTemplateTemplateFamilyappQuietTime,
		"baseline":                                     DeviceManagementConfigurationPolicyTemplateTemplateFamilybaseline,
		"deviceconfigurationpolicies":                  DeviceManagementConfigurationPolicyTemplateTemplateFamilydeviceConfigurationPolicies,
		"deviceconfigurationscripts":                   DeviceManagementConfigurationPolicyTemplateTemplateFamilydeviceConfigurationScripts,
		"endpointsecurityaccountprotection":            DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityAccountProtection,
		"endpointsecurityantivirus":                    DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityAntivirus,
		"endpointsecurityapplicationcontrol":           DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityApplicationControl,
		"endpointsecurityattacksurfacereduction":       DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityAttackSurfaceReduction,
		"endpointsecuritydiskencryption":               DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityDiskEncryption,
		"endpointsecurityendpointdetectionandresponse": DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityEndpointDetectionAndResponse,
		"endpointsecurityendpointprivilegemanagement":  DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityEndpointPrivilegeManagement,
		"endpointsecurityfirewall":                     DeviceManagementConfigurationPolicyTemplateTemplateFamilyendpointSecurityFirewall,
		"enrollmentconfiguration":                      DeviceManagementConfigurationPolicyTemplateTemplateFamilyenrollmentConfiguration,
		"none":                                         DeviceManagementConfigurationPolicyTemplateTemplateFamilynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyTemplateTemplateFamily(input)
	return &out, nil
}
