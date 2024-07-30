package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityNetworkConnectionEvidenceVerdict string

const (
	SecurityNetworkConnectionEvidenceVerdict_Malicious      SecurityNetworkConnectionEvidenceVerdict = "malicious"
	SecurityNetworkConnectionEvidenceVerdict_NoThreatsFound SecurityNetworkConnectionEvidenceVerdict = "noThreatsFound"
	SecurityNetworkConnectionEvidenceVerdict_Suspicious     SecurityNetworkConnectionEvidenceVerdict = "suspicious"
	SecurityNetworkConnectionEvidenceVerdict_Unknown        SecurityNetworkConnectionEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityNetworkConnectionEvidenceVerdict() []string {
	return []string{
		string(SecurityNetworkConnectionEvidenceVerdict_Malicious),
		string(SecurityNetworkConnectionEvidenceVerdict_NoThreatsFound),
		string(SecurityNetworkConnectionEvidenceVerdict_Suspicious),
		string(SecurityNetworkConnectionEvidenceVerdict_Unknown),
	}
}

func (s *SecurityNetworkConnectionEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityNetworkConnectionEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityNetworkConnectionEvidenceVerdict(input string) (*SecurityNetworkConnectionEvidenceVerdict, error) {
	vals := map[string]SecurityNetworkConnectionEvidenceVerdict{
		"malicious":      SecurityNetworkConnectionEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityNetworkConnectionEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityNetworkConnectionEvidenceVerdict_Suspicious,
		"unknown":        SecurityNetworkConnectionEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityNetworkConnectionEvidenceVerdict(input)
	return &out, nil
}
