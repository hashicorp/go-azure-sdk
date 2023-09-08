package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationIdentityMatchingOptionsAppliesTo string

const (
	EducationIdentityMatchingOptionsAppliesTofaculty EducationIdentityMatchingOptionsAppliesTo = "Faculty"
	EducationIdentityMatchingOptionsAppliesTonone    EducationIdentityMatchingOptionsAppliesTo = "None"
	EducationIdentityMatchingOptionsAppliesTostudent EducationIdentityMatchingOptionsAppliesTo = "Student"
	EducationIdentityMatchingOptionsAppliesToteacher EducationIdentityMatchingOptionsAppliesTo = "Teacher"
)

func PossibleValuesForEducationIdentityMatchingOptionsAppliesTo() []string {
	return []string{
		string(EducationIdentityMatchingOptionsAppliesTofaculty),
		string(EducationIdentityMatchingOptionsAppliesTonone),
		string(EducationIdentityMatchingOptionsAppliesTostudent),
		string(EducationIdentityMatchingOptionsAppliesToteacher),
	}
}

func parseEducationIdentityMatchingOptionsAppliesTo(input string) (*EducationIdentityMatchingOptionsAppliesTo, error) {
	vals := map[string]EducationIdentityMatchingOptionsAppliesTo{
		"faculty": EducationIdentityMatchingOptionsAppliesTofaculty,
		"none":    EducationIdentityMatchingOptionsAppliesTonone,
		"student": EducationIdentityMatchingOptionsAppliesTostudent,
		"teacher": EducationIdentityMatchingOptionsAppliesToteacher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationIdentityMatchingOptionsAppliesTo(input)
	return &out, nil
}
