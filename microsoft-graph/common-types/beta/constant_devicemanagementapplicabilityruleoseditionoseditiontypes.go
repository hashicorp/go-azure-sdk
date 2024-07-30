package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementApplicabilityRuleOsEditionOsEditionTypes string

const (
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_NotConfigured                     DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "notConfigured"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Education                DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10Education"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10EducationN               DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10EducationN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Enterprise               DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10Enterprise"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10EnterpriseN              DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10EnterpriseN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10HolographicEnterprise    DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10HolographicEnterprise"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Home                     DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10Home"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10HomeChina                DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10HomeChina"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10HomeN                    DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10HomeN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10HomeSingleLanguage       DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10HomeSingleLanguage"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10IoTCore                  DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10IoTCore"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10IoTCoreCommercial        DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10IoTCoreCommercial"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Mobile                   DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10Mobile"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10MobileEnterprise         DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10MobileEnterprise"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Professional             DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10Professional"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalEducation    DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10ProfessionalEducation"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalEducationN   DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10ProfessionalEducationN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalN            DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10ProfessionalN"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalWorkstation  DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10ProfessionalWorkstation"
	DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalWorkstationN DeviceManagementApplicabilityRuleOsEditionOsEditionTypes = "windows10ProfessionalWorkstationN"
)

func PossibleValuesForDeviceManagementApplicabilityRuleOsEditionOsEditionTypes() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_NotConfigured),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Education),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10EducationN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Enterprise),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10EnterpriseN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10HolographicEnterprise),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Home),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10HomeChina),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10HomeN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10HomeSingleLanguage),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10IoTCore),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10IoTCoreCommercial),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Mobile),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10MobileEnterprise),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Professional),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalEducation),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalEducationN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalN),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalWorkstation),
		string(DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalWorkstationN),
	}
}

func (s *DeviceManagementApplicabilityRuleOsEditionOsEditionTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementApplicabilityRuleOsEditionOsEditionTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementApplicabilityRuleOsEditionOsEditionTypes(input string) (*DeviceManagementApplicabilityRuleOsEditionOsEditionTypes, error) {
	vals := map[string]DeviceManagementApplicabilityRuleOsEditionOsEditionTypes{
		"notconfigured":                     DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_NotConfigured,
		"windows10education":                DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Education,
		"windows10educationn":               DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10EducationN,
		"windows10enterprise":               DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Enterprise,
		"windows10enterprisen":              DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10EnterpriseN,
		"windows10holographicenterprise":    DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10HolographicEnterprise,
		"windows10home":                     DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Home,
		"windows10homechina":                DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10HomeChina,
		"windows10homen":                    DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10HomeN,
		"windows10homesinglelanguage":       DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10HomeSingleLanguage,
		"windows10iotcore":                  DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10IoTCore,
		"windows10iotcorecommercial":        DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10IoTCoreCommercial,
		"windows10mobile":                   DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Mobile,
		"windows10mobileenterprise":         DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10MobileEnterprise,
		"windows10professional":             DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10Professional,
		"windows10professionaleducation":    DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalEducation,
		"windows10professionaleducationn":   DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalEducationN,
		"windows10professionaln":            DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalN,
		"windows10professionalworkstation":  DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalWorkstation,
		"windows10professionalworkstationn": DeviceManagementApplicabilityRuleOsEditionOsEditionTypes_Windows10ProfessionalWorkstationN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleOsEditionOsEditionTypes(input)
	return &out, nil
}
