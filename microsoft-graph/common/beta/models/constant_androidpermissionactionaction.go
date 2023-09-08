package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidPermissionActionAction string

const (
	AndroidPermissionActionActionautoDeny  AndroidPermissionActionAction = "AutoDeny"
	AndroidPermissionActionActionautoGrant AndroidPermissionActionAction = "AutoGrant"
	AndroidPermissionActionActionprompt    AndroidPermissionActionAction = "Prompt"
)

func PossibleValuesForAndroidPermissionActionAction() []string {
	return []string{
		string(AndroidPermissionActionActionautoDeny),
		string(AndroidPermissionActionActionautoGrant),
		string(AndroidPermissionActionActionprompt),
	}
}

func parseAndroidPermissionActionAction(input string) (*AndroidPermissionActionAction, error) {
	vals := map[string]AndroidPermissionActionAction{
		"autodeny":  AndroidPermissionActionActionautoDeny,
		"autogrant": AndroidPermissionActionActionautoGrant,
		"prompt":    AndroidPermissionActionActionprompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidPermissionActionAction(input)
	return &out, nil
}
