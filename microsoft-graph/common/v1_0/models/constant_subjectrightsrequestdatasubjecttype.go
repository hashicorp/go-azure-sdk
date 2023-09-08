package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestDataSubjectType string

const (
	SubjectRightsRequestDataSubjectTypecurrentEmployee     SubjectRightsRequestDataSubjectType = "CurrentEmployee"
	SubjectRightsRequestDataSubjectTypecustomer            SubjectRightsRequestDataSubjectType = "Customer"
	SubjectRightsRequestDataSubjectTypefaculty             SubjectRightsRequestDataSubjectType = "Faculty"
	SubjectRightsRequestDataSubjectTypeformerEmployee      SubjectRightsRequestDataSubjectType = "FormerEmployee"
	SubjectRightsRequestDataSubjectTypeother               SubjectRightsRequestDataSubjectType = "Other"
	SubjectRightsRequestDataSubjectTypeprospectiveEmployee SubjectRightsRequestDataSubjectType = "ProspectiveEmployee"
	SubjectRightsRequestDataSubjectTypestudent             SubjectRightsRequestDataSubjectType = "Student"
	SubjectRightsRequestDataSubjectTypeteacher             SubjectRightsRequestDataSubjectType = "Teacher"
)

func PossibleValuesForSubjectRightsRequestDataSubjectType() []string {
	return []string{
		string(SubjectRightsRequestDataSubjectTypecurrentEmployee),
		string(SubjectRightsRequestDataSubjectTypecustomer),
		string(SubjectRightsRequestDataSubjectTypefaculty),
		string(SubjectRightsRequestDataSubjectTypeformerEmployee),
		string(SubjectRightsRequestDataSubjectTypeother),
		string(SubjectRightsRequestDataSubjectTypeprospectiveEmployee),
		string(SubjectRightsRequestDataSubjectTypestudent),
		string(SubjectRightsRequestDataSubjectTypeteacher),
	}
}

func parseSubjectRightsRequestDataSubjectType(input string) (*SubjectRightsRequestDataSubjectType, error) {
	vals := map[string]SubjectRightsRequestDataSubjectType{
		"currentemployee":     SubjectRightsRequestDataSubjectTypecurrentEmployee,
		"customer":            SubjectRightsRequestDataSubjectTypecustomer,
		"faculty":             SubjectRightsRequestDataSubjectTypefaculty,
		"formeremployee":      SubjectRightsRequestDataSubjectTypeformerEmployee,
		"other":               SubjectRightsRequestDataSubjectTypeother,
		"prospectiveemployee": SubjectRightsRequestDataSubjectTypeprospectiveEmployee,
		"student":             SubjectRightsRequestDataSubjectTypestudent,
		"teacher":             SubjectRightsRequestDataSubjectTypeteacher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestDataSubjectType(input)
	return &out, nil
}
