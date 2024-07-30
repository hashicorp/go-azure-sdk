package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageImportance string

const (
	MessageImportance_High   MessageImportance = "high"
	MessageImportance_Low    MessageImportance = "low"
	MessageImportance_Normal MessageImportance = "normal"
)

func PossibleValuesForMessageImportance() []string {
	return []string{
		string(MessageImportance_High),
		string(MessageImportance_Low),
		string(MessageImportance_Normal),
	}
}

func (s *MessageImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageImportance(input string) (*MessageImportance, error) {
	vals := map[string]MessageImportance{
		"high":   MessageImportance_High,
		"low":    MessageImportance_Low,
		"normal": MessageImportance_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageImportance(input)
	return &out, nil
}
