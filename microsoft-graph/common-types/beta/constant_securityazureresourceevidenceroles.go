package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAzureResourceEvidenceRoles string

const (
	SecurityAzureResourceEvidenceRoles_Added             SecurityAzureResourceEvidenceRoles = "added"
	SecurityAzureResourceEvidenceRoles_Attacked          SecurityAzureResourceEvidenceRoles = "attacked"
	SecurityAzureResourceEvidenceRoles_Attacker          SecurityAzureResourceEvidenceRoles = "attacker"
	SecurityAzureResourceEvidenceRoles_CommandAndControl SecurityAzureResourceEvidenceRoles = "commandAndControl"
	SecurityAzureResourceEvidenceRoles_Compromised       SecurityAzureResourceEvidenceRoles = "compromised"
	SecurityAzureResourceEvidenceRoles_Contextual        SecurityAzureResourceEvidenceRoles = "contextual"
	SecurityAzureResourceEvidenceRoles_Created           SecurityAzureResourceEvidenceRoles = "created"
	SecurityAzureResourceEvidenceRoles_Destination       SecurityAzureResourceEvidenceRoles = "destination"
	SecurityAzureResourceEvidenceRoles_Edited            SecurityAzureResourceEvidenceRoles = "edited"
	SecurityAzureResourceEvidenceRoles_Loaded            SecurityAzureResourceEvidenceRoles = "loaded"
	SecurityAzureResourceEvidenceRoles_PolicyViolator    SecurityAzureResourceEvidenceRoles = "policyViolator"
	SecurityAzureResourceEvidenceRoles_Scanned           SecurityAzureResourceEvidenceRoles = "scanned"
	SecurityAzureResourceEvidenceRoles_Source            SecurityAzureResourceEvidenceRoles = "source"
	SecurityAzureResourceEvidenceRoles_Suspicious        SecurityAzureResourceEvidenceRoles = "suspicious"
	SecurityAzureResourceEvidenceRoles_Unknown           SecurityAzureResourceEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityAzureResourceEvidenceRoles() []string {
	return []string{
		string(SecurityAzureResourceEvidenceRoles_Added),
		string(SecurityAzureResourceEvidenceRoles_Attacked),
		string(SecurityAzureResourceEvidenceRoles_Attacker),
		string(SecurityAzureResourceEvidenceRoles_CommandAndControl),
		string(SecurityAzureResourceEvidenceRoles_Compromised),
		string(SecurityAzureResourceEvidenceRoles_Contextual),
		string(SecurityAzureResourceEvidenceRoles_Created),
		string(SecurityAzureResourceEvidenceRoles_Destination),
		string(SecurityAzureResourceEvidenceRoles_Edited),
		string(SecurityAzureResourceEvidenceRoles_Loaded),
		string(SecurityAzureResourceEvidenceRoles_PolicyViolator),
		string(SecurityAzureResourceEvidenceRoles_Scanned),
		string(SecurityAzureResourceEvidenceRoles_Source),
		string(SecurityAzureResourceEvidenceRoles_Suspicious),
		string(SecurityAzureResourceEvidenceRoles_Unknown),
	}
}

func (s *SecurityAzureResourceEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAzureResourceEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAzureResourceEvidenceRoles(input string) (*SecurityAzureResourceEvidenceRoles, error) {
	vals := map[string]SecurityAzureResourceEvidenceRoles{
		"added":             SecurityAzureResourceEvidenceRoles_Added,
		"attacked":          SecurityAzureResourceEvidenceRoles_Attacked,
		"attacker":          SecurityAzureResourceEvidenceRoles_Attacker,
		"commandandcontrol": SecurityAzureResourceEvidenceRoles_CommandAndControl,
		"compromised":       SecurityAzureResourceEvidenceRoles_Compromised,
		"contextual":        SecurityAzureResourceEvidenceRoles_Contextual,
		"created":           SecurityAzureResourceEvidenceRoles_Created,
		"destination":       SecurityAzureResourceEvidenceRoles_Destination,
		"edited":            SecurityAzureResourceEvidenceRoles_Edited,
		"loaded":            SecurityAzureResourceEvidenceRoles_Loaded,
		"policyviolator":    SecurityAzureResourceEvidenceRoles_PolicyViolator,
		"scanned":           SecurityAzureResourceEvidenceRoles_Scanned,
		"source":            SecurityAzureResourceEvidenceRoles_Source,
		"suspicious":        SecurityAzureResourceEvidenceRoles_Suspicious,
		"unknown":           SecurityAzureResourceEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAzureResourceEvidenceRoles(input)
	return &out, nil
}
