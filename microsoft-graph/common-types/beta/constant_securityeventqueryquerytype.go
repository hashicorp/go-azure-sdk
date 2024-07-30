package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventQueryQueryType string

const (
	SecurityEventQueryQueryType_Files    SecurityEventQueryQueryType = "files"
	SecurityEventQueryQueryType_Messages SecurityEventQueryQueryType = "messages"
)

func PossibleValuesForSecurityEventQueryQueryType() []string {
	return []string{
		string(SecurityEventQueryQueryType_Files),
		string(SecurityEventQueryQueryType_Messages),
	}
}

func (s *SecurityEventQueryQueryType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEventQueryQueryType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEventQueryQueryType(input string) (*SecurityEventQueryQueryType, error) {
	vals := map[string]SecurityEventQueryQueryType{
		"files":    SecurityEventQueryQueryType_Files,
		"messages": SecurityEventQueryQueryType_Messages,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEventQueryQueryType(input)
	return &out, nil
}
