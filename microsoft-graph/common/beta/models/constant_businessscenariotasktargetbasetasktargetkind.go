package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioTaskTargetBaseTaskTargetKind string

const (
	BusinessScenarioTaskTargetBaseTaskTargetKindgroup BusinessScenarioTaskTargetBaseTaskTargetKind = "Group"
)

func PossibleValuesForBusinessScenarioTaskTargetBaseTaskTargetKind() []string {
	return []string{
		string(BusinessScenarioTaskTargetBaseTaskTargetKindgroup),
	}
}

func parseBusinessScenarioTaskTargetBaseTaskTargetKind(input string) (*BusinessScenarioTaskTargetBaseTaskTargetKind, error) {
	vals := map[string]BusinessScenarioTaskTargetBaseTaskTargetKind{
		"group": BusinessScenarioTaskTargetBaseTaskTargetKindgroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BusinessScenarioTaskTargetBaseTaskTargetKind(input)
	return &out, nil
}
