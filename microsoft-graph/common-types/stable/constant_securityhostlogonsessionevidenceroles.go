package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostLogonSessionEvidenceRoles string

const (
	SecurityHostLogonSessionEvidenceRoles_Added             SecurityHostLogonSessionEvidenceRoles = "added"
	SecurityHostLogonSessionEvidenceRoles_Attacked          SecurityHostLogonSessionEvidenceRoles = "attacked"
	SecurityHostLogonSessionEvidenceRoles_Attacker          SecurityHostLogonSessionEvidenceRoles = "attacker"
	SecurityHostLogonSessionEvidenceRoles_CommandAndControl SecurityHostLogonSessionEvidenceRoles = "commandAndControl"
	SecurityHostLogonSessionEvidenceRoles_Compromised       SecurityHostLogonSessionEvidenceRoles = "compromised"
	SecurityHostLogonSessionEvidenceRoles_Contextual        SecurityHostLogonSessionEvidenceRoles = "contextual"
	SecurityHostLogonSessionEvidenceRoles_Created           SecurityHostLogonSessionEvidenceRoles = "created"
	SecurityHostLogonSessionEvidenceRoles_Destination       SecurityHostLogonSessionEvidenceRoles = "destination"
	SecurityHostLogonSessionEvidenceRoles_Edited            SecurityHostLogonSessionEvidenceRoles = "edited"
	SecurityHostLogonSessionEvidenceRoles_Loaded            SecurityHostLogonSessionEvidenceRoles = "loaded"
	SecurityHostLogonSessionEvidenceRoles_PolicyViolator    SecurityHostLogonSessionEvidenceRoles = "policyViolator"
	SecurityHostLogonSessionEvidenceRoles_Scanned           SecurityHostLogonSessionEvidenceRoles = "scanned"
	SecurityHostLogonSessionEvidenceRoles_Source            SecurityHostLogonSessionEvidenceRoles = "source"
	SecurityHostLogonSessionEvidenceRoles_Suspicious        SecurityHostLogonSessionEvidenceRoles = "suspicious"
	SecurityHostLogonSessionEvidenceRoles_Unknown           SecurityHostLogonSessionEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityHostLogonSessionEvidenceRoles() []string {
	return []string{
		string(SecurityHostLogonSessionEvidenceRoles_Added),
		string(SecurityHostLogonSessionEvidenceRoles_Attacked),
		string(SecurityHostLogonSessionEvidenceRoles_Attacker),
		string(SecurityHostLogonSessionEvidenceRoles_CommandAndControl),
		string(SecurityHostLogonSessionEvidenceRoles_Compromised),
		string(SecurityHostLogonSessionEvidenceRoles_Contextual),
		string(SecurityHostLogonSessionEvidenceRoles_Created),
		string(SecurityHostLogonSessionEvidenceRoles_Destination),
		string(SecurityHostLogonSessionEvidenceRoles_Edited),
		string(SecurityHostLogonSessionEvidenceRoles_Loaded),
		string(SecurityHostLogonSessionEvidenceRoles_PolicyViolator),
		string(SecurityHostLogonSessionEvidenceRoles_Scanned),
		string(SecurityHostLogonSessionEvidenceRoles_Source),
		string(SecurityHostLogonSessionEvidenceRoles_Suspicious),
		string(SecurityHostLogonSessionEvidenceRoles_Unknown),
	}
}

func (s *SecurityHostLogonSessionEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityHostLogonSessionEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityHostLogonSessionEvidenceRoles(input string) (*SecurityHostLogonSessionEvidenceRoles, error) {
	vals := map[string]SecurityHostLogonSessionEvidenceRoles{
		"added":             SecurityHostLogonSessionEvidenceRoles_Added,
		"attacked":          SecurityHostLogonSessionEvidenceRoles_Attacked,
		"attacker":          SecurityHostLogonSessionEvidenceRoles_Attacker,
		"commandandcontrol": SecurityHostLogonSessionEvidenceRoles_CommandAndControl,
		"compromised":       SecurityHostLogonSessionEvidenceRoles_Compromised,
		"contextual":        SecurityHostLogonSessionEvidenceRoles_Contextual,
		"created":           SecurityHostLogonSessionEvidenceRoles_Created,
		"destination":       SecurityHostLogonSessionEvidenceRoles_Destination,
		"edited":            SecurityHostLogonSessionEvidenceRoles_Edited,
		"loaded":            SecurityHostLogonSessionEvidenceRoles_Loaded,
		"policyviolator":    SecurityHostLogonSessionEvidenceRoles_PolicyViolator,
		"scanned":           SecurityHostLogonSessionEvidenceRoles_Scanned,
		"source":            SecurityHostLogonSessionEvidenceRoles_Source,
		"suspicious":        SecurityHostLogonSessionEvidenceRoles_Suspicious,
		"unknown":           SecurityHostLogonSessionEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHostLogonSessionEvidenceRoles(input)
	return &out, nil
}
