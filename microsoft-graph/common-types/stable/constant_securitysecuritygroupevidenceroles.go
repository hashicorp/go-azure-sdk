package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySecurityGroupEvidenceRoles string

const (
	SecuritySecurityGroupEvidenceRoles_Added             SecuritySecurityGroupEvidenceRoles = "added"
	SecuritySecurityGroupEvidenceRoles_Attacked          SecuritySecurityGroupEvidenceRoles = "attacked"
	SecuritySecurityGroupEvidenceRoles_Attacker          SecuritySecurityGroupEvidenceRoles = "attacker"
	SecuritySecurityGroupEvidenceRoles_CommandAndControl SecuritySecurityGroupEvidenceRoles = "commandAndControl"
	SecuritySecurityGroupEvidenceRoles_Compromised       SecuritySecurityGroupEvidenceRoles = "compromised"
	SecuritySecurityGroupEvidenceRoles_Contextual        SecuritySecurityGroupEvidenceRoles = "contextual"
	SecuritySecurityGroupEvidenceRoles_Created           SecuritySecurityGroupEvidenceRoles = "created"
	SecuritySecurityGroupEvidenceRoles_Destination       SecuritySecurityGroupEvidenceRoles = "destination"
	SecuritySecurityGroupEvidenceRoles_Edited            SecuritySecurityGroupEvidenceRoles = "edited"
	SecuritySecurityGroupEvidenceRoles_Loaded            SecuritySecurityGroupEvidenceRoles = "loaded"
	SecuritySecurityGroupEvidenceRoles_PolicyViolator    SecuritySecurityGroupEvidenceRoles = "policyViolator"
	SecuritySecurityGroupEvidenceRoles_Scanned           SecuritySecurityGroupEvidenceRoles = "scanned"
	SecuritySecurityGroupEvidenceRoles_Source            SecuritySecurityGroupEvidenceRoles = "source"
	SecuritySecurityGroupEvidenceRoles_Suspicious        SecuritySecurityGroupEvidenceRoles = "suspicious"
	SecuritySecurityGroupEvidenceRoles_Unknown           SecuritySecurityGroupEvidenceRoles = "unknown"
)

func PossibleValuesForSecuritySecurityGroupEvidenceRoles() []string {
	return []string{
		string(SecuritySecurityGroupEvidenceRoles_Added),
		string(SecuritySecurityGroupEvidenceRoles_Attacked),
		string(SecuritySecurityGroupEvidenceRoles_Attacker),
		string(SecuritySecurityGroupEvidenceRoles_CommandAndControl),
		string(SecuritySecurityGroupEvidenceRoles_Compromised),
		string(SecuritySecurityGroupEvidenceRoles_Contextual),
		string(SecuritySecurityGroupEvidenceRoles_Created),
		string(SecuritySecurityGroupEvidenceRoles_Destination),
		string(SecuritySecurityGroupEvidenceRoles_Edited),
		string(SecuritySecurityGroupEvidenceRoles_Loaded),
		string(SecuritySecurityGroupEvidenceRoles_PolicyViolator),
		string(SecuritySecurityGroupEvidenceRoles_Scanned),
		string(SecuritySecurityGroupEvidenceRoles_Source),
		string(SecuritySecurityGroupEvidenceRoles_Suspicious),
		string(SecuritySecurityGroupEvidenceRoles_Unknown),
	}
}

func (s *SecuritySecurityGroupEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySecurityGroupEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySecurityGroupEvidenceRoles(input string) (*SecuritySecurityGroupEvidenceRoles, error) {
	vals := map[string]SecuritySecurityGroupEvidenceRoles{
		"added":             SecuritySecurityGroupEvidenceRoles_Added,
		"attacked":          SecuritySecurityGroupEvidenceRoles_Attacked,
		"attacker":          SecuritySecurityGroupEvidenceRoles_Attacker,
		"commandandcontrol": SecuritySecurityGroupEvidenceRoles_CommandAndControl,
		"compromised":       SecuritySecurityGroupEvidenceRoles_Compromised,
		"contextual":        SecuritySecurityGroupEvidenceRoles_Contextual,
		"created":           SecuritySecurityGroupEvidenceRoles_Created,
		"destination":       SecuritySecurityGroupEvidenceRoles_Destination,
		"edited":            SecuritySecurityGroupEvidenceRoles_Edited,
		"loaded":            SecuritySecurityGroupEvidenceRoles_Loaded,
		"policyviolator":    SecuritySecurityGroupEvidenceRoles_PolicyViolator,
		"scanned":           SecuritySecurityGroupEvidenceRoles_Scanned,
		"source":            SecuritySecurityGroupEvidenceRoles_Source,
		"suspicious":        SecuritySecurityGroupEvidenceRoles_Suspicious,
		"unknown":           SecuritySecurityGroupEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySecurityGroupEvidenceRoles(input)
	return &out, nil
}
