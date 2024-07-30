package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRuleActionsMarkImportance string

const (
	MessageRuleActionsMarkImportance_High   MessageRuleActionsMarkImportance = "high"
	MessageRuleActionsMarkImportance_Low    MessageRuleActionsMarkImportance = "low"
	MessageRuleActionsMarkImportance_Normal MessageRuleActionsMarkImportance = "normal"
)

func PossibleValuesForMessageRuleActionsMarkImportance() []string {
	return []string{
		string(MessageRuleActionsMarkImportance_High),
		string(MessageRuleActionsMarkImportance_Low),
		string(MessageRuleActionsMarkImportance_Normal),
	}
}

func (s *MessageRuleActionsMarkImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageRuleActionsMarkImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageRuleActionsMarkImportance(input string) (*MessageRuleActionsMarkImportance, error) {
	vals := map[string]MessageRuleActionsMarkImportance{
		"high":   MessageRuleActionsMarkImportance_High,
		"low":    MessageRuleActionsMarkImportance_Low,
		"normal": MessageRuleActionsMarkImportance_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageRuleActionsMarkImportance(input)
	return &out, nil
}
