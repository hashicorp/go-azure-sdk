package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMappingFlowBehavior string

const (
	AttributeMappingFlowBehaviorFlowAlways      AttributeMappingFlowBehavior = "FlowAlways"
	AttributeMappingFlowBehaviorFlowWhenChanged AttributeMappingFlowBehavior = "FlowWhenChanged"
)

func PossibleValuesForAttributeMappingFlowBehavior() []string {
	return []string{
		string(AttributeMappingFlowBehaviorFlowAlways),
		string(AttributeMappingFlowBehaviorFlowWhenChanged),
	}
}

func parseAttributeMappingFlowBehavior(input string) (*AttributeMappingFlowBehavior, error) {
	vals := map[string]AttributeMappingFlowBehavior{
		"flowalways":      AttributeMappingFlowBehaviorFlowAlways,
		"flowwhenchanged": AttributeMappingFlowBehaviorFlowWhenChanged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeMappingFlowBehavior(input)
	return &out, nil
}
