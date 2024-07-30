package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceEvidenceVerdict string

const (
	SecurityKubernetesServiceEvidenceVerdict_Malicious      SecurityKubernetesServiceEvidenceVerdict = "malicious"
	SecurityKubernetesServiceEvidenceVerdict_NoThreatsFound SecurityKubernetesServiceEvidenceVerdict = "noThreatsFound"
	SecurityKubernetesServiceEvidenceVerdict_Suspicious     SecurityKubernetesServiceEvidenceVerdict = "suspicious"
	SecurityKubernetesServiceEvidenceVerdict_Unknown        SecurityKubernetesServiceEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityKubernetesServiceEvidenceVerdict() []string {
	return []string{
		string(SecurityKubernetesServiceEvidenceVerdict_Malicious),
		string(SecurityKubernetesServiceEvidenceVerdict_NoThreatsFound),
		string(SecurityKubernetesServiceEvidenceVerdict_Suspicious),
		string(SecurityKubernetesServiceEvidenceVerdict_Unknown),
	}
}

func (s *SecurityKubernetesServiceEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesServiceEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesServiceEvidenceVerdict(input string) (*SecurityKubernetesServiceEvidenceVerdict, error) {
	vals := map[string]SecurityKubernetesServiceEvidenceVerdict{
		"malicious":      SecurityKubernetesServiceEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityKubernetesServiceEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityKubernetesServiceEvidenceVerdict_Suspicious,
		"unknown":        SecurityKubernetesServiceEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServiceEvidenceVerdict(input)
	return &out, nil
}
