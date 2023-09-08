package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AclIdentitySource string

const (
	AclIdentitySourceazureActiveDirectory AclIdentitySource = "AzureActiveDirectory"
	AclIdentitySourceexternal             AclIdentitySource = "External"
)

func PossibleValuesForAclIdentitySource() []string {
	return []string{
		string(AclIdentitySourceazureActiveDirectory),
		string(AclIdentitySourceexternal),
	}
}

func parseAclIdentitySource(input string) (*AclIdentitySource, error) {
	vals := map[string]AclIdentitySource{
		"azureactivedirectory": AclIdentitySourceazureActiveDirectory,
		"external":             AclIdentitySourceexternal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AclIdentitySource(input)
	return &out, nil
}
