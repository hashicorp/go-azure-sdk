package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SystemCredentialPreferencesState string

const (
	SystemCredentialPreferencesStatedefault  SystemCredentialPreferencesState = "Default"
	SystemCredentialPreferencesStatedisabled SystemCredentialPreferencesState = "Disabled"
	SystemCredentialPreferencesStateenabled  SystemCredentialPreferencesState = "Enabled"
)

func PossibleValuesForSystemCredentialPreferencesState() []string {
	return []string{
		string(SystemCredentialPreferencesStatedefault),
		string(SystemCredentialPreferencesStatedisabled),
		string(SystemCredentialPreferencesStateenabled),
	}
}

func parseSystemCredentialPreferencesState(input string) (*SystemCredentialPreferencesState, error) {
	vals := map[string]SystemCredentialPreferencesState{
		"default":  SystemCredentialPreferencesStatedefault,
		"disabled": SystemCredentialPreferencesStatedisabled,
		"enabled":  SystemCredentialPreferencesStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SystemCredentialPreferencesState(input)
	return &out, nil
}
