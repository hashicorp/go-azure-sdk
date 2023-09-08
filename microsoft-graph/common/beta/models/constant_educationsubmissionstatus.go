package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSubmissionStatus string

const (
	EducationSubmissionStatusreassigned EducationSubmissionStatus = "Reassigned"
	EducationSubmissionStatusreleased   EducationSubmissionStatus = "Released"
	EducationSubmissionStatusreturned   EducationSubmissionStatus = "Returned"
	EducationSubmissionStatussubmitted  EducationSubmissionStatus = "Submitted"
	EducationSubmissionStatusworking    EducationSubmissionStatus = "Working"
)

func PossibleValuesForEducationSubmissionStatus() []string {
	return []string{
		string(EducationSubmissionStatusreassigned),
		string(EducationSubmissionStatusreleased),
		string(EducationSubmissionStatusreturned),
		string(EducationSubmissionStatussubmitted),
		string(EducationSubmissionStatusworking),
	}
}

func parseEducationSubmissionStatus(input string) (*EducationSubmissionStatus, error) {
	vals := map[string]EducationSubmissionStatus{
		"reassigned": EducationSubmissionStatusreassigned,
		"released":   EducationSubmissionStatusreleased,
		"returned":   EducationSubmissionStatusreturned,
		"submitted":  EducationSubmissionStatussubmitted,
		"working":    EducationSubmissionStatusworking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationSubmissionStatus(input)
	return &out, nil
}
