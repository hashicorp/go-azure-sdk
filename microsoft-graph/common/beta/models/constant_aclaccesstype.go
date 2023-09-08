package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AclAccessType string

const (
	AclAccessTypedeny  AclAccessType = "Deny"
	AclAccessTypegrant AclAccessType = "Grant"
)

func PossibleValuesForAclAccessType() []string {
	return []string{
		string(AclAccessTypedeny),
		string(AclAccessTypegrant),
	}
}

func parseAclAccessType(input string) (*AclAccessType, error) {
	vals := map[string]AclAccessType{
		"deny":  AclAccessTypedeny,
		"grant": AclAccessTypegrant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AclAccessType(input)
	return &out, nil
}
