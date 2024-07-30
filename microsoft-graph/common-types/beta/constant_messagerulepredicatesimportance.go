package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRulePredicatesImportance string

const (
	MessageRulePredicatesImportance_High   MessageRulePredicatesImportance = "high"
	MessageRulePredicatesImportance_Low    MessageRulePredicatesImportance = "low"
	MessageRulePredicatesImportance_Normal MessageRulePredicatesImportance = "normal"
)

func PossibleValuesForMessageRulePredicatesImportance() []string {
	return []string{
		string(MessageRulePredicatesImportance_High),
		string(MessageRulePredicatesImportance_Low),
		string(MessageRulePredicatesImportance_Normal),
	}
}

func (s *MessageRulePredicatesImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageRulePredicatesImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageRulePredicatesImportance(input string) (*MessageRulePredicatesImportance, error) {
	vals := map[string]MessageRulePredicatesImportance{
		"high":   MessageRulePredicatesImportance_High,
		"low":    MessageRulePredicatesImportance_Low,
		"normal": MessageRulePredicatesImportance_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageRulePredicatesImportance(input)
	return &out, nil
}
