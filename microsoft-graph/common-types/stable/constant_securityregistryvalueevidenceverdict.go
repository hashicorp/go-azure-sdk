package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRegistryValueEvidenceVerdict string

const (
	SecurityRegistryValueEvidenceVerdict_Malicious      SecurityRegistryValueEvidenceVerdict = "malicious"
	SecurityRegistryValueEvidenceVerdict_NoThreatsFound SecurityRegistryValueEvidenceVerdict = "noThreatsFound"
	SecurityRegistryValueEvidenceVerdict_Suspicious     SecurityRegistryValueEvidenceVerdict = "suspicious"
	SecurityRegistryValueEvidenceVerdict_Unknown        SecurityRegistryValueEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityRegistryValueEvidenceVerdict() []string {
	return []string{
		string(SecurityRegistryValueEvidenceVerdict_Malicious),
		string(SecurityRegistryValueEvidenceVerdict_NoThreatsFound),
		string(SecurityRegistryValueEvidenceVerdict_Suspicious),
		string(SecurityRegistryValueEvidenceVerdict_Unknown),
	}
}

func (s *SecurityRegistryValueEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRegistryValueEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRegistryValueEvidenceVerdict(input string) (*SecurityRegistryValueEvidenceVerdict, error) {
	vals := map[string]SecurityRegistryValueEvidenceVerdict{
		"malicious":      SecurityRegistryValueEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityRegistryValueEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityRegistryValueEvidenceVerdict_Suspicious,
		"unknown":        SecurityRegistryValueEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRegistryValueEvidenceVerdict(input)
	return &out, nil
}
