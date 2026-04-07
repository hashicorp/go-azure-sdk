package networksecurityperimeterconfigurations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NspAccessRuleDirection string

const (
	NspAccessRuleDirectionInbound  NspAccessRuleDirection = "Inbound"
	NspAccessRuleDirectionOutbound NspAccessRuleDirection = "Outbound"
)

func PossibleValuesForNspAccessRuleDirection() []string {
	return []string{
		string(NspAccessRuleDirectionInbound),
		string(NspAccessRuleDirectionOutbound),
	}
}

func (s *NspAccessRuleDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNspAccessRuleDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNspAccessRuleDirection(input string) (*NspAccessRuleDirection, error) {
	vals := map[string]NspAccessRuleDirection{
		"inbound":  NspAccessRuleDirectionInbound,
		"outbound": NspAccessRuleDirectionOutbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NspAccessRuleDirection(input)
	return &out, nil
}
