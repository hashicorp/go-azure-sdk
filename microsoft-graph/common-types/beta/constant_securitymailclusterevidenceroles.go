package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailClusterEvidenceRoles string

const (
	SecurityMailClusterEvidenceRoles_Added             SecurityMailClusterEvidenceRoles = "added"
	SecurityMailClusterEvidenceRoles_Attacked          SecurityMailClusterEvidenceRoles = "attacked"
	SecurityMailClusterEvidenceRoles_Attacker          SecurityMailClusterEvidenceRoles = "attacker"
	SecurityMailClusterEvidenceRoles_CommandAndControl SecurityMailClusterEvidenceRoles = "commandAndControl"
	SecurityMailClusterEvidenceRoles_Compromised       SecurityMailClusterEvidenceRoles = "compromised"
	SecurityMailClusterEvidenceRoles_Contextual        SecurityMailClusterEvidenceRoles = "contextual"
	SecurityMailClusterEvidenceRoles_Created           SecurityMailClusterEvidenceRoles = "created"
	SecurityMailClusterEvidenceRoles_Destination       SecurityMailClusterEvidenceRoles = "destination"
	SecurityMailClusterEvidenceRoles_Edited            SecurityMailClusterEvidenceRoles = "edited"
	SecurityMailClusterEvidenceRoles_Loaded            SecurityMailClusterEvidenceRoles = "loaded"
	SecurityMailClusterEvidenceRoles_PolicyViolator    SecurityMailClusterEvidenceRoles = "policyViolator"
	SecurityMailClusterEvidenceRoles_Scanned           SecurityMailClusterEvidenceRoles = "scanned"
	SecurityMailClusterEvidenceRoles_Source            SecurityMailClusterEvidenceRoles = "source"
	SecurityMailClusterEvidenceRoles_Suspicious        SecurityMailClusterEvidenceRoles = "suspicious"
	SecurityMailClusterEvidenceRoles_Unknown           SecurityMailClusterEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityMailClusterEvidenceRoles() []string {
	return []string{
		string(SecurityMailClusterEvidenceRoles_Added),
		string(SecurityMailClusterEvidenceRoles_Attacked),
		string(SecurityMailClusterEvidenceRoles_Attacker),
		string(SecurityMailClusterEvidenceRoles_CommandAndControl),
		string(SecurityMailClusterEvidenceRoles_Compromised),
		string(SecurityMailClusterEvidenceRoles_Contextual),
		string(SecurityMailClusterEvidenceRoles_Created),
		string(SecurityMailClusterEvidenceRoles_Destination),
		string(SecurityMailClusterEvidenceRoles_Edited),
		string(SecurityMailClusterEvidenceRoles_Loaded),
		string(SecurityMailClusterEvidenceRoles_PolicyViolator),
		string(SecurityMailClusterEvidenceRoles_Scanned),
		string(SecurityMailClusterEvidenceRoles_Source),
		string(SecurityMailClusterEvidenceRoles_Suspicious),
		string(SecurityMailClusterEvidenceRoles_Unknown),
	}
}

func (s *SecurityMailClusterEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityMailClusterEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityMailClusterEvidenceRoles(input string) (*SecurityMailClusterEvidenceRoles, error) {
	vals := map[string]SecurityMailClusterEvidenceRoles{
		"added":             SecurityMailClusterEvidenceRoles_Added,
		"attacked":          SecurityMailClusterEvidenceRoles_Attacked,
		"attacker":          SecurityMailClusterEvidenceRoles_Attacker,
		"commandandcontrol": SecurityMailClusterEvidenceRoles_CommandAndControl,
		"compromised":       SecurityMailClusterEvidenceRoles_Compromised,
		"contextual":        SecurityMailClusterEvidenceRoles_Contextual,
		"created":           SecurityMailClusterEvidenceRoles_Created,
		"destination":       SecurityMailClusterEvidenceRoles_Destination,
		"edited":            SecurityMailClusterEvidenceRoles_Edited,
		"loaded":            SecurityMailClusterEvidenceRoles_Loaded,
		"policyviolator":    SecurityMailClusterEvidenceRoles_PolicyViolator,
		"scanned":           SecurityMailClusterEvidenceRoles_Scanned,
		"source":            SecurityMailClusterEvidenceRoles_Source,
		"suspicious":        SecurityMailClusterEvidenceRoles_Suspicious,
		"unknown":           SecurityMailClusterEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMailClusterEvidenceRoles(input)
	return &out, nil
}
