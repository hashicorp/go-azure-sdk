package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioTaskSpecifiedCompletionRequirements string

const (
	BusinessScenarioTaskSpecifiedCompletionRequirements_ChecklistCompletion BusinessScenarioTaskSpecifiedCompletionRequirements = "checklistCompletion"
	BusinessScenarioTaskSpecifiedCompletionRequirements_None                BusinessScenarioTaskSpecifiedCompletionRequirements = "none"
)

func PossibleValuesForBusinessScenarioTaskSpecifiedCompletionRequirements() []string {
	return []string{
		string(BusinessScenarioTaskSpecifiedCompletionRequirements_ChecklistCompletion),
		string(BusinessScenarioTaskSpecifiedCompletionRequirements_None),
	}
}

func (s *BusinessScenarioTaskSpecifiedCompletionRequirements) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBusinessScenarioTaskSpecifiedCompletionRequirements(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBusinessScenarioTaskSpecifiedCompletionRequirements(input string) (*BusinessScenarioTaskSpecifiedCompletionRequirements, error) {
	vals := map[string]BusinessScenarioTaskSpecifiedCompletionRequirements{
		"checklistcompletion": BusinessScenarioTaskSpecifiedCompletionRequirements_ChecklistCompletion,
		"none":                BusinessScenarioTaskSpecifiedCompletionRequirements_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BusinessScenarioTaskSpecifiedCompletionRequirements(input)
	return &out, nil
}
