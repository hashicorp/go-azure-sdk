package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailClusterEvidenceVerdict string

const (
	SecurityMailClusterEvidenceVerdict_Malicious      SecurityMailClusterEvidenceVerdict = "malicious"
	SecurityMailClusterEvidenceVerdict_NoThreatsFound SecurityMailClusterEvidenceVerdict = "noThreatsFound"
	SecurityMailClusterEvidenceVerdict_Suspicious     SecurityMailClusterEvidenceVerdict = "suspicious"
	SecurityMailClusterEvidenceVerdict_Unknown        SecurityMailClusterEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityMailClusterEvidenceVerdict() []string {
	return []string{
		string(SecurityMailClusterEvidenceVerdict_Malicious),
		string(SecurityMailClusterEvidenceVerdict_NoThreatsFound),
		string(SecurityMailClusterEvidenceVerdict_Suspicious),
		string(SecurityMailClusterEvidenceVerdict_Unknown),
	}
}

func (s *SecurityMailClusterEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityMailClusterEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityMailClusterEvidenceVerdict(input string) (*SecurityMailClusterEvidenceVerdict, error) {
	vals := map[string]SecurityMailClusterEvidenceVerdict{
		"malicious":      SecurityMailClusterEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityMailClusterEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityMailClusterEvidenceVerdict_Suspicious,
		"unknown":        SecurityMailClusterEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMailClusterEvidenceVerdict(input)
	return &out, nil
}
