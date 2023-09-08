package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTenantAllowOrBlockListActionAction string

const (
	SecurityTenantAllowOrBlockListActionActionallow SecurityTenantAllowOrBlockListActionAction = "Allow"
	SecurityTenantAllowOrBlockListActionActionblock SecurityTenantAllowOrBlockListActionAction = "Block"
)

func PossibleValuesForSecurityTenantAllowOrBlockListActionAction() []string {
	return []string{
		string(SecurityTenantAllowOrBlockListActionActionallow),
		string(SecurityTenantAllowOrBlockListActionActionblock),
	}
}

func parseSecurityTenantAllowOrBlockListActionAction(input string) (*SecurityTenantAllowOrBlockListActionAction, error) {
	vals := map[string]SecurityTenantAllowOrBlockListActionAction{
		"allow": SecurityTenantAllowOrBlockListActionActionallow,
		"block": SecurityTenantAllowOrBlockListActionActionblock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityTenantAllowOrBlockListActionAction(input)
	return &out, nil
}
