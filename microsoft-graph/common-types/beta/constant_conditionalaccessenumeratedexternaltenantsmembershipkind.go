package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessEnumeratedExternalTenantsMembershipKind string

const (
	ConditionalAccessEnumeratedExternalTenantsMembershipKind_All        ConditionalAccessEnumeratedExternalTenantsMembershipKind = "all"
	ConditionalAccessEnumeratedExternalTenantsMembershipKind_Enumerated ConditionalAccessEnumeratedExternalTenantsMembershipKind = "enumerated"
)

func PossibleValuesForConditionalAccessEnumeratedExternalTenantsMembershipKind() []string {
	return []string{
		string(ConditionalAccessEnumeratedExternalTenantsMembershipKind_All),
		string(ConditionalAccessEnumeratedExternalTenantsMembershipKind_Enumerated),
	}
}

func (s *ConditionalAccessEnumeratedExternalTenantsMembershipKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessEnumeratedExternalTenantsMembershipKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessEnumeratedExternalTenantsMembershipKind(input string) (*ConditionalAccessEnumeratedExternalTenantsMembershipKind, error) {
	vals := map[string]ConditionalAccessEnumeratedExternalTenantsMembershipKind{
		"all":        ConditionalAccessEnumeratedExternalTenantsMembershipKind_All,
		"enumerated": ConditionalAccessEnumeratedExternalTenantsMembershipKind_Enumerated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessEnumeratedExternalTenantsMembershipKind(input)
	return &out, nil
}
