package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleOwnerTypeEnrollmentTypeOwnerType string

const (
	AppleOwnerTypeEnrollmentTypeOwnerTypecompany  AppleOwnerTypeEnrollmentTypeOwnerType = "Company"
	AppleOwnerTypeEnrollmentTypeOwnerTypepersonal AppleOwnerTypeEnrollmentTypeOwnerType = "Personal"
	AppleOwnerTypeEnrollmentTypeOwnerTypeunknown  AppleOwnerTypeEnrollmentTypeOwnerType = "Unknown"
)

func PossibleValuesForAppleOwnerTypeEnrollmentTypeOwnerType() []string {
	return []string{
		string(AppleOwnerTypeEnrollmentTypeOwnerTypecompany),
		string(AppleOwnerTypeEnrollmentTypeOwnerTypepersonal),
		string(AppleOwnerTypeEnrollmentTypeOwnerTypeunknown),
	}
}

func parseAppleOwnerTypeEnrollmentTypeOwnerType(input string) (*AppleOwnerTypeEnrollmentTypeOwnerType, error) {
	vals := map[string]AppleOwnerTypeEnrollmentTypeOwnerType{
		"company":  AppleOwnerTypeEnrollmentTypeOwnerTypecompany,
		"personal": AppleOwnerTypeEnrollmentTypeOwnerTypepersonal,
		"unknown":  AppleOwnerTypeEnrollmentTypeOwnerTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleOwnerTypeEnrollmentTypeOwnerType(input)
	return &out, nil
}
