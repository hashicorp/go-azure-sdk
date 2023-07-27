package usertoken

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyType string

const (
	KeyTypePrimary   KeyType = "primary"
	KeyTypeSecondary KeyType = "secondary"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypePrimary),
		string(KeyTypeSecondary),
	}
}

func parseKeyType(input string) (*KeyType, error) {
	vals := map[string]KeyType{
		"primary":   KeyTypePrimary,
		"secondary": KeyTypeSecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeyType(input)
	return &out, nil
}
