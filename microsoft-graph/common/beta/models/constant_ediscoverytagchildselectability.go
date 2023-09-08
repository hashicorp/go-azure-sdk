package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryTagChildSelectability string

const (
	EdiscoveryTagChildSelectabilityMany EdiscoveryTagChildSelectability = "Many"
	EdiscoveryTagChildSelectabilityOne  EdiscoveryTagChildSelectability = "One"
)

func PossibleValuesForEdiscoveryTagChildSelectability() []string {
	return []string{
		string(EdiscoveryTagChildSelectabilityMany),
		string(EdiscoveryTagChildSelectabilityOne),
	}
}

func parseEdiscoveryTagChildSelectability(input string) (*EdiscoveryTagChildSelectability, error) {
	vals := map[string]EdiscoveryTagChildSelectability{
		"many": EdiscoveryTagChildSelectabilityMany,
		"one":  EdiscoveryTagChildSelectabilityOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryTagChildSelectability(input)
	return &out, nil
}
