package availableworkloadprofiles

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Applicability string

const (
	ApplicabilityCustom          Applicability = "Custom"
	ApplicabilityLocationDefault Applicability = "LocationDefault"
)

func PossibleValuesForApplicability() []string {
	return []string{
		string(ApplicabilityCustom),
		string(ApplicabilityLocationDefault),
	}
}

func parseApplicability(input string) (*Applicability, error) {
	vals := map[string]Applicability{
		"custom":          ApplicabilityCustom,
		"locationdefault": ApplicabilityLocationDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Applicability(input)
	return &out, nil
}
