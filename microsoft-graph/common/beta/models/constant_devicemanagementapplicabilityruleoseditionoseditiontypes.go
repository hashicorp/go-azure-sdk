package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementApplicabilityRuleOsEditionOsEditionTypes string

const (
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypesnotConfigured                     DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "NotConfigured"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Education                DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10Education"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10EducationN               DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10EducationN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Enterprise               DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10Enterprise"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10EnterpriseN              DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10EnterpriseN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10HolographicEnterprise    DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10HolographicEnterprise"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Home                     DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10Home"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10HomeChina                DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10HomeChina"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10HomeN                    DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10HomeN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10HomeSingleLanguage       DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10HomeSingleLanguage"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10IoTCore                  DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10IoTCore"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10IoTCoreCommercial        DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10IoTCoreCommercial"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Mobile                   DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10Mobile"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10MobileEnterprise         DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10MobileEnterprise"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Professional             DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10Professional"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalEducation    DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10ProfessionalEducation"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalEducationN   DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10ProfessionalEducationN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalN            DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10ProfessionalN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalWorkstation  DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10ProfessionalWorkstation"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalWorkstationN DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "Windows10ProfessionalWorkstationN"
)

func PossibleValuesForDeviceManagementApplicabilityRuleOsEditionOsEditionTypes() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypesnotConfigured),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Education),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10EducationN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Enterprise),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10EnterpriseN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10HolographicEnterprise),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Home),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10HomeChina),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10HomeN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10HomeSingleLanguage),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10IoTCore),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10IoTCoreCommercial),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Mobile),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10MobileEnterprise),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Professional),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalEducation),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalEducationN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalWorkstation),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalWorkstationN),
	}
}

func parseDeviceManagementApplicabilityRuleOsEditionOsEditionTypes(input string) (*DeviceManagementApplicabilityRuleOsEditionOsEditionTypes, error) {
	vals := map[string]DeviceManagementApplicabilityRuleOsEditionOsEditionTypes{
		"notconfigured":                     DeviceManagementApplicabilityRuleOsEditionOsEditionTypesnotConfigured,
		"windows10education":                DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Education,
		"windows10educationn":               DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10EducationN,
		"windows10enterprise":               DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Enterprise,
		"windows10enterprisen":              DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10EnterpriseN,
		"windows10holographicenterprise":    DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10HolographicEnterprise,
		"windows10home":                     DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Home,
		"windows10homechina":                DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10HomeChina,
		"windows10homen":                    DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10HomeN,
		"windows10homesinglelanguage":       DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10HomeSingleLanguage,
		"windows10iotcore":                  DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10IoTCore,
		"windows10iotcorecommercial":        DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10IoTCoreCommercial,
		"windows10mobile":                   DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Mobile,
		"windows10mobileenterprise":         DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10MobileEnterprise,
		"windows10professional":             DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10Professional,
		"windows10professionaleducation":    DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalEducation,
		"windows10professionaleducationn":   DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalEducationN,
		"windows10professionaln":            DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalN,
		"windows10professionalworkstation":  DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalWorkstation,
		"windows10professionalworkstationn": DeviceManagementApplicabilityRuleOsEditionOsEditionTypeswindows10ProfessionalWorkstationN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleOsEditionOsEditionTypes(input)
	return &out, nil
}
