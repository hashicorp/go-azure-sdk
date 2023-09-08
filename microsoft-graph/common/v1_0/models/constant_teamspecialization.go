package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamSpecialization string

const (
	TeamSpecializationeducationClass                         TeamSpecialization = "EducationClass"
	TeamSpecializationeducationProfessionalLearningCommunity TeamSpecialization = "EducationProfessionalLearningCommunity"
	TeamSpecializationeducationStaff                         TeamSpecialization = "EducationStaff"
	TeamSpecializationeducationStandard                      TeamSpecialization = "EducationStandard"
	TeamSpecializationhealthcareCareCoordination             TeamSpecialization = "HealthcareCareCoordination"
	TeamSpecializationhealthcareStandard                     TeamSpecialization = "HealthcareStandard"
	TeamSpecializationnone                                   TeamSpecialization = "None"
)

func PossibleValuesForTeamSpecialization() []string {
	return []string{
		string(TeamSpecializationeducationClass),
		string(TeamSpecializationeducationProfessionalLearningCommunity),
		string(TeamSpecializationeducationStaff),
		string(TeamSpecializationeducationStandard),
		string(TeamSpecializationhealthcareCareCoordination),
		string(TeamSpecializationhealthcareStandard),
		string(TeamSpecializationnone),
	}
}

func parseTeamSpecialization(input string) (*TeamSpecialization, error) {
	vals := map[string]TeamSpecialization{
		"educationclass":                         TeamSpecializationeducationClass,
		"educationprofessionallearningcommunity": TeamSpecializationeducationProfessionalLearningCommunity,
		"educationstaff":                         TeamSpecializationeducationStaff,
		"educationstandard":                      TeamSpecializationeducationStandard,
		"healthcarecarecoordination":             TeamSpecializationhealthcareCareCoordination,
		"healthcarestandard":                     TeamSpecializationhealthcareStandard,
		"none":                                   TeamSpecializationnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamSpecialization(input)
	return &out, nil
}
