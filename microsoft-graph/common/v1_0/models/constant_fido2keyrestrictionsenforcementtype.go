package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2KeyRestrictionsEnforcementType string

const (
	Fido2KeyRestrictionsEnforcementTypeallow Fido2KeyRestrictionsEnforcementType = "Allow"
	Fido2KeyRestrictionsEnforcementTypeblock Fido2KeyRestrictionsEnforcementType = "Block"
)

func PossibleValuesForFido2KeyRestrictionsEnforcementType() []string {
	return []string{
		string(Fido2KeyRestrictionsEnforcementTypeallow),
		string(Fido2KeyRestrictionsEnforcementTypeblock),
	}
}

func parseFido2KeyRestrictionsEnforcementType(input string) (*Fido2KeyRestrictionsEnforcementType, error) {
	vals := map[string]Fido2KeyRestrictionsEnforcementType{
		"allow": Fido2KeyRestrictionsEnforcementTypeallow,
		"block": Fido2KeyRestrictionsEnforcementTypeblock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Fido2KeyRestrictionsEnforcementType(input)
	return &out, nil
}
