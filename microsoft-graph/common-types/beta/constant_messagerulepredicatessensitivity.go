package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRulePredicatesSensitivity string

const (
	MessageRulePredicatesSensitivity_Confidential MessageRulePredicatesSensitivity = "confidential"
	MessageRulePredicatesSensitivity_Normal       MessageRulePredicatesSensitivity = "normal"
	MessageRulePredicatesSensitivity_Personal     MessageRulePredicatesSensitivity = "personal"
	MessageRulePredicatesSensitivity_Private      MessageRulePredicatesSensitivity = "private"
)

func PossibleValuesForMessageRulePredicatesSensitivity() []string {
	return []string{
		string(MessageRulePredicatesSensitivity_Confidential),
		string(MessageRulePredicatesSensitivity_Normal),
		string(MessageRulePredicatesSensitivity_Personal),
		string(MessageRulePredicatesSensitivity_Private),
	}
}

func (s *MessageRulePredicatesSensitivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageRulePredicatesSensitivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageRulePredicatesSensitivity(input string) (*MessageRulePredicatesSensitivity, error) {
	vals := map[string]MessageRulePredicatesSensitivity{
		"confidential": MessageRulePredicatesSensitivity_Confidential,
		"normal":       MessageRulePredicatesSensitivity_Normal,
		"personal":     MessageRulePredicatesSensitivity_Personal,
		"private":      MessageRulePredicatesSensitivity_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageRulePredicatesSensitivity(input)
	return &out, nil
}
