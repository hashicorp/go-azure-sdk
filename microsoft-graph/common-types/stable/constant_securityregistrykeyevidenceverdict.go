package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRegistryKeyEvidenceVerdict string

const (
	SecurityRegistryKeyEvidenceVerdict_Malicious      SecurityRegistryKeyEvidenceVerdict = "malicious"
	SecurityRegistryKeyEvidenceVerdict_NoThreatsFound SecurityRegistryKeyEvidenceVerdict = "noThreatsFound"
	SecurityRegistryKeyEvidenceVerdict_Suspicious     SecurityRegistryKeyEvidenceVerdict = "suspicious"
	SecurityRegistryKeyEvidenceVerdict_Unknown        SecurityRegistryKeyEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityRegistryKeyEvidenceVerdict() []string {
	return []string{
		string(SecurityRegistryKeyEvidenceVerdict_Malicious),
		string(SecurityRegistryKeyEvidenceVerdict_NoThreatsFound),
		string(SecurityRegistryKeyEvidenceVerdict_Suspicious),
		string(SecurityRegistryKeyEvidenceVerdict_Unknown),
	}
}

func (s *SecurityRegistryKeyEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRegistryKeyEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRegistryKeyEvidenceVerdict(input string) (*SecurityRegistryKeyEvidenceVerdict, error) {
	vals := map[string]SecurityRegistryKeyEvidenceVerdict{
		"malicious":      SecurityRegistryKeyEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityRegistryKeyEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityRegistryKeyEvidenceVerdict_Suspicious,
		"unknown":        SecurityRegistryKeyEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRegistryKeyEvidenceVerdict(input)
	return &out, nil
}
