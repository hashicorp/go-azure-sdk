package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTenantAllowOrBlockListActionAction string

const (
	SecurityTenantAllowOrBlockListActionAction_Allow SecurityTenantAllowOrBlockListActionAction = "allow"
	SecurityTenantAllowOrBlockListActionAction_Block SecurityTenantAllowOrBlockListActionAction = "block"
)

func PossibleValuesForSecurityTenantAllowOrBlockListActionAction() []string {
	return []string{
		string(SecurityTenantAllowOrBlockListActionAction_Allow),
		string(SecurityTenantAllowOrBlockListActionAction_Block),
	}
}

func (s *SecurityTenantAllowOrBlockListActionAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityTenantAllowOrBlockListActionAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityTenantAllowOrBlockListActionAction(input string) (*SecurityTenantAllowOrBlockListActionAction, error) {
	vals := map[string]SecurityTenantAllowOrBlockListActionAction{
		"allow": SecurityTenantAllowOrBlockListActionAction_Allow,
		"block": SecurityTenantAllowOrBlockListActionAction_Block,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityTenantAllowOrBlockListActionAction(input)
	return &out, nil
}
