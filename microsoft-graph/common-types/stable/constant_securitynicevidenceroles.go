package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityNicEvidenceRoles string

const (
	SecurityNicEvidenceRoles_Added             SecurityNicEvidenceRoles = "added"
	SecurityNicEvidenceRoles_Attacked          SecurityNicEvidenceRoles = "attacked"
	SecurityNicEvidenceRoles_Attacker          SecurityNicEvidenceRoles = "attacker"
	SecurityNicEvidenceRoles_CommandAndControl SecurityNicEvidenceRoles = "commandAndControl"
	SecurityNicEvidenceRoles_Compromised       SecurityNicEvidenceRoles = "compromised"
	SecurityNicEvidenceRoles_Contextual        SecurityNicEvidenceRoles = "contextual"
	SecurityNicEvidenceRoles_Created           SecurityNicEvidenceRoles = "created"
	SecurityNicEvidenceRoles_Destination       SecurityNicEvidenceRoles = "destination"
	SecurityNicEvidenceRoles_Edited            SecurityNicEvidenceRoles = "edited"
	SecurityNicEvidenceRoles_Loaded            SecurityNicEvidenceRoles = "loaded"
	SecurityNicEvidenceRoles_PolicyViolator    SecurityNicEvidenceRoles = "policyViolator"
	SecurityNicEvidenceRoles_Scanned           SecurityNicEvidenceRoles = "scanned"
	SecurityNicEvidenceRoles_Source            SecurityNicEvidenceRoles = "source"
	SecurityNicEvidenceRoles_Suspicious        SecurityNicEvidenceRoles = "suspicious"
	SecurityNicEvidenceRoles_Unknown           SecurityNicEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityNicEvidenceRoles() []string {
	return []string{
		string(SecurityNicEvidenceRoles_Added),
		string(SecurityNicEvidenceRoles_Attacked),
		string(SecurityNicEvidenceRoles_Attacker),
		string(SecurityNicEvidenceRoles_CommandAndControl),
		string(SecurityNicEvidenceRoles_Compromised),
		string(SecurityNicEvidenceRoles_Contextual),
		string(SecurityNicEvidenceRoles_Created),
		string(SecurityNicEvidenceRoles_Destination),
		string(SecurityNicEvidenceRoles_Edited),
		string(SecurityNicEvidenceRoles_Loaded),
		string(SecurityNicEvidenceRoles_PolicyViolator),
		string(SecurityNicEvidenceRoles_Scanned),
		string(SecurityNicEvidenceRoles_Source),
		string(SecurityNicEvidenceRoles_Suspicious),
		string(SecurityNicEvidenceRoles_Unknown),
	}
}

func (s *SecurityNicEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityNicEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityNicEvidenceRoles(input string) (*SecurityNicEvidenceRoles, error) {
	vals := map[string]SecurityNicEvidenceRoles{
		"added":             SecurityNicEvidenceRoles_Added,
		"attacked":          SecurityNicEvidenceRoles_Attacked,
		"attacker":          SecurityNicEvidenceRoles_Attacker,
		"commandandcontrol": SecurityNicEvidenceRoles_CommandAndControl,
		"compromised":       SecurityNicEvidenceRoles_Compromised,
		"contextual":        SecurityNicEvidenceRoles_Contextual,
		"created":           SecurityNicEvidenceRoles_Created,
		"destination":       SecurityNicEvidenceRoles_Destination,
		"edited":            SecurityNicEvidenceRoles_Edited,
		"loaded":            SecurityNicEvidenceRoles_Loaded,
		"policyviolator":    SecurityNicEvidenceRoles_PolicyViolator,
		"scanned":           SecurityNicEvidenceRoles_Scanned,
		"source":            SecurityNicEvidenceRoles_Source,
		"suspicious":        SecurityNicEvidenceRoles_Suspicious,
		"unknown":           SecurityNicEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityNicEvidenceRoles(input)
	return &out, nil
}
