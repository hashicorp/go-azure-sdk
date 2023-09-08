package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentInfoFormat string

const (
	ContentInfoFormatdefault ContentInfoFormat = "Default"
	ContentInfoFormatemail   ContentInfoFormat = "Email"
)

func PossibleValuesForContentInfoFormat() []string {
	return []string{
		string(ContentInfoFormatdefault),
		string(ContentInfoFormatemail),
	}
}

func parseContentInfoFormat(input string) (*ContentInfoFormat, error) {
	vals := map[string]ContentInfoFormat{
		"default": ContentInfoFormatdefault,
		"email":   ContentInfoFormatemail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentInfoFormat(input)
	return &out, nil
}
