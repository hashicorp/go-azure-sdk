package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRegistryValueEvidenceRoles string

const (
	SecurityRegistryValueEvidenceRoles_Added             SecurityRegistryValueEvidenceRoles = "added"
	SecurityRegistryValueEvidenceRoles_Attacked          SecurityRegistryValueEvidenceRoles = "attacked"
	SecurityRegistryValueEvidenceRoles_Attacker          SecurityRegistryValueEvidenceRoles = "attacker"
	SecurityRegistryValueEvidenceRoles_CommandAndControl SecurityRegistryValueEvidenceRoles = "commandAndControl"
	SecurityRegistryValueEvidenceRoles_Compromised       SecurityRegistryValueEvidenceRoles = "compromised"
	SecurityRegistryValueEvidenceRoles_Contextual        SecurityRegistryValueEvidenceRoles = "contextual"
	SecurityRegistryValueEvidenceRoles_Created           SecurityRegistryValueEvidenceRoles = "created"
	SecurityRegistryValueEvidenceRoles_Destination       SecurityRegistryValueEvidenceRoles = "destination"
	SecurityRegistryValueEvidenceRoles_Edited            SecurityRegistryValueEvidenceRoles = "edited"
	SecurityRegistryValueEvidenceRoles_Loaded            SecurityRegistryValueEvidenceRoles = "loaded"
	SecurityRegistryValueEvidenceRoles_PolicyViolator    SecurityRegistryValueEvidenceRoles = "policyViolator"
	SecurityRegistryValueEvidenceRoles_Scanned           SecurityRegistryValueEvidenceRoles = "scanned"
	SecurityRegistryValueEvidenceRoles_Source            SecurityRegistryValueEvidenceRoles = "source"
	SecurityRegistryValueEvidenceRoles_Suspicious        SecurityRegistryValueEvidenceRoles = "suspicious"
	SecurityRegistryValueEvidenceRoles_Unknown           SecurityRegistryValueEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityRegistryValueEvidenceRoles() []string {
	return []string{
		string(SecurityRegistryValueEvidenceRoles_Added),
		string(SecurityRegistryValueEvidenceRoles_Attacked),
		string(SecurityRegistryValueEvidenceRoles_Attacker),
		string(SecurityRegistryValueEvidenceRoles_CommandAndControl),
		string(SecurityRegistryValueEvidenceRoles_Compromised),
		string(SecurityRegistryValueEvidenceRoles_Contextual),
		string(SecurityRegistryValueEvidenceRoles_Created),
		string(SecurityRegistryValueEvidenceRoles_Destination),
		string(SecurityRegistryValueEvidenceRoles_Edited),
		string(SecurityRegistryValueEvidenceRoles_Loaded),
		string(SecurityRegistryValueEvidenceRoles_PolicyViolator),
		string(SecurityRegistryValueEvidenceRoles_Scanned),
		string(SecurityRegistryValueEvidenceRoles_Source),
		string(SecurityRegistryValueEvidenceRoles_Suspicious),
		string(SecurityRegistryValueEvidenceRoles_Unknown),
	}
}

func (s *SecurityRegistryValueEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRegistryValueEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRegistryValueEvidenceRoles(input string) (*SecurityRegistryValueEvidenceRoles, error) {
	vals := map[string]SecurityRegistryValueEvidenceRoles{
		"added":             SecurityRegistryValueEvidenceRoles_Added,
		"attacked":          SecurityRegistryValueEvidenceRoles_Attacked,
		"attacker":          SecurityRegistryValueEvidenceRoles_Attacker,
		"commandandcontrol": SecurityRegistryValueEvidenceRoles_CommandAndControl,
		"compromised":       SecurityRegistryValueEvidenceRoles_Compromised,
		"contextual":        SecurityRegistryValueEvidenceRoles_Contextual,
		"created":           SecurityRegistryValueEvidenceRoles_Created,
		"destination":       SecurityRegistryValueEvidenceRoles_Destination,
		"edited":            SecurityRegistryValueEvidenceRoles_Edited,
		"loaded":            SecurityRegistryValueEvidenceRoles_Loaded,
		"policyviolator":    SecurityRegistryValueEvidenceRoles_PolicyViolator,
		"scanned":           SecurityRegistryValueEvidenceRoles_Scanned,
		"source":            SecurityRegistryValueEvidenceRoles_Source,
		"suspicious":        SecurityRegistryValueEvidenceRoles_Suspicious,
		"unknown":           SecurityRegistryValueEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRegistryValueEvidenceRoles(input)
	return &out, nil
}
