package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMappingFlowBehavior string

const (
	AttributeMappingFlowBehavior_FlowAlways      AttributeMappingFlowBehavior = "FlowAlways"
	AttributeMappingFlowBehavior_FlowWhenChanged AttributeMappingFlowBehavior = "FlowWhenChanged"
)

func PossibleValuesForAttributeMappingFlowBehavior() []string {
	return []string{
		string(AttributeMappingFlowBehavior_FlowAlways),
		string(AttributeMappingFlowBehavior_FlowWhenChanged),
	}
}

func (s *AttributeMappingFlowBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeMappingFlowBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeMappingFlowBehavior(input string) (*AttributeMappingFlowBehavior, error) {
	vals := map[string]AttributeMappingFlowBehavior{
		"flowalways":      AttributeMappingFlowBehavior_FlowAlways,
		"flowwhenchanged": AttributeMappingFlowBehavior_FlowWhenChanged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeMappingFlowBehavior(input)
	return &out, nil
}
