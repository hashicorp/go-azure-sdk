package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentConfigurationAssignmentSource string

const (
	EnrollmentConfigurationAssignmentSourcedirect     EnrollmentConfigurationAssignmentSource = "Direct"
	EnrollmentConfigurationAssignmentSourcepolicySets EnrollmentConfigurationAssignmentSource = "PolicySets"
)

func PossibleValuesForEnrollmentConfigurationAssignmentSource() []string {
	return []string{
		string(EnrollmentConfigurationAssignmentSourcedirect),
		string(EnrollmentConfigurationAssignmentSourcepolicySets),
	}
}

func parseEnrollmentConfigurationAssignmentSource(input string) (*EnrollmentConfigurationAssignmentSource, error) {
	vals := map[string]EnrollmentConfigurationAssignmentSource{
		"direct":     EnrollmentConfigurationAssignmentSourcedirect,
		"policysets": EnrollmentConfigurationAssignmentSourcepolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentConfigurationAssignmentSource(input)
	return &out, nil
}
