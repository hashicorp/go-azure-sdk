package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsIdentityType string

const (
	ExternalConnectorsIdentityTypeexternalGroup ExternalConnectorsIdentityType = "ExternalGroup"
	ExternalConnectorsIdentityTypegroup         ExternalConnectorsIdentityType = "Group"
	ExternalConnectorsIdentityTypeuser          ExternalConnectorsIdentityType = "User"
)

func PossibleValuesForExternalConnectorsIdentityType() []string {
	return []string{
		string(ExternalConnectorsIdentityTypeexternalGroup),
		string(ExternalConnectorsIdentityTypegroup),
		string(ExternalConnectorsIdentityTypeuser),
	}
}

func parseExternalConnectorsIdentityType(input string) (*ExternalConnectorsIdentityType, error) {
	vals := map[string]ExternalConnectorsIdentityType{
		"externalgroup": ExternalConnectorsIdentityTypeexternalGroup,
		"group":         ExternalConnectorsIdentityTypegroup,
		"user":          ExternalConnectorsIdentityTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsIdentityType(input)
	return &out, nil
}
