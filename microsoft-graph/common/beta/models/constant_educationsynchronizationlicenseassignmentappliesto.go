package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationLicenseAssignmentAppliesTo string

const (
	EducationSynchronizationLicenseAssignmentAppliesTofaculty EducationSynchronizationLicenseAssignmentAppliesTo = "Faculty"
	EducationSynchronizationLicenseAssignmentAppliesTonone    EducationSynchronizationLicenseAssignmentAppliesTo = "None"
	EducationSynchronizationLicenseAssignmentAppliesTostudent EducationSynchronizationLicenseAssignmentAppliesTo = "Student"
	EducationSynchronizationLicenseAssignmentAppliesToteacher EducationSynchronizationLicenseAssignmentAppliesTo = "Teacher"
)

func PossibleValuesForEducationSynchronizationLicenseAssignmentAppliesTo() []string {
	return []string{
		string(EducationSynchronizationLicenseAssignmentAppliesTofaculty),
		string(EducationSynchronizationLicenseAssignmentAppliesTonone),
		string(EducationSynchronizationLicenseAssignmentAppliesTostudent),
		string(EducationSynchronizationLicenseAssignmentAppliesToteacher),
	}
}

func parseEducationSynchronizationLicenseAssignmentAppliesTo(input string) (*EducationSynchronizationLicenseAssignmentAppliesTo, error) {
	vals := map[string]EducationSynchronizationLicenseAssignmentAppliesTo{
		"faculty": EducationSynchronizationLicenseAssignmentAppliesTofaculty,
		"none":    EducationSynchronizationLicenseAssignmentAppliesTonone,
		"student": EducationSynchronizationLicenseAssignmentAppliesTostudent,
		"teacher": EducationSynchronizationLicenseAssignmentAppliesToteacher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationSynchronizationLicenseAssignmentAppliesTo(input)
	return &out, nil
}
