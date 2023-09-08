package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EditionUpgradeConfigurationTargetEdition string

const (
	EditionUpgradeConfigurationTargetEditionwindows10Education                EditionUpgradeConfigurationTargetEdition = "Windows10Education"
	EditionUpgradeConfigurationTargetEditionwindows10EducationN               EditionUpgradeConfigurationTargetEdition = "Windows10EducationN"
	EditionUpgradeConfigurationTargetEditionwindows10Enterprise               EditionUpgradeConfigurationTargetEdition = "Windows10Enterprise"
	EditionUpgradeConfigurationTargetEditionwindows10EnterpriseN              EditionUpgradeConfigurationTargetEdition = "Windows10EnterpriseN"
	EditionUpgradeConfigurationTargetEditionwindows10HolographicEnterprise    EditionUpgradeConfigurationTargetEdition = "Windows10HolographicEnterprise"
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
		string(EditionUpgradeConfigurationTargetEditionwindows10Education),
		string(EditionUpgradeConfigurationTargetEditionwindows10EducationN),
		string(EditionUpgradeConfigurationTargetEditionwindows10Enterprise),
		string(EditionUpgradeConfigurationTargetEditionwindows10EnterpriseN),
		string(EditionUpgradeConfigurationTargetEditionwindows10HolographicEnterprise),
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
		"windows10education":                EditionUpgradeConfigurationTargetEditionwindows10Education,
		"windows10educationn":               EditionUpgradeConfigurationTargetEditionwindows10EducationN,
		"windows10enterprise":               EditionUpgradeConfigurationTargetEditionwindows10Enterprise,
		"windows10enterprisen":              EditionUpgradeConfigurationTargetEditionwindows10EnterpriseN,
		"windows10holographicenterprise":    EditionUpgradeConfigurationTargetEditionwindows10HolographicEnterprise,
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
