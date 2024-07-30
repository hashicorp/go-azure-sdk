package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesSecretEvidenceRoles string

const (
	SecurityKubernetesSecretEvidenceRoles_Added             SecurityKubernetesSecretEvidenceRoles = "added"
	SecurityKubernetesSecretEvidenceRoles_Attacked          SecurityKubernetesSecretEvidenceRoles = "attacked"
	SecurityKubernetesSecretEvidenceRoles_Attacker          SecurityKubernetesSecretEvidenceRoles = "attacker"
	SecurityKubernetesSecretEvidenceRoles_CommandAndControl SecurityKubernetesSecretEvidenceRoles = "commandAndControl"
	SecurityKubernetesSecretEvidenceRoles_Compromised       SecurityKubernetesSecretEvidenceRoles = "compromised"
	SecurityKubernetesSecretEvidenceRoles_Contextual        SecurityKubernetesSecretEvidenceRoles = "contextual"
	SecurityKubernetesSecretEvidenceRoles_Created           SecurityKubernetesSecretEvidenceRoles = "created"
	SecurityKubernetesSecretEvidenceRoles_Destination       SecurityKubernetesSecretEvidenceRoles = "destination"
	SecurityKubernetesSecretEvidenceRoles_Edited            SecurityKubernetesSecretEvidenceRoles = "edited"
	SecurityKubernetesSecretEvidenceRoles_Loaded            SecurityKubernetesSecretEvidenceRoles = "loaded"
	SecurityKubernetesSecretEvidenceRoles_PolicyViolator    SecurityKubernetesSecretEvidenceRoles = "policyViolator"
	SecurityKubernetesSecretEvidenceRoles_Scanned           SecurityKubernetesSecretEvidenceRoles = "scanned"
	SecurityKubernetesSecretEvidenceRoles_Source            SecurityKubernetesSecretEvidenceRoles = "source"
	SecurityKubernetesSecretEvidenceRoles_Suspicious        SecurityKubernetesSecretEvidenceRoles = "suspicious"
	SecurityKubernetesSecretEvidenceRoles_Unknown           SecurityKubernetesSecretEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityKubernetesSecretEvidenceRoles() []string {
	return []string{
		string(SecurityKubernetesSecretEvidenceRoles_Added),
		string(SecurityKubernetesSecretEvidenceRoles_Attacked),
		string(SecurityKubernetesSecretEvidenceRoles_Attacker),
		string(SecurityKubernetesSecretEvidenceRoles_CommandAndControl),
		string(SecurityKubernetesSecretEvidenceRoles_Compromised),
		string(SecurityKubernetesSecretEvidenceRoles_Contextual),
		string(SecurityKubernetesSecretEvidenceRoles_Created),
		string(SecurityKubernetesSecretEvidenceRoles_Destination),
		string(SecurityKubernetesSecretEvidenceRoles_Edited),
		string(SecurityKubernetesSecretEvidenceRoles_Loaded),
		string(SecurityKubernetesSecretEvidenceRoles_PolicyViolator),
		string(SecurityKubernetesSecretEvidenceRoles_Scanned),
		string(SecurityKubernetesSecretEvidenceRoles_Source),
		string(SecurityKubernetesSecretEvidenceRoles_Suspicious),
		string(SecurityKubernetesSecretEvidenceRoles_Unknown),
	}
}

func (s *SecurityKubernetesSecretEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesSecretEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesSecretEvidenceRoles(input string) (*SecurityKubernetesSecretEvidenceRoles, error) {
	vals := map[string]SecurityKubernetesSecretEvidenceRoles{
		"added":             SecurityKubernetesSecretEvidenceRoles_Added,
		"attacked":          SecurityKubernetesSecretEvidenceRoles_Attacked,
		"attacker":          SecurityKubernetesSecretEvidenceRoles_Attacker,
		"commandandcontrol": SecurityKubernetesSecretEvidenceRoles_CommandAndControl,
		"compromised":       SecurityKubernetesSecretEvidenceRoles_Compromised,
		"contextual":        SecurityKubernetesSecretEvidenceRoles_Contextual,
		"created":           SecurityKubernetesSecretEvidenceRoles_Created,
		"destination":       SecurityKubernetesSecretEvidenceRoles_Destination,
		"edited":            SecurityKubernetesSecretEvidenceRoles_Edited,
		"loaded":            SecurityKubernetesSecretEvidenceRoles_Loaded,
		"policyviolator":    SecurityKubernetesSecretEvidenceRoles_PolicyViolator,
		"scanned":           SecurityKubernetesSecretEvidenceRoles_Scanned,
		"source":            SecurityKubernetesSecretEvidenceRoles_Source,
		"suspicious":        SecurityKubernetesSecretEvidenceRoles_Suspicious,
		"unknown":           SecurityKubernetesSecretEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesSecretEvidenceRoles(input)
	return &out, nil
}
