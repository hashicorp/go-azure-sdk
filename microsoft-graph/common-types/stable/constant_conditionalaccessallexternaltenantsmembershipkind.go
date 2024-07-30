package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessAllExternalTenantsMembershipKind string

const (
	ConditionalAccessAllExternalTenantsMembershipKind_All        ConditionalAccessAllExternalTenantsMembershipKind = "all"
	ConditionalAccessAllExternalTenantsMembershipKind_Enumerated ConditionalAccessAllExternalTenantsMembershipKind = "enumerated"
)

func PossibleValuesForConditionalAccessAllExternalTenantsMembershipKind() []string {
	return []string{
		string(ConditionalAccessAllExternalTenantsMembershipKind_All),
		string(ConditionalAccessAllExternalTenantsMembershipKind_Enumerated),
	}
}

func (s *ConditionalAccessAllExternalTenantsMembershipKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessAllExternalTenantsMembershipKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessAllExternalTenantsMembershipKind(input string) (*ConditionalAccessAllExternalTenantsMembershipKind, error) {
	vals := map[string]ConditionalAccessAllExternalTenantsMembershipKind{
		"all":        ConditionalAccessAllExternalTenantsMembershipKind_All,
		"enumerated": ConditionalAccessAllExternalTenantsMembershipKind_Enumerated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessAllExternalTenantsMembershipKind(input)
	return &out, nil
}
