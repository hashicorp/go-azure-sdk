package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationIdentityDomainAppliesTo string

const (
	EducationIdentityDomainAppliesTofaculty EducationIdentityDomainAppliesTo = "Faculty"
	EducationIdentityDomainAppliesTonone    EducationIdentityDomainAppliesTo = "None"
	EducationIdentityDomainAppliesTostudent EducationIdentityDomainAppliesTo = "Student"
	EducationIdentityDomainAppliesToteacher EducationIdentityDomainAppliesTo = "Teacher"
)

func PossibleValuesForEducationIdentityDomainAppliesTo() []string {
	return []string{
		string(EducationIdentityDomainAppliesTofaculty),
		string(EducationIdentityDomainAppliesTonone),
		string(EducationIdentityDomainAppliesTostudent),
		string(EducationIdentityDomainAppliesToteacher),
	}
}

func parseEducationIdentityDomainAppliesTo(input string) (*EducationIdentityDomainAppliesTo, error) {
	vals := map[string]EducationIdentityDomainAppliesTo{
		"faculty": EducationIdentityDomainAppliesTofaculty,
		"none":    EducationIdentityDomainAppliesTonone,
		"student": EducationIdentityDomainAppliesTostudent,
		"teacher": EducationIdentityDomainAppliesToteacher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationIdentityDomainAppliesTo(input)
	return &out, nil
}
