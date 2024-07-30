package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EditionUpgradeConfigurationTargetEdition string

const (
	EditionUpgradeConfigurationTargetEdition_NotConfigured                     EditionUpgradeConfigurationTargetEdition = "notConfigured"
	EditionUpgradeConfigurationTargetEdition_Windows10Education                EditionUpgradeConfigurationTargetEdition = "windows10Education"
	EditionUpgradeConfigurationTargetEdition_Windows10EducationN               EditionUpgradeConfigurationTargetEdition = "windows10EducationN"
	EditionUpgradeConfigurationTargetEdition_Windows10Enterprise               EditionUpgradeConfigurationTargetEdition = "windows10Enterprise"
	EditionUpgradeConfigurationTargetEdition_Windows10EnterpriseN              EditionUpgradeConfigurationTargetEdition = "windows10EnterpriseN"
	EditionUpgradeConfigurationTargetEdition_Windows10HolographicEnterprise    EditionUpgradeConfigurationTargetEdition = "windows10HolographicEnterprise"
	EditionUpgradeConfigurationTargetEdition_Windows10Home                     EditionUpgradeConfigurationTargetEdition = "windows10Home"
	EditionUpgradeConfigurationTargetEdition_Windows10HomeChina                EditionUpgradeConfigurationTargetEdition = "windows10HomeChina"
	EditionUpgradeConfigurationTargetEdition_Windows10HomeN                    EditionUpgradeConfigurationTargetEdition = "windows10HomeN"
	EditionUpgradeConfigurationTargetEdition_Windows10HomeSingleLanguage       EditionUpgradeConfigurationTargetEdition = "windows10HomeSingleLanguage"
	EditionUpgradeConfigurationTargetEdition_Windows10IoTCore                  EditionUpgradeConfigurationTargetEdition = "windows10IoTCore"
	EditionUpgradeConfigurationTargetEdition_Windows10IoTCoreCommercial        EditionUpgradeConfigurationTargetEdition = "windows10IoTCoreCommercial"
	EditionUpgradeConfigurationTargetEdition_Windows10Mobile                   EditionUpgradeConfigurationTargetEdition = "windows10Mobile"
	EditionUpgradeConfigurationTargetEdition_Windows10MobileEnterprise         EditionUpgradeConfigurationTargetEdition = "windows10MobileEnterprise"
	EditionUpgradeConfigurationTargetEdition_Windows10Professional             EditionUpgradeConfigurationTargetEdition = "windows10Professional"
	EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalEducation    EditionUpgradeConfigurationTargetEdition = "windows10ProfessionalEducation"
	EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalEducationN   EditionUpgradeConfigurationTargetEdition = "windows10ProfessionalEducationN"
	EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalN            EditionUpgradeConfigurationTargetEdition = "windows10ProfessionalN"
	EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalWorkstation  EditionUpgradeConfigurationTargetEdition = "windows10ProfessionalWorkstation"
	EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalWorkstationN EditionUpgradeConfigurationTargetEdition = "windows10ProfessionalWorkstationN"
)

func PossibleValuesForEditionUpgradeConfigurationTargetEdition() []string {
	return []string{
		string(EditionUpgradeConfigurationTargetEdition_NotConfigured),
		string(EditionUpgradeConfigurationTargetEdition_Windows10Education),
		string(EditionUpgradeConfigurationTargetEdition_Windows10EducationN),
		string(EditionUpgradeConfigurationTargetEdition_Windows10Enterprise),
		string(EditionUpgradeConfigurationTargetEdition_Windows10EnterpriseN),
		string(EditionUpgradeConfigurationTargetEdition_Windows10HolographicEnterprise),
		string(EditionUpgradeConfigurationTargetEdition_Windows10Home),
		string(EditionUpgradeConfigurationTargetEdition_Windows10HomeChina),
		string(EditionUpgradeConfigurationTargetEdition_Windows10HomeN),
		string(EditionUpgradeConfigurationTargetEdition_Windows10HomeSingleLanguage),
		string(EditionUpgradeConfigurationTargetEdition_Windows10IoTCore),
		string(EditionUpgradeConfigurationTargetEdition_Windows10IoTCoreCommercial),
		string(EditionUpgradeConfigurationTargetEdition_Windows10Mobile),
		string(EditionUpgradeConfigurationTargetEdition_Windows10MobileEnterprise),
		string(EditionUpgradeConfigurationTargetEdition_Windows10Professional),
		string(EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalEducation),
		string(EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalEducationN),
		string(EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalN),
		string(EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalWorkstation),
		string(EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalWorkstationN),
	}
}

func (s *EditionUpgradeConfigurationTargetEdition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEditionUpgradeConfigurationTargetEdition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEditionUpgradeConfigurationTargetEdition(input string) (*EditionUpgradeConfigurationTargetEdition, error) {
	vals := map[string]EditionUpgradeConfigurationTargetEdition{
		"notconfigured":                     EditionUpgradeConfigurationTargetEdition_NotConfigured,
		"windows10education":                EditionUpgradeConfigurationTargetEdition_Windows10Education,
		"windows10educationn":               EditionUpgradeConfigurationTargetEdition_Windows10EducationN,
		"windows10enterprise":               EditionUpgradeConfigurationTargetEdition_Windows10Enterprise,
		"windows10enterprisen":              EditionUpgradeConfigurationTargetEdition_Windows10EnterpriseN,
		"windows10holographicenterprise":    EditionUpgradeConfigurationTargetEdition_Windows10HolographicEnterprise,
		"windows10home":                     EditionUpgradeConfigurationTargetEdition_Windows10Home,
		"windows10homechina":                EditionUpgradeConfigurationTargetEdition_Windows10HomeChina,
		"windows10homen":                    EditionUpgradeConfigurationTargetEdition_Windows10HomeN,
		"windows10homesinglelanguage":       EditionUpgradeConfigurationTargetEdition_Windows10HomeSingleLanguage,
		"windows10iotcore":                  EditionUpgradeConfigurationTargetEdition_Windows10IoTCore,
		"windows10iotcorecommercial":        EditionUpgradeConfigurationTargetEdition_Windows10IoTCoreCommercial,
		"windows10mobile":                   EditionUpgradeConfigurationTargetEdition_Windows10Mobile,
		"windows10mobileenterprise":         EditionUpgradeConfigurationTargetEdition_Windows10MobileEnterprise,
		"windows10professional":             EditionUpgradeConfigurationTargetEdition_Windows10Professional,
		"windows10professionaleducation":    EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalEducation,
		"windows10professionaleducationn":   EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalEducationN,
		"windows10professionaln":            EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalN,
		"windows10professionalworkstation":  EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalWorkstation,
		"windows10professionalworkstationn": EditionUpgradeConfigurationTargetEdition_Windows10ProfessionalWorkstationN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EditionUpgradeConfigurationTargetEdition(input)
	return &out, nil
}
