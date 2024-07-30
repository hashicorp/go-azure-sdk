package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGitHubUserEvidenceRoles string

const (
	SecurityGitHubUserEvidenceRoles_Added             SecurityGitHubUserEvidenceRoles = "added"
	SecurityGitHubUserEvidenceRoles_Attacked          SecurityGitHubUserEvidenceRoles = "attacked"
	SecurityGitHubUserEvidenceRoles_Attacker          SecurityGitHubUserEvidenceRoles = "attacker"
	SecurityGitHubUserEvidenceRoles_CommandAndControl SecurityGitHubUserEvidenceRoles = "commandAndControl"
	SecurityGitHubUserEvidenceRoles_Compromised       SecurityGitHubUserEvidenceRoles = "compromised"
	SecurityGitHubUserEvidenceRoles_Contextual        SecurityGitHubUserEvidenceRoles = "contextual"
	SecurityGitHubUserEvidenceRoles_Created           SecurityGitHubUserEvidenceRoles = "created"
	SecurityGitHubUserEvidenceRoles_Destination       SecurityGitHubUserEvidenceRoles = "destination"
	SecurityGitHubUserEvidenceRoles_Edited            SecurityGitHubUserEvidenceRoles = "edited"
	SecurityGitHubUserEvidenceRoles_Loaded            SecurityGitHubUserEvidenceRoles = "loaded"
	SecurityGitHubUserEvidenceRoles_PolicyViolator    SecurityGitHubUserEvidenceRoles = "policyViolator"
	SecurityGitHubUserEvidenceRoles_Scanned           SecurityGitHubUserEvidenceRoles = "scanned"
	SecurityGitHubUserEvidenceRoles_Source            SecurityGitHubUserEvidenceRoles = "source"
	SecurityGitHubUserEvidenceRoles_Suspicious        SecurityGitHubUserEvidenceRoles = "suspicious"
	SecurityGitHubUserEvidenceRoles_Unknown           SecurityGitHubUserEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityGitHubUserEvidenceRoles() []string {
	return []string{
		string(SecurityGitHubUserEvidenceRoles_Added),
		string(SecurityGitHubUserEvidenceRoles_Attacked),
		string(SecurityGitHubUserEvidenceRoles_Attacker),
		string(SecurityGitHubUserEvidenceRoles_CommandAndControl),
		string(SecurityGitHubUserEvidenceRoles_Compromised),
		string(SecurityGitHubUserEvidenceRoles_Contextual),
		string(SecurityGitHubUserEvidenceRoles_Created),
		string(SecurityGitHubUserEvidenceRoles_Destination),
		string(SecurityGitHubUserEvidenceRoles_Edited),
		string(SecurityGitHubUserEvidenceRoles_Loaded),
		string(SecurityGitHubUserEvidenceRoles_PolicyViolator),
		string(SecurityGitHubUserEvidenceRoles_Scanned),
		string(SecurityGitHubUserEvidenceRoles_Source),
		string(SecurityGitHubUserEvidenceRoles_Suspicious),
		string(SecurityGitHubUserEvidenceRoles_Unknown),
	}
}

func (s *SecurityGitHubUserEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityGitHubUserEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityGitHubUserEvidenceRoles(input string) (*SecurityGitHubUserEvidenceRoles, error) {
	vals := map[string]SecurityGitHubUserEvidenceRoles{
		"added":             SecurityGitHubUserEvidenceRoles_Added,
		"attacked":          SecurityGitHubUserEvidenceRoles_Attacked,
		"attacker":          SecurityGitHubUserEvidenceRoles_Attacker,
		"commandandcontrol": SecurityGitHubUserEvidenceRoles_CommandAndControl,
		"compromised":       SecurityGitHubUserEvidenceRoles_Compromised,
		"contextual":        SecurityGitHubUserEvidenceRoles_Contextual,
		"created":           SecurityGitHubUserEvidenceRoles_Created,
		"destination":       SecurityGitHubUserEvidenceRoles_Destination,
		"edited":            SecurityGitHubUserEvidenceRoles_Edited,
		"loaded":            SecurityGitHubUserEvidenceRoles_Loaded,
		"policyviolator":    SecurityGitHubUserEvidenceRoles_PolicyViolator,
		"scanned":           SecurityGitHubUserEvidenceRoles_Scanned,
		"source":            SecurityGitHubUserEvidenceRoles_Source,
		"suspicious":        SecurityGitHubUserEvidenceRoles_Suspicious,
		"unknown":           SecurityGitHubUserEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGitHubUserEvidenceRoles(input)
	return &out, nil
}
