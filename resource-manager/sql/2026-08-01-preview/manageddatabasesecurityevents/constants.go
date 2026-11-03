package manageddatabasesecurityevents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventType string

const (
	SecurityEventTypeSqlInjectionExploit       SecurityEventType = "SqlInjectionExploit"
	SecurityEventTypeSqlInjectionVulnerability SecurityEventType = "SqlInjectionVulnerability"
	SecurityEventTypeUndefined                 SecurityEventType = "Undefined"
)

func PossibleValuesForSecurityEventType() []string {
	return []string{
		string(SecurityEventTypeSqlInjectionExploit),
		string(SecurityEventTypeSqlInjectionVulnerability),
		string(SecurityEventTypeUndefined),
	}
}

func (s *SecurityEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEventType(input string) (*SecurityEventType, error) {
	vals := map[string]SecurityEventType{
		"sqlinjectionexploit":       SecurityEventTypeSqlInjectionExploit,
		"sqlinjectionvulnerability": SecurityEventTypeSqlInjectionVulnerability,
		"undefined":                 SecurityEventTypeUndefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEventType(input)
	return &out, nil
}
