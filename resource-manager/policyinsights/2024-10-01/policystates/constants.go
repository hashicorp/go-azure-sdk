package policystates

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyStatesResource string

const (
	PolicyStatesResourceDefault PolicyStatesResource = "default"
	PolicyStatesResourceLatest  PolicyStatesResource = "latest"
)

func PossibleValuesForPolicyStatesResource() []string {
	return []string{
		string(PolicyStatesResourceDefault),
		string(PolicyStatesResourceLatest),
	}
}

func (s *PolicyStatesResource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicyStatesResource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicyStatesResource(input string) (*PolicyStatesResource, error) {
	vals := map[string]PolicyStatesResource{
		"default": PolicyStatesResourceDefault,
		"latest":  PolicyStatesResourceLatest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyStatesResource(input)
	return &out, nil
}
