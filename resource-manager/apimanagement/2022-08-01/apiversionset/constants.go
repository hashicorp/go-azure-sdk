package apiversionset

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VersioningScheme string

const (
	VersioningSchemeHeader  VersioningScheme = "Header"
	VersioningSchemeQuery   VersioningScheme = "Query"
	VersioningSchemeSegment VersioningScheme = "Segment"
)

func PossibleValuesForVersioningScheme() []string {
	return []string{
		string(VersioningSchemeHeader),
		string(VersioningSchemeQuery),
		string(VersioningSchemeSegment),
	}
}

func parseVersioningScheme(input string) (*VersioningScheme, error) {
	vals := map[string]VersioningScheme{
		"header":  VersioningSchemeHeader,
		"query":   VersioningSchemeQuery,
		"segment": VersioningSchemeSegment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VersioningScheme(input)
	return &out, nil
}
