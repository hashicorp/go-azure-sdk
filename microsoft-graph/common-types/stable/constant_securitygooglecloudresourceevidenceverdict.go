package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGoogleCloudResourceEvidenceVerdict string

const (
	SecurityGoogleCloudResourceEvidenceVerdict_Malicious      SecurityGoogleCloudResourceEvidenceVerdict = "malicious"
	SecurityGoogleCloudResourceEvidenceVerdict_NoThreatsFound SecurityGoogleCloudResourceEvidenceVerdict = "noThreatsFound"
	SecurityGoogleCloudResourceEvidenceVerdict_Suspicious     SecurityGoogleCloudResourceEvidenceVerdict = "suspicious"
	SecurityGoogleCloudResourceEvidenceVerdict_Unknown        SecurityGoogleCloudResourceEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityGoogleCloudResourceEvidenceVerdict() []string {
	return []string{
		string(SecurityGoogleCloudResourceEvidenceVerdict_Malicious),
		string(SecurityGoogleCloudResourceEvidenceVerdict_NoThreatsFound),
		string(SecurityGoogleCloudResourceEvidenceVerdict_Suspicious),
		string(SecurityGoogleCloudResourceEvidenceVerdict_Unknown),
	}
}

func (s *SecurityGoogleCloudResourceEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityGoogleCloudResourceEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityGoogleCloudResourceEvidenceVerdict(input string) (*SecurityGoogleCloudResourceEvidenceVerdict, error) {
	vals := map[string]SecurityGoogleCloudResourceEvidenceVerdict{
		"malicious":      SecurityGoogleCloudResourceEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityGoogleCloudResourceEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityGoogleCloudResourceEvidenceVerdict_Suspicious,
		"unknown":        SecurityGoogleCloudResourceEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGoogleCloudResourceEvidenceVerdict(input)
	return &out, nil
}
