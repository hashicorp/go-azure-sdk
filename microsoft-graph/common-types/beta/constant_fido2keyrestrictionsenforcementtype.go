package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2KeyRestrictionsEnforcementType string

const (
	Fido2KeyRestrictionsEnforcementType_Allow Fido2KeyRestrictionsEnforcementType = "allow"
	Fido2KeyRestrictionsEnforcementType_Block Fido2KeyRestrictionsEnforcementType = "block"
)

func PossibleValuesForFido2KeyRestrictionsEnforcementType() []string {
	return []string{
		string(Fido2KeyRestrictionsEnforcementType_Allow),
		string(Fido2KeyRestrictionsEnforcementType_Block),
	}
}

func (s *Fido2KeyRestrictionsEnforcementType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFido2KeyRestrictionsEnforcementType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFido2KeyRestrictionsEnforcementType(input string) (*Fido2KeyRestrictionsEnforcementType, error) {
	vals := map[string]Fido2KeyRestrictionsEnforcementType{
		"allow": Fido2KeyRestrictionsEnforcementType_Allow,
		"block": Fido2KeyRestrictionsEnforcementType_Block,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Fido2KeyRestrictionsEnforcementType(input)
	return &out, nil
}
