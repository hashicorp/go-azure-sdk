package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsAclType string

const (
	ExternalConnectorsAclTypeeveryone             ExternalConnectorsAclType = "Everyone"
	ExternalConnectorsAclTypeeveryoneExceptGuests ExternalConnectorsAclType = "EveryoneExceptGuests"
	ExternalConnectorsAclTypeexternalGroup        ExternalConnectorsAclType = "ExternalGroup"
	ExternalConnectorsAclTypegroup                ExternalConnectorsAclType = "Group"
	ExternalConnectorsAclTypeuser                 ExternalConnectorsAclType = "User"
)

func PossibleValuesForExternalConnectorsAclType() []string {
	return []string{
		string(ExternalConnectorsAclTypeeveryone),
		string(ExternalConnectorsAclTypeeveryoneExceptGuests),
		string(ExternalConnectorsAclTypeexternalGroup),
		string(ExternalConnectorsAclTypegroup),
		string(ExternalConnectorsAclTypeuser),
	}
}

func parseExternalConnectorsAclType(input string) (*ExternalConnectorsAclType, error) {
	vals := map[string]ExternalConnectorsAclType{
		"everyone":             ExternalConnectorsAclTypeeveryone,
		"everyoneexceptguests": ExternalConnectorsAclTypeeveryoneExceptGuests,
		"externalgroup":        ExternalConnectorsAclTypeexternalGroup,
		"group":                ExternalConnectorsAclTypegroup,
		"user":                 ExternalConnectorsAclTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsAclType(input)
	return &out, nil
}
