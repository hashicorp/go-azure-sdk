package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationItemBodyContentType string

const (
	EducationItemBodyContentTypehtml EducationItemBodyContentType = "Html"
	EducationItemBodyContentTypetext EducationItemBodyContentType = "Text"
)

func PossibleValuesForEducationItemBodyContentType() []string {
	return []string{
		string(EducationItemBodyContentTypehtml),
		string(EducationItemBodyContentTypetext),
	}
}

func parseEducationItemBodyContentType(input string) (*EducationItemBodyContentType, error) {
	vals := map[string]EducationItemBodyContentType{
		"html": EducationItemBodyContentTypehtml,
		"text": EducationItemBodyContentTypetext,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationItemBodyContentType(input)
	return &out, nil
}
