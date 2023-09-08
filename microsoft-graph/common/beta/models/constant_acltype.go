package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AclType string

const (
	AclTypeeveryone             AclType = "Everyone"
	AclTypeeveryoneExceptGuests AclType = "EveryoneExceptGuests"
	AclTypeexternalGroup        AclType = "ExternalGroup"
	AclTypegroup                AclType = "Group"
	AclTypeuser                 AclType = "User"
)

func PossibleValuesForAclType() []string {
	return []string{
		string(AclTypeeveryone),
		string(AclTypeeveryoneExceptGuests),
		string(AclTypeexternalGroup),
		string(AclTypegroup),
		string(AclTypeuser),
	}
}

func parseAclType(input string) (*AclType, error) {
	vals := map[string]AclType{
		"everyone":             AclTypeeveryone,
		"everyoneexceptguests": AclTypeeveryoneExceptGuests,
		"externalgroup":        AclTypeexternalGroup,
		"group":                AclTypegroup,
		"user":                 AclTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AclType(input)
	return &out, nil
}
