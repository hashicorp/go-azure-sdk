package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesClusterEvidenceVerdict string

const (
	SecurityKubernetesClusterEvidenceVerdict_Malicious      SecurityKubernetesClusterEvidenceVerdict = "malicious"
	SecurityKubernetesClusterEvidenceVerdict_NoThreatsFound SecurityKubernetesClusterEvidenceVerdict = "noThreatsFound"
	SecurityKubernetesClusterEvidenceVerdict_Suspicious     SecurityKubernetesClusterEvidenceVerdict = "suspicious"
	SecurityKubernetesClusterEvidenceVerdict_Unknown        SecurityKubernetesClusterEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityKubernetesClusterEvidenceVerdict() []string {
	return []string{
		string(SecurityKubernetesClusterEvidenceVerdict_Malicious),
		string(SecurityKubernetesClusterEvidenceVerdict_NoThreatsFound),
		string(SecurityKubernetesClusterEvidenceVerdict_Suspicious),
		string(SecurityKubernetesClusterEvidenceVerdict_Unknown),
	}
}

func (s *SecurityKubernetesClusterEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesClusterEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesClusterEvidenceVerdict(input string) (*SecurityKubernetesClusterEvidenceVerdict, error) {
	vals := map[string]SecurityKubernetesClusterEvidenceVerdict{
		"malicious":      SecurityKubernetesClusterEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityKubernetesClusterEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityKubernetesClusterEvidenceVerdict_Suspicious,
		"unknown":        SecurityKubernetesClusterEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesClusterEvidenceVerdict(input)
	return &out, nil
}
