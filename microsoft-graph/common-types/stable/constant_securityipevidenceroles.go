package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIpEvidenceRoles string

const (
	SecurityIpEvidenceRoles_Added             SecurityIpEvidenceRoles = "added"
	SecurityIpEvidenceRoles_Attacked          SecurityIpEvidenceRoles = "attacked"
	SecurityIpEvidenceRoles_Attacker          SecurityIpEvidenceRoles = "attacker"
	SecurityIpEvidenceRoles_CommandAndControl SecurityIpEvidenceRoles = "commandAndControl"
	SecurityIpEvidenceRoles_Compromised       SecurityIpEvidenceRoles = "compromised"
	SecurityIpEvidenceRoles_Contextual        SecurityIpEvidenceRoles = "contextual"
	SecurityIpEvidenceRoles_Created           SecurityIpEvidenceRoles = "created"
	SecurityIpEvidenceRoles_Destination       SecurityIpEvidenceRoles = "destination"
	SecurityIpEvidenceRoles_Edited            SecurityIpEvidenceRoles = "edited"
	SecurityIpEvidenceRoles_Loaded            SecurityIpEvidenceRoles = "loaded"
	SecurityIpEvidenceRoles_PolicyViolator    SecurityIpEvidenceRoles = "policyViolator"
	SecurityIpEvidenceRoles_Scanned           SecurityIpEvidenceRoles = "scanned"
	SecurityIpEvidenceRoles_Source            SecurityIpEvidenceRoles = "source"
	SecurityIpEvidenceRoles_Suspicious        SecurityIpEvidenceRoles = "suspicious"
	SecurityIpEvidenceRoles_Unknown           SecurityIpEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityIpEvidenceRoles() []string {
	return []string{
		string(SecurityIpEvidenceRoles_Added),
		string(SecurityIpEvidenceRoles_Attacked),
		string(SecurityIpEvidenceRoles_Attacker),
		string(SecurityIpEvidenceRoles_CommandAndControl),
		string(SecurityIpEvidenceRoles_Compromised),
		string(SecurityIpEvidenceRoles_Contextual),
		string(SecurityIpEvidenceRoles_Created),
		string(SecurityIpEvidenceRoles_Destination),
		string(SecurityIpEvidenceRoles_Edited),
		string(SecurityIpEvidenceRoles_Loaded),
		string(SecurityIpEvidenceRoles_PolicyViolator),
		string(SecurityIpEvidenceRoles_Scanned),
		string(SecurityIpEvidenceRoles_Source),
		string(SecurityIpEvidenceRoles_Suspicious),
		string(SecurityIpEvidenceRoles_Unknown),
	}
}

func (s *SecurityIpEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIpEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIpEvidenceRoles(input string) (*SecurityIpEvidenceRoles, error) {
	vals := map[string]SecurityIpEvidenceRoles{
		"added":             SecurityIpEvidenceRoles_Added,
		"attacked":          SecurityIpEvidenceRoles_Attacked,
		"attacker":          SecurityIpEvidenceRoles_Attacker,
		"commandandcontrol": SecurityIpEvidenceRoles_CommandAndControl,
		"compromised":       SecurityIpEvidenceRoles_Compromised,
		"contextual":        SecurityIpEvidenceRoles_Contextual,
		"created":           SecurityIpEvidenceRoles_Created,
		"destination":       SecurityIpEvidenceRoles_Destination,
		"edited":            SecurityIpEvidenceRoles_Edited,
		"loaded":            SecurityIpEvidenceRoles_Loaded,
		"policyviolator":    SecurityIpEvidenceRoles_PolicyViolator,
		"scanned":           SecurityIpEvidenceRoles_Scanned,
		"source":            SecurityIpEvidenceRoles_Source,
		"suspicious":        SecurityIpEvidenceRoles_Suspicious,
		"unknown":           SecurityIpEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIpEvidenceRoles(input)
	return &out, nil
}
