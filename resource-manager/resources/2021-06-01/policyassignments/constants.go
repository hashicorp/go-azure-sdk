package policyassignments

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnforcementMode string

const (
	EnforcementModeDefault      EnforcementMode = "Default"
	EnforcementModeDoNotEnforce EnforcementMode = "DoNotEnforce"
)

func PossibleValuesForEnforcementMode() []string {
	return []string{
		string(EnforcementModeDefault),
		string(EnforcementModeDoNotEnforce),
	}
}

func parseEnforcementMode(input string) (*EnforcementMode, error) {
	vals := map[string]EnforcementMode{
		"default":      EnforcementModeDefault,
		"donotenforce": EnforcementModeDoNotEnforce,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnforcementMode(input)
	return &out, nil
}
