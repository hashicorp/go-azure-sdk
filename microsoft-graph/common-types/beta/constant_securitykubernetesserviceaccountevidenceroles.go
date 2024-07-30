package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceAccountEvidenceRoles string

const (
	SecurityKubernetesServiceAccountEvidenceRoles_Added             SecurityKubernetesServiceAccountEvidenceRoles = "added"
	SecurityKubernetesServiceAccountEvidenceRoles_Attacked          SecurityKubernetesServiceAccountEvidenceRoles = "attacked"
	SecurityKubernetesServiceAccountEvidenceRoles_Attacker          SecurityKubernetesServiceAccountEvidenceRoles = "attacker"
	SecurityKubernetesServiceAccountEvidenceRoles_CommandAndControl SecurityKubernetesServiceAccountEvidenceRoles = "commandAndControl"
	SecurityKubernetesServiceAccountEvidenceRoles_Compromised       SecurityKubernetesServiceAccountEvidenceRoles = "compromised"
	SecurityKubernetesServiceAccountEvidenceRoles_Contextual        SecurityKubernetesServiceAccountEvidenceRoles = "contextual"
	SecurityKubernetesServiceAccountEvidenceRoles_Created           SecurityKubernetesServiceAccountEvidenceRoles = "created"
	SecurityKubernetesServiceAccountEvidenceRoles_Destination       SecurityKubernetesServiceAccountEvidenceRoles = "destination"
	SecurityKubernetesServiceAccountEvidenceRoles_Edited            SecurityKubernetesServiceAccountEvidenceRoles = "edited"
	SecurityKubernetesServiceAccountEvidenceRoles_Loaded            SecurityKubernetesServiceAccountEvidenceRoles = "loaded"
	SecurityKubernetesServiceAccountEvidenceRoles_PolicyViolator    SecurityKubernetesServiceAccountEvidenceRoles = "policyViolator"
	SecurityKubernetesServiceAccountEvidenceRoles_Scanned           SecurityKubernetesServiceAccountEvidenceRoles = "scanned"
	SecurityKubernetesServiceAccountEvidenceRoles_Source            SecurityKubernetesServiceAccountEvidenceRoles = "source"
	SecurityKubernetesServiceAccountEvidenceRoles_Suspicious        SecurityKubernetesServiceAccountEvidenceRoles = "suspicious"
	SecurityKubernetesServiceAccountEvidenceRoles_Unknown           SecurityKubernetesServiceAccountEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityKubernetesServiceAccountEvidenceRoles() []string {
	return []string{
		string(SecurityKubernetesServiceAccountEvidenceRoles_Added),
		string(SecurityKubernetesServiceAccountEvidenceRoles_Attacked),
		string(SecurityKubernetesServiceAccountEvidenceRoles_Attacker),
		string(SecurityKubernetesServiceAccountEvidenceRoles_CommandAndControl),
		string(SecurityKubernetesServiceAccountEvidenceRoles_Compromised),
		string(SecurityKubernetesServiceAccountEvidenceRoles_Contextual),
		string(SecurityKubernetesServiceAccountEvidenceRoles_Created),
		string(SecurityKubernetesServiceAccountEvidenceRoles_Destination),
		string(SecurityKubernetesServiceAccountEvidenceRoles_Edited),
		string(SecurityKubernetesServiceAccountEvidenceRoles_Loaded),
		string(SecurityKubernetesServiceAccountEvidenceRoles_PolicyViolator),
		string(SecurityKubernetesServiceAccountEvidenceRoles_Scanned),
		string(SecurityKubernetesServiceAccountEvidenceRoles_Source),
		string(SecurityKubernetesServiceAccountEvidenceRoles_Suspicious),
		string(SecurityKubernetesServiceAccountEvidenceRoles_Unknown),
	}
}

func (s *SecurityKubernetesServiceAccountEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesServiceAccountEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesServiceAccountEvidenceRoles(input string) (*SecurityKubernetesServiceAccountEvidenceRoles, error) {
	vals := map[string]SecurityKubernetesServiceAccountEvidenceRoles{
		"added":             SecurityKubernetesServiceAccountEvidenceRoles_Added,
		"attacked":          SecurityKubernetesServiceAccountEvidenceRoles_Attacked,
		"attacker":          SecurityKubernetesServiceAccountEvidenceRoles_Attacker,
		"commandandcontrol": SecurityKubernetesServiceAccountEvidenceRoles_CommandAndControl,
		"compromised":       SecurityKubernetesServiceAccountEvidenceRoles_Compromised,
		"contextual":        SecurityKubernetesServiceAccountEvidenceRoles_Contextual,
		"created":           SecurityKubernetesServiceAccountEvidenceRoles_Created,
		"destination":       SecurityKubernetesServiceAccountEvidenceRoles_Destination,
		"edited":            SecurityKubernetesServiceAccountEvidenceRoles_Edited,
		"loaded":            SecurityKubernetesServiceAccountEvidenceRoles_Loaded,
		"policyviolator":    SecurityKubernetesServiceAccountEvidenceRoles_PolicyViolator,
		"scanned":           SecurityKubernetesServiceAccountEvidenceRoles_Scanned,
		"source":            SecurityKubernetesServiceAccountEvidenceRoles_Source,
		"suspicious":        SecurityKubernetesServiceAccountEvidenceRoles_Suspicious,
		"unknown":           SecurityKubernetesServiceAccountEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServiceAccountEvidenceRoles(input)
	return &out, nil
}
