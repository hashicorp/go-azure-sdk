package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EditionUpgradeConfigurationTargetEdition string

const (
	EditionUpgradeConfigurationTargetEditionnotConfigured                     EditionUpgradeConfigurationTargetEdition = "NotConfigured"
	EditionUpgradeConfigurationTargetEditionwindows10Education                EditionUpgradeConfigurationTargetEdition = "Windows10Education"
	EditionUpgradeConfigurationTargetEditionwindows10EducationN               EditionUpgradeConfigurationTargetEdition = "Windows10EducationN"
	EditionUpgradeConfigurationTargetEditionwindows10Enterprise               EditionUpgradeConfigurationTargetEdition = "Windows10Enterprise"
	EditionUpgradeConfigurationTargetEditionwindows10EnterpriseN              EditionUpgradeConfigurationTargetEdition = "Windows10EnterpriseN"
	EditionUpgradeConfigurationTargetEditionwindows10HolographicEnterprise    EditionUpgradeConfigurationTargetEdition = "Windows10HolographicEnterprise"
	EditionUpgradeConfigurationTargetEditionwindows10Home                     EditionUpgradeConfigurationTargetEdition = "Windows10Home"
	EditionUpgradeConfigurationTargetEditionwindows10HomeChina                EditionUpgradeConfigurationTargetEdition = "Windows10HomeChina"
	EditionUpgradeConfigurationTargetEditionwindows10HomeN                    EditionUpgradeConfigurationTargetEdition = "Windows10HomeN"
	EditionUpgradeConfigurationTargetEditionwindows10HomeSingleLanguage       EditionUpgradeConfigurationTargetEdition = "Windows10HomeSingleLanguage"
	EditionUpgradeConfigurationTargetEditionwindows10IoTCore                  EditionUpgradeConfigurationTargetEdition = "Windows10IoTCore"
	EditionUpgradeConfigurationTargetEditionwindows10IoTCoreCommercial        EditionUpgradeConfigurationTargetEdition = "Windows10IoTCoreCommercial"
	EditionUpgradeConfigurationTargetEditionwindows10Mobile                   EditionUpgradeConfigurationTargetEdition = "Windows10Mobile"
	EditionUpgradeConfigurationTargetEditionwindows10MobileEnterprise         EditionUpgradeConfigurationTargetEdition = "Windows10MobileEnterprise"
	EditionUpgradeConfigurationTargetEditionwindows10Professional             EditionUpgradeConfigurationTargetEdition = "Windows10Professional"
	EditionUpgradeConfigurationTargetEditionwindows10ProfessionalEducation    EditionUpgradeConfigurationTargetEdition = "Windows10ProfessionalEducation"
	EditionUpgradeConfigurationTargetEditionwindows10ProfessionalEducationN   EditionUpgradeConfigurationTargetEdition = "Windows10ProfessionalEducationN"
	EditionUpgradeConfigurationTargetEditionwindows10ProfessionalN            EditionUpgradeConfigurationTargetEdition = "Windows10ProfessionalN"
	EditionUpgradeConfigurationTargetEditionwindows10ProfessionalWorkstation  EditionUpgradeConfigurationTargetEdition = "Windows10ProfessionalWorkstation"
	EditionUpgradeConfigurationTargetEditionwindows10ProfessionalWorkstationN EditionUpgradeConfigurationTargetEdition = "Windows10ProfessionalWorkstationN"
)

func PossibleValuesForEditionUpgradeConfigurationTargetEdition() []string {
	return []string{
		string(EditionUpgradeConfigurationTargetEditionnotConfigured),
		string(EditionUpgradeConfigurationTargetEditionwindows10Education),
		string(EditionUpgradeConfigurationTargetEditionwindows10EducationN),
		string(EditionUpgradeConfigurationTargetEditionwindows10Enterprise),
		string(EditionUpgradeConfigurationTargetEditionwindows10EnterpriseN),
		string(EditionUpgradeConfigurationTargetEditionwindows10HolographicEnterprise),
		string(EditionUpgradeConfigurationTargetEditionwindows10Home),
		string(EditionUpgradeConfigurationTargetEditionwindows10HomeChina),
		string(EditionUpgradeConfigurationTargetEditionwindows10HomeN),
		string(EditionUpgradeConfigurationTargetEditionwindows10HomeSingleLanguage),
		string(EditionUpgradeConfigurationTargetEditionwindows10IoTCore),
		string(EditionUpgradeConfigurationTargetEditionwindows10IoTCoreCommercial),
		string(EditionUpgradeConfigurationTargetEditionwindows10Mobile),
		string(EditionUpgradeConfigurationTargetEditionwindows10MobileEnterprise),
		string(EditionUpgradeConfigurationTargetEditionwindows10Professional),
		string(EditionUpgradeConfigurationTargetEditionwindows10ProfessionalEducation),
		string(EditionUpgradeConfigurationTargetEditionwindows10ProfessionalEducationN),
		string(EditionUpgradeConfigurationTargetEditionwindows10ProfessionalN),
		string(EditionUpgradeConfigurationTargetEditionwindows10ProfessionalWorkstation),
		string(EditionUpgradeConfigurationTargetEditionwindows10ProfessionalWorkstationN),
	}
}

func parseEditionUpgradeConfigurationTargetEdition(input string) (*EditionUpgradeConfigurationTargetEdition, error) {
	vals := map[string]EditionUpgradeConfigurationTargetEdition{
		"notconfigured":                     EditionUpgradeConfigurationTargetEditionnotConfigured,
		"windows10education":                EditionUpgradeConfigurationTargetEditionwindows10Education,
		"windows10educationn":               EditionUpgradeConfigurationTargetEditionwindows10EducationN,
		"windows10enterprise":               EditionUpgradeConfigurationTargetEditionwindows10Enterprise,
		"windows10enterprisen":              EditionUpgradeConfigurationTargetEditionwindows10EnterpriseN,
		"windows10holographicenterprise":    EditionUpgradeConfigurationTargetEditionwindows10HolographicEnterprise,
		"windows10home":                     EditionUpgradeConfigurationTargetEditionwindows10Home,
		"windows10homechina":                EditionUpgradeConfigurationTargetEditionwindows10HomeChina,
		"windows10homen":                    EditionUpgradeConfigurationTargetEditionwindows10HomeN,
		"windows10homesinglelanguage":       EditionUpgradeConfigurationTargetEditionwindows10HomeSingleLanguage,
		"windows10iotcore":                  EditionUpgradeConfigurationTargetEditionwindows10IoTCore,
		"windows10iotcorecommercial":        EditionUpgradeConfigurationTargetEditionwindows10IoTCoreCommercial,
		"windows10mobile":                   EditionUpgradeConfigurationTargetEditionwindows10Mobile,
		"windows10mobileenterprise":         EditionUpgradeConfigurationTargetEditionwindows10MobileEnterprise,
		"windows10professional":             EditionUpgradeConfigurationTargetEditionwindows10Professional,
		"windows10professionaleducation":    EditionUpgradeConfigurationTargetEditionwindows10ProfessionalEducation,
		"windows10professionaleducationn":   EditionUpgradeConfigurationTargetEditionwindows10ProfessionalEducationN,
		"windows10professionaln":            EditionUpgradeConfigurationTargetEditionwindows10ProfessionalN,
		"windows10professionalworkstation":  EditionUpgradeConfigurationTargetEditionwindows10ProfessionalWorkstation,
		"windows10professionalworkstationn": EditionUpgradeConfigurationTargetEditionwindows10ProfessionalWorkstationN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EditionUpgradeConfigurationTargetEdition(input)
	return &out, nil
}
