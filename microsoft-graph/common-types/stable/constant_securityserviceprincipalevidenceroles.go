package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityServicePrincipalEvidenceRoles string

const (
	SecurityServicePrincipalEvidenceRoles_Added             SecurityServicePrincipalEvidenceRoles = "added"
	SecurityServicePrincipalEvidenceRoles_Attacked          SecurityServicePrincipalEvidenceRoles = "attacked"
	SecurityServicePrincipalEvidenceRoles_Attacker          SecurityServicePrincipalEvidenceRoles = "attacker"
	SecurityServicePrincipalEvidenceRoles_CommandAndControl SecurityServicePrincipalEvidenceRoles = "commandAndControl"
	SecurityServicePrincipalEvidenceRoles_Compromised       SecurityServicePrincipalEvidenceRoles = "compromised"
	SecurityServicePrincipalEvidenceRoles_Contextual        SecurityServicePrincipalEvidenceRoles = "contextual"
	SecurityServicePrincipalEvidenceRoles_Created           SecurityServicePrincipalEvidenceRoles = "created"
	SecurityServicePrincipalEvidenceRoles_Destination       SecurityServicePrincipalEvidenceRoles = "destination"
	SecurityServicePrincipalEvidenceRoles_Edited            SecurityServicePrincipalEvidenceRoles = "edited"
	SecurityServicePrincipalEvidenceRoles_Loaded            SecurityServicePrincipalEvidenceRoles = "loaded"
	SecurityServicePrincipalEvidenceRoles_PolicyViolator    SecurityServicePrincipalEvidenceRoles = "policyViolator"
	SecurityServicePrincipalEvidenceRoles_Scanned           SecurityServicePrincipalEvidenceRoles = "scanned"
	SecurityServicePrincipalEvidenceRoles_Source            SecurityServicePrincipalEvidenceRoles = "source"
	SecurityServicePrincipalEvidenceRoles_Suspicious        SecurityServicePrincipalEvidenceRoles = "suspicious"
	SecurityServicePrincipalEvidenceRoles_Unknown           SecurityServicePrincipalEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityServicePrincipalEvidenceRoles() []string {
	return []string{
		string(SecurityServicePrincipalEvidenceRoles_Added),
		string(SecurityServicePrincipalEvidenceRoles_Attacked),
		string(SecurityServicePrincipalEvidenceRoles_Attacker),
		string(SecurityServicePrincipalEvidenceRoles_CommandAndControl),
		string(SecurityServicePrincipalEvidenceRoles_Compromised),
		string(SecurityServicePrincipalEvidenceRoles_Contextual),
		string(SecurityServicePrincipalEvidenceRoles_Created),
		string(SecurityServicePrincipalEvidenceRoles_Destination),
		string(SecurityServicePrincipalEvidenceRoles_Edited),
		string(SecurityServicePrincipalEvidenceRoles_Loaded),
		string(SecurityServicePrincipalEvidenceRoles_PolicyViolator),
		string(SecurityServicePrincipalEvidenceRoles_Scanned),
		string(SecurityServicePrincipalEvidenceRoles_Source),
		string(SecurityServicePrincipalEvidenceRoles_Suspicious),
		string(SecurityServicePrincipalEvidenceRoles_Unknown),
	}
}

func (s *SecurityServicePrincipalEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityServicePrincipalEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityServicePrincipalEvidenceRoles(input string) (*SecurityServicePrincipalEvidenceRoles, error) {
	vals := map[string]SecurityServicePrincipalEvidenceRoles{
		"added":             SecurityServicePrincipalEvidenceRoles_Added,
		"attacked":          SecurityServicePrincipalEvidenceRoles_Attacked,
		"attacker":          SecurityServicePrincipalEvidenceRoles_Attacker,
		"commandandcontrol": SecurityServicePrincipalEvidenceRoles_CommandAndControl,
		"compromised":       SecurityServicePrincipalEvidenceRoles_Compromised,
		"contextual":        SecurityServicePrincipalEvidenceRoles_Contextual,
		"created":           SecurityServicePrincipalEvidenceRoles_Created,
		"destination":       SecurityServicePrincipalEvidenceRoles_Destination,
		"edited":            SecurityServicePrincipalEvidenceRoles_Edited,
		"loaded":            SecurityServicePrincipalEvidenceRoles_Loaded,
		"policyviolator":    SecurityServicePrincipalEvidenceRoles_PolicyViolator,
		"scanned":           SecurityServicePrincipalEvidenceRoles_Scanned,
		"source":            SecurityServicePrincipalEvidenceRoles_Source,
		"suspicious":        SecurityServicePrincipalEvidenceRoles_Suspicious,
		"unknown":           SecurityServicePrincipalEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityServicePrincipalEvidenceRoles(input)
	return &out, nil
}
