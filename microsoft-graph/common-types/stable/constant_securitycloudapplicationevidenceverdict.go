package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCloudApplicationEvidenceVerdict string

const (
	SecurityCloudApplicationEvidenceVerdict_Malicious      SecurityCloudApplicationEvidenceVerdict = "malicious"
	SecurityCloudApplicationEvidenceVerdict_NoThreatsFound SecurityCloudApplicationEvidenceVerdict = "noThreatsFound"
	SecurityCloudApplicationEvidenceVerdict_Suspicious     SecurityCloudApplicationEvidenceVerdict = "suspicious"
	SecurityCloudApplicationEvidenceVerdict_Unknown        SecurityCloudApplicationEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityCloudApplicationEvidenceVerdict() []string {
	return []string{
		string(SecurityCloudApplicationEvidenceVerdict_Malicious),
		string(SecurityCloudApplicationEvidenceVerdict_NoThreatsFound),
		string(SecurityCloudApplicationEvidenceVerdict_Suspicious),
		string(SecurityCloudApplicationEvidenceVerdict_Unknown),
	}
}

func (s *SecurityCloudApplicationEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityCloudApplicationEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityCloudApplicationEvidenceVerdict(input string) (*SecurityCloudApplicationEvidenceVerdict, error) {
	vals := map[string]SecurityCloudApplicationEvidenceVerdict{
		"malicious":      SecurityCloudApplicationEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityCloudApplicationEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityCloudApplicationEvidenceVerdict_Suspicious,
		"unknown":        SecurityCloudApplicationEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCloudApplicationEvidenceVerdict(input)
	return &out, nil
}
