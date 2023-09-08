package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureTargetTargetType string

const (
	FeatureTargetTargetTypeadministrativeUnit FeatureTargetTargetType = "AdministrativeUnit"
	FeatureTargetTargetTypegroup              FeatureTargetTargetType = "Group"
	FeatureTargetTargetTyperole               FeatureTargetTargetType = "Role"
)

func PossibleValuesForFeatureTargetTargetType() []string {
	return []string{
		string(FeatureTargetTargetTypeadministrativeUnit),
		string(FeatureTargetTargetTypegroup),
		string(FeatureTargetTargetTyperole),
	}
}

func parseFeatureTargetTargetType(input string) (*FeatureTargetTargetType, error) {
	vals := map[string]FeatureTargetTargetType{
		"administrativeunit": FeatureTargetTargetTypeadministrativeUnit,
		"group":              FeatureTargetTargetTypegroup,
		"role":               FeatureTargetTargetTyperole,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureTargetTargetType(input)
	return &out, nil
}
