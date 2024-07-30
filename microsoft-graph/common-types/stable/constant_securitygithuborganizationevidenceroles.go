package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGitHubOrganizationEvidenceRoles string

const (
	SecurityGitHubOrganizationEvidenceRoles_Added             SecurityGitHubOrganizationEvidenceRoles = "added"
	SecurityGitHubOrganizationEvidenceRoles_Attacked          SecurityGitHubOrganizationEvidenceRoles = "attacked"
	SecurityGitHubOrganizationEvidenceRoles_Attacker          SecurityGitHubOrganizationEvidenceRoles = "attacker"
	SecurityGitHubOrganizationEvidenceRoles_CommandAndControl SecurityGitHubOrganizationEvidenceRoles = "commandAndControl"
	SecurityGitHubOrganizationEvidenceRoles_Compromised       SecurityGitHubOrganizationEvidenceRoles = "compromised"
	SecurityGitHubOrganizationEvidenceRoles_Contextual        SecurityGitHubOrganizationEvidenceRoles = "contextual"
	SecurityGitHubOrganizationEvidenceRoles_Created           SecurityGitHubOrganizationEvidenceRoles = "created"
	SecurityGitHubOrganizationEvidenceRoles_Destination       SecurityGitHubOrganizationEvidenceRoles = "destination"
	SecurityGitHubOrganizationEvidenceRoles_Edited            SecurityGitHubOrganizationEvidenceRoles = "edited"
	SecurityGitHubOrganizationEvidenceRoles_Loaded            SecurityGitHubOrganizationEvidenceRoles = "loaded"
	SecurityGitHubOrganizationEvidenceRoles_PolicyViolator    SecurityGitHubOrganizationEvidenceRoles = "policyViolator"
	SecurityGitHubOrganizationEvidenceRoles_Scanned           SecurityGitHubOrganizationEvidenceRoles = "scanned"
	SecurityGitHubOrganizationEvidenceRoles_Source            SecurityGitHubOrganizationEvidenceRoles = "source"
	SecurityGitHubOrganizationEvidenceRoles_Suspicious        SecurityGitHubOrganizationEvidenceRoles = "suspicious"
	SecurityGitHubOrganizationEvidenceRoles_Unknown           SecurityGitHubOrganizationEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityGitHubOrganizationEvidenceRoles() []string {
	return []string{
		string(SecurityGitHubOrganizationEvidenceRoles_Added),
		string(SecurityGitHubOrganizationEvidenceRoles_Attacked),
		string(SecurityGitHubOrganizationEvidenceRoles_Attacker),
		string(SecurityGitHubOrganizationEvidenceRoles_CommandAndControl),
		string(SecurityGitHubOrganizationEvidenceRoles_Compromised),
		string(SecurityGitHubOrganizationEvidenceRoles_Contextual),
		string(SecurityGitHubOrganizationEvidenceRoles_Created),
		string(SecurityGitHubOrganizationEvidenceRoles_Destination),
		string(SecurityGitHubOrganizationEvidenceRoles_Edited),
		string(SecurityGitHubOrganizationEvidenceRoles_Loaded),
		string(SecurityGitHubOrganizationEvidenceRoles_PolicyViolator),
		string(SecurityGitHubOrganizationEvidenceRoles_Scanned),
		string(SecurityGitHubOrganizationEvidenceRoles_Source),
		string(SecurityGitHubOrganizationEvidenceRoles_Suspicious),
		string(SecurityGitHubOrganizationEvidenceRoles_Unknown),
	}
}

func (s *SecurityGitHubOrganizationEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityGitHubOrganizationEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityGitHubOrganizationEvidenceRoles(input string) (*SecurityGitHubOrganizationEvidenceRoles, error) {
	vals := map[string]SecurityGitHubOrganizationEvidenceRoles{
		"added":             SecurityGitHubOrganizationEvidenceRoles_Added,
		"attacked":          SecurityGitHubOrganizationEvidenceRoles_Attacked,
		"attacker":          SecurityGitHubOrganizationEvidenceRoles_Attacker,
		"commandandcontrol": SecurityGitHubOrganizationEvidenceRoles_CommandAndControl,
		"compromised":       SecurityGitHubOrganizationEvidenceRoles_Compromised,
		"contextual":        SecurityGitHubOrganizationEvidenceRoles_Contextual,
		"created":           SecurityGitHubOrganizationEvidenceRoles_Created,
		"destination":       SecurityGitHubOrganizationEvidenceRoles_Destination,
		"edited":            SecurityGitHubOrganizationEvidenceRoles_Edited,
		"loaded":            SecurityGitHubOrganizationEvidenceRoles_Loaded,
		"policyviolator":    SecurityGitHubOrganizationEvidenceRoles_PolicyViolator,
		"scanned":           SecurityGitHubOrganizationEvidenceRoles_Scanned,
		"source":            SecurityGitHubOrganizationEvidenceRoles_Source,
		"suspicious":        SecurityGitHubOrganizationEvidenceRoles_Suspicious,
		"unknown":           SecurityGitHubOrganizationEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGitHubOrganizationEvidenceRoles(input)
	return &out, nil
}
