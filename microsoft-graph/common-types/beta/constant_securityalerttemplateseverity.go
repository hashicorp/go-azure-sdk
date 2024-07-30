package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertTemplateSeverity string

const (
	SecurityAlertTemplateSeverity_High          SecurityAlertTemplateSeverity = "high"
	SecurityAlertTemplateSeverity_Informational SecurityAlertTemplateSeverity = "informational"
	SecurityAlertTemplateSeverity_Low           SecurityAlertTemplateSeverity = "low"
	SecurityAlertTemplateSeverity_Medium        SecurityAlertTemplateSeverity = "medium"
	SecurityAlertTemplateSeverity_Unknown       SecurityAlertTemplateSeverity = "unknown"
)

func PossibleValuesForSecurityAlertTemplateSeverity() []string {
	return []string{
		string(SecurityAlertTemplateSeverity_High),
		string(SecurityAlertTemplateSeverity_Informational),
		string(SecurityAlertTemplateSeverity_Low),
		string(SecurityAlertTemplateSeverity_Medium),
		string(SecurityAlertTemplateSeverity_Unknown),
	}
}

func (s *SecurityAlertTemplateSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAlertTemplateSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAlertTemplateSeverity(input string) (*SecurityAlertTemplateSeverity, error) {
	vals := map[string]SecurityAlertTemplateSeverity{
		"high":          SecurityAlertTemplateSeverity_High,
		"informational": SecurityAlertTemplateSeverity_Informational,
		"low":           SecurityAlertTemplateSeverity_Low,
		"medium":        SecurityAlertTemplateSeverity_Medium,
		"unknown":       SecurityAlertTemplateSeverity_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertTemplateSeverity(input)
	return &out, nil
}
