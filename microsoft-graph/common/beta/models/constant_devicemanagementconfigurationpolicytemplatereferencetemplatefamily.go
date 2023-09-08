package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily string

const (
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyappQuietTime                                 DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "AppQuietTime"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilybaseline                                     DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "Baseline"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilydeviceConfigurationPolicies                  DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "DeviceConfigurationPolicies"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilydeviceConfigurationScripts                   DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "DeviceConfigurationScripts"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityAccountProtection            DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "EndpointSecurityAccountProtection"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityAntivirus                    DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "EndpointSecurityAntivirus"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityApplicationControl           DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "EndpointSecurityApplicationControl"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityAttackSurfaceReduction       DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "EndpointSecurityAttackSurfaceReduction"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityDiskEncryption               DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "EndpointSecurityDiskEncryption"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityEndpointDetectionAndResponse DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "EndpointSecurityEndpointDetectionAndResponse"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityEndpointPrivilegeManagement  DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "EndpointSecurityEndpointPrivilegeManagement"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityFirewall                     DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "EndpointSecurityFirewall"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyenrollmentConfiguration                      DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "EnrollmentConfiguration"
	DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilynone                                         DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily = "None"
)

func PossibleValuesForDeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyappQuietTime),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilybaseline),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilydeviceConfigurationPolicies),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilydeviceConfigurationScripts),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityAccountProtection),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityAntivirus),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityApplicationControl),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityAttackSurfaceReduction),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityDiskEncryption),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityEndpointDetectionAndResponse),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityEndpointPrivilegeManagement),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityFirewall),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyenrollmentConfiguration),
		string(DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilynone),
	}
}

func parseDeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily(input string) (*DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily, error) {
	vals := map[string]DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily{
		"appquiettime":                                 DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyappQuietTime,
		"baseline":                                     DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilybaseline,
		"deviceconfigurationpolicies":                  DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilydeviceConfigurationPolicies,
		"deviceconfigurationscripts":                   DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilydeviceConfigurationScripts,
		"endpointsecurityaccountprotection":            DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityAccountProtection,
		"endpointsecurityantivirus":                    DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityAntivirus,
		"endpointsecurityapplicationcontrol":           DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityApplicationControl,
		"endpointsecurityattacksurfacereduction":       DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityAttackSurfaceReduction,
		"endpointsecuritydiskencryption":               DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityDiskEncryption,
		"endpointsecurityendpointdetectionandresponse": DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityEndpointDetectionAndResponse,
		"endpointsecurityendpointprivilegemanagement":  DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityEndpointPrivilegeManagement,
		"endpointsecurityfirewall":                     DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyendpointSecurityFirewall,
		"enrollmentconfiguration":                      DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilyenrollmentConfiguration,
		"none":                                         DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamilynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily(input)
	return &out, nil
}
