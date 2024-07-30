package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesSecretEvidenceVerdict string

const (
	SecurityKubernetesSecretEvidenceVerdict_Malicious      SecurityKubernetesSecretEvidenceVerdict = "malicious"
	SecurityKubernetesSecretEvidenceVerdict_NoThreatsFound SecurityKubernetesSecretEvidenceVerdict = "noThreatsFound"
	SecurityKubernetesSecretEvidenceVerdict_Suspicious     SecurityKubernetesSecretEvidenceVerdict = "suspicious"
	SecurityKubernetesSecretEvidenceVerdict_Unknown        SecurityKubernetesSecretEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityKubernetesSecretEvidenceVerdict() []string {
	return []string{
		string(SecurityKubernetesSecretEvidenceVerdict_Malicious),
		string(SecurityKubernetesSecretEvidenceVerdict_NoThreatsFound),
		string(SecurityKubernetesSecretEvidenceVerdict_Suspicious),
		string(SecurityKubernetesSecretEvidenceVerdict_Unknown),
	}
}

func (s *SecurityKubernetesSecretEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesSecretEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesSecretEvidenceVerdict(input string) (*SecurityKubernetesSecretEvidenceVerdict, error) {
	vals := map[string]SecurityKubernetesSecretEvidenceVerdict{
		"malicious":      SecurityKubernetesSecretEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityKubernetesSecretEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityKubernetesSecretEvidenceVerdict_Suspicious,
		"unknown":        SecurityKubernetesSecretEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesSecretEvidenceVerdict(input)
	return &out, nil
}
