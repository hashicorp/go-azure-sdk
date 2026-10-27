package hcpoperatoridentityrolesets

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperatorIdentityRequired string

const (
	OperatorIdentityRequiredAlways       OperatorIdentityRequired = "Always"
	OperatorIdentityRequiredOnEnablement OperatorIdentityRequired = "OnEnablement"
)

func PossibleValuesForOperatorIdentityRequired() []string {
	return []string{
		string(OperatorIdentityRequiredAlways),
		string(OperatorIdentityRequiredOnEnablement),
	}
}

func (s *OperatorIdentityRequired) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperatorIdentityRequired(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperatorIdentityRequired(input string) (*OperatorIdentityRequired, error) {
	vals := map[string]OperatorIdentityRequired{
		"always":       OperatorIdentityRequiredAlways,
		"onenablement": OperatorIdentityRequiredOnEnablement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperatorIdentityRequired(input)
	return &out, nil
}
