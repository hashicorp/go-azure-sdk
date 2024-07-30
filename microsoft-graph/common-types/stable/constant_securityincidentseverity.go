package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIncidentSeverity string

const (
	SecurityIncidentSeverity_High          SecurityIncidentSeverity = "high"
	SecurityIncidentSeverity_Informational SecurityIncidentSeverity = "informational"
	SecurityIncidentSeverity_Low           SecurityIncidentSeverity = "low"
	SecurityIncidentSeverity_Medium        SecurityIncidentSeverity = "medium"
	SecurityIncidentSeverity_Unknown       SecurityIncidentSeverity = "unknown"
)

func PossibleValuesForSecurityIncidentSeverity() []string {
	return []string{
		string(SecurityIncidentSeverity_High),
		string(SecurityIncidentSeverity_Informational),
		string(SecurityIncidentSeverity_Low),
		string(SecurityIncidentSeverity_Medium),
		string(SecurityIncidentSeverity_Unknown),
	}
}

func (s *SecurityIncidentSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIncidentSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIncidentSeverity(input string) (*SecurityIncidentSeverity, error) {
	vals := map[string]SecurityIncidentSeverity{
		"high":          SecurityIncidentSeverity_High,
		"informational": SecurityIncidentSeverity_Informational,
		"low":           SecurityIncidentSeverity_Low,
		"medium":        SecurityIncidentSeverity_Medium,
		"unknown":       SecurityIncidentSeverity_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIncidentSeverity(input)
	return &out, nil
}
