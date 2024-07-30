package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRegistryKeyEvidenceRoles string

const (
	SecurityRegistryKeyEvidenceRoles_Added             SecurityRegistryKeyEvidenceRoles = "added"
	SecurityRegistryKeyEvidenceRoles_Attacked          SecurityRegistryKeyEvidenceRoles = "attacked"
	SecurityRegistryKeyEvidenceRoles_Attacker          SecurityRegistryKeyEvidenceRoles = "attacker"
	SecurityRegistryKeyEvidenceRoles_CommandAndControl SecurityRegistryKeyEvidenceRoles = "commandAndControl"
	SecurityRegistryKeyEvidenceRoles_Compromised       SecurityRegistryKeyEvidenceRoles = "compromised"
	SecurityRegistryKeyEvidenceRoles_Contextual        SecurityRegistryKeyEvidenceRoles = "contextual"
	SecurityRegistryKeyEvidenceRoles_Created           SecurityRegistryKeyEvidenceRoles = "created"
	SecurityRegistryKeyEvidenceRoles_Destination       SecurityRegistryKeyEvidenceRoles = "destination"
	SecurityRegistryKeyEvidenceRoles_Edited            SecurityRegistryKeyEvidenceRoles = "edited"
	SecurityRegistryKeyEvidenceRoles_Loaded            SecurityRegistryKeyEvidenceRoles = "loaded"
	SecurityRegistryKeyEvidenceRoles_PolicyViolator    SecurityRegistryKeyEvidenceRoles = "policyViolator"
	SecurityRegistryKeyEvidenceRoles_Scanned           SecurityRegistryKeyEvidenceRoles = "scanned"
	SecurityRegistryKeyEvidenceRoles_Source            SecurityRegistryKeyEvidenceRoles = "source"
	SecurityRegistryKeyEvidenceRoles_Suspicious        SecurityRegistryKeyEvidenceRoles = "suspicious"
	SecurityRegistryKeyEvidenceRoles_Unknown           SecurityRegistryKeyEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityRegistryKeyEvidenceRoles() []string {
	return []string{
		string(SecurityRegistryKeyEvidenceRoles_Added),
		string(SecurityRegistryKeyEvidenceRoles_Attacked),
		string(SecurityRegistryKeyEvidenceRoles_Attacker),
		string(SecurityRegistryKeyEvidenceRoles_CommandAndControl),
		string(SecurityRegistryKeyEvidenceRoles_Compromised),
		string(SecurityRegistryKeyEvidenceRoles_Contextual),
		string(SecurityRegistryKeyEvidenceRoles_Created),
		string(SecurityRegistryKeyEvidenceRoles_Destination),
		string(SecurityRegistryKeyEvidenceRoles_Edited),
		string(SecurityRegistryKeyEvidenceRoles_Loaded),
		string(SecurityRegistryKeyEvidenceRoles_PolicyViolator),
		string(SecurityRegistryKeyEvidenceRoles_Scanned),
		string(SecurityRegistryKeyEvidenceRoles_Source),
		string(SecurityRegistryKeyEvidenceRoles_Suspicious),
		string(SecurityRegistryKeyEvidenceRoles_Unknown),
	}
}

func (s *SecurityRegistryKeyEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRegistryKeyEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRegistryKeyEvidenceRoles(input string) (*SecurityRegistryKeyEvidenceRoles, error) {
	vals := map[string]SecurityRegistryKeyEvidenceRoles{
		"added":             SecurityRegistryKeyEvidenceRoles_Added,
		"attacked":          SecurityRegistryKeyEvidenceRoles_Attacked,
		"attacker":          SecurityRegistryKeyEvidenceRoles_Attacker,
		"commandandcontrol": SecurityRegistryKeyEvidenceRoles_CommandAndControl,
		"compromised":       SecurityRegistryKeyEvidenceRoles_Compromised,
		"contextual":        SecurityRegistryKeyEvidenceRoles_Contextual,
		"created":           SecurityRegistryKeyEvidenceRoles_Created,
		"destination":       SecurityRegistryKeyEvidenceRoles_Destination,
		"edited":            SecurityRegistryKeyEvidenceRoles_Edited,
		"loaded":            SecurityRegistryKeyEvidenceRoles_Loaded,
		"policyviolator":    SecurityRegistryKeyEvidenceRoles_PolicyViolator,
		"scanned":           SecurityRegistryKeyEvidenceRoles_Scanned,
		"source":            SecurityRegistryKeyEvidenceRoles_Source,
		"suspicious":        SecurityRegistryKeyEvidenceRoles_Suspicious,
		"unknown":           SecurityRegistryKeyEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRegistryKeyEvidenceRoles(input)
	return &out, nil
}
