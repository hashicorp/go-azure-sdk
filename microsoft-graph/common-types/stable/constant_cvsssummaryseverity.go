package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CvssSummarySeverity string

const (
	CvssSummarySeverity_Critical CvssSummarySeverity = "critical"
	CvssSummarySeverity_High     CvssSummarySeverity = "high"
	CvssSummarySeverity_Low      CvssSummarySeverity = "low"
	CvssSummarySeverity_Medium   CvssSummarySeverity = "medium"
	CvssSummarySeverity_None     CvssSummarySeverity = "none"
)

func PossibleValuesForCvssSummarySeverity() []string {
	return []string{
		string(CvssSummarySeverity_Critical),
		string(CvssSummarySeverity_High),
		string(CvssSummarySeverity_Low),
		string(CvssSummarySeverity_Medium),
		string(CvssSummarySeverity_None),
	}
}

func (s *CvssSummarySeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCvssSummarySeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCvssSummarySeverity(input string) (*CvssSummarySeverity, error) {
	vals := map[string]CvssSummarySeverity{
		"critical": CvssSummarySeverity_Critical,
		"high":     CvssSummarySeverity_High,
		"low":      CvssSummarySeverity_Low,
		"medium":   CvssSummarySeverity_Medium,
		"none":     CvssSummarySeverity_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CvssSummarySeverity(input)
	return &out, nil
}
