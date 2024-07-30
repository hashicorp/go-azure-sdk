package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PostImportance string

const (
	PostImportance_High   PostImportance = "high"
	PostImportance_Low    PostImportance = "low"
	PostImportance_Normal PostImportance = "normal"
)

func PossibleValuesForPostImportance() []string {
	return []string{
		string(PostImportance_High),
		string(PostImportance_Low),
		string(PostImportance_Normal),
	}
}

func (s *PostImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePostImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePostImportance(input string) (*PostImportance, error) {
	vals := map[string]PostImportance{
		"high":   PostImportance_High,
		"low":    PostImportance_Low,
		"normal": PostImportance_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PostImportance(input)
	return &out, nil
}
