package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCvssSummarySeverity string

const (
	SecurityCvssSummarySeverity_Critical SecurityCvssSummarySeverity = "critical"
	SecurityCvssSummarySeverity_High     SecurityCvssSummarySeverity = "high"
	SecurityCvssSummarySeverity_Low      SecurityCvssSummarySeverity = "low"
	SecurityCvssSummarySeverity_Medium   SecurityCvssSummarySeverity = "medium"
	SecurityCvssSummarySeverity_None     SecurityCvssSummarySeverity = "none"
)

func PossibleValuesForSecurityCvssSummarySeverity() []string {
	return []string{
		string(SecurityCvssSummarySeverity_Critical),
		string(SecurityCvssSummarySeverity_High),
		string(SecurityCvssSummarySeverity_Low),
		string(SecurityCvssSummarySeverity_Medium),
		string(SecurityCvssSummarySeverity_None),
	}
}

func (s *SecurityCvssSummarySeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityCvssSummarySeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityCvssSummarySeverity(input string) (*SecurityCvssSummarySeverity, error) {
	vals := map[string]SecurityCvssSummarySeverity{
		"critical": SecurityCvssSummarySeverity_Critical,
		"high":     SecurityCvssSummarySeverity_High,
		"low":      SecurityCvssSummarySeverity_Low,
		"medium":   SecurityCvssSummarySeverity_Medium,
		"none":     SecurityCvssSummarySeverity_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCvssSummarySeverity(input)
	return &out, nil
}
