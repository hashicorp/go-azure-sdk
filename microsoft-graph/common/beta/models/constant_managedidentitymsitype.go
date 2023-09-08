package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIdentityMsiType string

const (
	ManagedIdentityMsiTypenone           ManagedIdentityMsiType = "None"
	ManagedIdentityMsiTypesystemAssigned ManagedIdentityMsiType = "SystemAssigned"
	ManagedIdentityMsiTypeuserAssigned   ManagedIdentityMsiType = "UserAssigned"
)

func PossibleValuesForManagedIdentityMsiType() []string {
	return []string{
		string(ManagedIdentityMsiTypenone),
		string(ManagedIdentityMsiTypesystemAssigned),
		string(ManagedIdentityMsiTypeuserAssigned),
	}
}

func parseManagedIdentityMsiType(input string) (*ManagedIdentityMsiType, error) {
	vals := map[string]ManagedIdentityMsiType{
		"none":           ManagedIdentityMsiTypenone,
		"systemassigned": ManagedIdentityMsiTypesystemAssigned,
		"userassigned":   ManagedIdentityMsiTypeuserAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedIdentityMsiType(input)
	return &out, nil
}
