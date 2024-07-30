package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceEvidenceRoles string

const (
	SecurityKubernetesServiceEvidenceRoles_Added             SecurityKubernetesServiceEvidenceRoles = "added"
	SecurityKubernetesServiceEvidenceRoles_Attacked          SecurityKubernetesServiceEvidenceRoles = "attacked"
	SecurityKubernetesServiceEvidenceRoles_Attacker          SecurityKubernetesServiceEvidenceRoles = "attacker"
	SecurityKubernetesServiceEvidenceRoles_CommandAndControl SecurityKubernetesServiceEvidenceRoles = "commandAndControl"
	SecurityKubernetesServiceEvidenceRoles_Compromised       SecurityKubernetesServiceEvidenceRoles = "compromised"
	SecurityKubernetesServiceEvidenceRoles_Contextual        SecurityKubernetesServiceEvidenceRoles = "contextual"
	SecurityKubernetesServiceEvidenceRoles_Created           SecurityKubernetesServiceEvidenceRoles = "created"
	SecurityKubernetesServiceEvidenceRoles_Destination       SecurityKubernetesServiceEvidenceRoles = "destination"
	SecurityKubernetesServiceEvidenceRoles_Edited            SecurityKubernetesServiceEvidenceRoles = "edited"
	SecurityKubernetesServiceEvidenceRoles_Loaded            SecurityKubernetesServiceEvidenceRoles = "loaded"
	SecurityKubernetesServiceEvidenceRoles_PolicyViolator    SecurityKubernetesServiceEvidenceRoles = "policyViolator"
	SecurityKubernetesServiceEvidenceRoles_Scanned           SecurityKubernetesServiceEvidenceRoles = "scanned"
	SecurityKubernetesServiceEvidenceRoles_Source            SecurityKubernetesServiceEvidenceRoles = "source"
	SecurityKubernetesServiceEvidenceRoles_Suspicious        SecurityKubernetesServiceEvidenceRoles = "suspicious"
	SecurityKubernetesServiceEvidenceRoles_Unknown           SecurityKubernetesServiceEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityKubernetesServiceEvidenceRoles() []string {
	return []string{
		string(SecurityKubernetesServiceEvidenceRoles_Added),
		string(SecurityKubernetesServiceEvidenceRoles_Attacked),
		string(SecurityKubernetesServiceEvidenceRoles_Attacker),
		string(SecurityKubernetesServiceEvidenceRoles_CommandAndControl),
		string(SecurityKubernetesServiceEvidenceRoles_Compromised),
		string(SecurityKubernetesServiceEvidenceRoles_Contextual),
		string(SecurityKubernetesServiceEvidenceRoles_Created),
		string(SecurityKubernetesServiceEvidenceRoles_Destination),
		string(SecurityKubernetesServiceEvidenceRoles_Edited),
		string(SecurityKubernetesServiceEvidenceRoles_Loaded),
		string(SecurityKubernetesServiceEvidenceRoles_PolicyViolator),
		string(SecurityKubernetesServiceEvidenceRoles_Scanned),
		string(SecurityKubernetesServiceEvidenceRoles_Source),
		string(SecurityKubernetesServiceEvidenceRoles_Suspicious),
		string(SecurityKubernetesServiceEvidenceRoles_Unknown),
	}
}

func (s *SecurityKubernetesServiceEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesServiceEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesServiceEvidenceRoles(input string) (*SecurityKubernetesServiceEvidenceRoles, error) {
	vals := map[string]SecurityKubernetesServiceEvidenceRoles{
		"added":             SecurityKubernetesServiceEvidenceRoles_Added,
		"attacked":          SecurityKubernetesServiceEvidenceRoles_Attacked,
		"attacker":          SecurityKubernetesServiceEvidenceRoles_Attacker,
		"commandandcontrol": SecurityKubernetesServiceEvidenceRoles_CommandAndControl,
		"compromised":       SecurityKubernetesServiceEvidenceRoles_Compromised,
		"contextual":        SecurityKubernetesServiceEvidenceRoles_Contextual,
		"created":           SecurityKubernetesServiceEvidenceRoles_Created,
		"destination":       SecurityKubernetesServiceEvidenceRoles_Destination,
		"edited":            SecurityKubernetesServiceEvidenceRoles_Edited,
		"loaded":            SecurityKubernetesServiceEvidenceRoles_Loaded,
		"policyviolator":    SecurityKubernetesServiceEvidenceRoles_PolicyViolator,
		"scanned":           SecurityKubernetesServiceEvidenceRoles_Scanned,
		"source":            SecurityKubernetesServiceEvidenceRoles_Source,
		"suspicious":        SecurityKubernetesServiceEvidenceRoles_Suspicious,
		"unknown":           SecurityKubernetesServiceEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServiceEvidenceRoles(input)
	return &out, nil
}
