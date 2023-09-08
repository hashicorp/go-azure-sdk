package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioGroupTargetTaskTargetKind string

const (
	BusinessScenarioGroupTargetTaskTargetKindgroup BusinessScenarioGroupTargetTaskTargetKind = "Group"
)

func PossibleValuesForBusinessScenarioGroupTargetTaskTargetKind() []string {
	return []string{
		string(BusinessScenarioGroupTargetTaskTargetKindgroup),
	}
}

func parseBusinessScenarioGroupTargetTaskTargetKind(input string) (*BusinessScenarioGroupTargetTaskTargetKind, error) {
	vals := map[string]BusinessScenarioGroupTargetTaskTargetKind{
		"group": BusinessScenarioGroupTargetTaskTargetKindgroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BusinessScenarioGroupTargetTaskTargetKind(input)
	return &out, nil
}
