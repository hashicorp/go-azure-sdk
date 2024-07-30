package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGitHubOrganizationEvidenceVerdict string

const (
	SecurityGitHubOrganizationEvidenceVerdict_Malicious      SecurityGitHubOrganizationEvidenceVerdict = "malicious"
	SecurityGitHubOrganizationEvidenceVerdict_NoThreatsFound SecurityGitHubOrganizationEvidenceVerdict = "noThreatsFound"
	SecurityGitHubOrganizationEvidenceVerdict_Suspicious     SecurityGitHubOrganizationEvidenceVerdict = "suspicious"
	SecurityGitHubOrganizationEvidenceVerdict_Unknown        SecurityGitHubOrganizationEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityGitHubOrganizationEvidenceVerdict() []string {
	return []string{
		string(SecurityGitHubOrganizationEvidenceVerdict_Malicious),
		string(SecurityGitHubOrganizationEvidenceVerdict_NoThreatsFound),
		string(SecurityGitHubOrganizationEvidenceVerdict_Suspicious),
		string(SecurityGitHubOrganizationEvidenceVerdict_Unknown),
	}
}

func (s *SecurityGitHubOrganizationEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityGitHubOrganizationEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityGitHubOrganizationEvidenceVerdict(input string) (*SecurityGitHubOrganizationEvidenceVerdict, error) {
	vals := map[string]SecurityGitHubOrganizationEvidenceVerdict{
		"malicious":      SecurityGitHubOrganizationEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityGitHubOrganizationEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityGitHubOrganizationEvidenceVerdict_Suspicious,
		"unknown":        SecurityGitHubOrganizationEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGitHubOrganizationEvidenceVerdict(input)
	return &out, nil
}
