package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesControllerEvidenceVerdict string

const (
	SecurityKubernetesControllerEvidenceVerdict_Malicious      SecurityKubernetesControllerEvidenceVerdict = "malicious"
	SecurityKubernetesControllerEvidenceVerdict_NoThreatsFound SecurityKubernetesControllerEvidenceVerdict = "noThreatsFound"
	SecurityKubernetesControllerEvidenceVerdict_Suspicious     SecurityKubernetesControllerEvidenceVerdict = "suspicious"
	SecurityKubernetesControllerEvidenceVerdict_Unknown        SecurityKubernetesControllerEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityKubernetesControllerEvidenceVerdict() []string {
	return []string{
		string(SecurityKubernetesControllerEvidenceVerdict_Malicious),
		string(SecurityKubernetesControllerEvidenceVerdict_NoThreatsFound),
		string(SecurityKubernetesControllerEvidenceVerdict_Suspicious),
		string(SecurityKubernetesControllerEvidenceVerdict_Unknown),
	}
}

func (s *SecurityKubernetesControllerEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesControllerEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesControllerEvidenceVerdict(input string) (*SecurityKubernetesControllerEvidenceVerdict, error) {
	vals := map[string]SecurityKubernetesControllerEvidenceVerdict{
		"malicious":      SecurityKubernetesControllerEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityKubernetesControllerEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityKubernetesControllerEvidenceVerdict_Suspicious,
		"unknown":        SecurityKubernetesControllerEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesControllerEvidenceVerdict(input)
	return &out, nil
}
