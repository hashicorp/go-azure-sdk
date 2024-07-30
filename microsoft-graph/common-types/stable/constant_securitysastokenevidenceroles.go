package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySasTokenEvidenceRoles string

const (
	SecuritySasTokenEvidenceRoles_Added             SecuritySasTokenEvidenceRoles = "added"
	SecuritySasTokenEvidenceRoles_Attacked          SecuritySasTokenEvidenceRoles = "attacked"
	SecuritySasTokenEvidenceRoles_Attacker          SecuritySasTokenEvidenceRoles = "attacker"
	SecuritySasTokenEvidenceRoles_CommandAndControl SecuritySasTokenEvidenceRoles = "commandAndControl"
	SecuritySasTokenEvidenceRoles_Compromised       SecuritySasTokenEvidenceRoles = "compromised"
	SecuritySasTokenEvidenceRoles_Contextual        SecuritySasTokenEvidenceRoles = "contextual"
	SecuritySasTokenEvidenceRoles_Created           SecuritySasTokenEvidenceRoles = "created"
	SecuritySasTokenEvidenceRoles_Destination       SecuritySasTokenEvidenceRoles = "destination"
	SecuritySasTokenEvidenceRoles_Edited            SecuritySasTokenEvidenceRoles = "edited"
	SecuritySasTokenEvidenceRoles_Loaded            SecuritySasTokenEvidenceRoles = "loaded"
	SecuritySasTokenEvidenceRoles_PolicyViolator    SecuritySasTokenEvidenceRoles = "policyViolator"
	SecuritySasTokenEvidenceRoles_Scanned           SecuritySasTokenEvidenceRoles = "scanned"
	SecuritySasTokenEvidenceRoles_Source            SecuritySasTokenEvidenceRoles = "source"
	SecuritySasTokenEvidenceRoles_Suspicious        SecuritySasTokenEvidenceRoles = "suspicious"
	SecuritySasTokenEvidenceRoles_Unknown           SecuritySasTokenEvidenceRoles = "unknown"
)

func PossibleValuesForSecuritySasTokenEvidenceRoles() []string {
	return []string{
		string(SecuritySasTokenEvidenceRoles_Added),
		string(SecuritySasTokenEvidenceRoles_Attacked),
		string(SecuritySasTokenEvidenceRoles_Attacker),
		string(SecuritySasTokenEvidenceRoles_CommandAndControl),
		string(SecuritySasTokenEvidenceRoles_Compromised),
		string(SecuritySasTokenEvidenceRoles_Contextual),
		string(SecuritySasTokenEvidenceRoles_Created),
		string(SecuritySasTokenEvidenceRoles_Destination),
		string(SecuritySasTokenEvidenceRoles_Edited),
		string(SecuritySasTokenEvidenceRoles_Loaded),
		string(SecuritySasTokenEvidenceRoles_PolicyViolator),
		string(SecuritySasTokenEvidenceRoles_Scanned),
		string(SecuritySasTokenEvidenceRoles_Source),
		string(SecuritySasTokenEvidenceRoles_Suspicious),
		string(SecuritySasTokenEvidenceRoles_Unknown),
	}
}

func (s *SecuritySasTokenEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySasTokenEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySasTokenEvidenceRoles(input string) (*SecuritySasTokenEvidenceRoles, error) {
	vals := map[string]SecuritySasTokenEvidenceRoles{
		"added":             SecuritySasTokenEvidenceRoles_Added,
		"attacked":          SecuritySasTokenEvidenceRoles_Attacked,
		"attacker":          SecuritySasTokenEvidenceRoles_Attacker,
		"commandandcontrol": SecuritySasTokenEvidenceRoles_CommandAndControl,
		"compromised":       SecuritySasTokenEvidenceRoles_Compromised,
		"contextual":        SecuritySasTokenEvidenceRoles_Contextual,
		"created":           SecuritySasTokenEvidenceRoles_Created,
		"destination":       SecuritySasTokenEvidenceRoles_Destination,
		"edited":            SecuritySasTokenEvidenceRoles_Edited,
		"loaded":            SecuritySasTokenEvidenceRoles_Loaded,
		"policyviolator":    SecuritySasTokenEvidenceRoles_PolicyViolator,
		"scanned":           SecuritySasTokenEvidenceRoles_Scanned,
		"source":            SecuritySasTokenEvidenceRoles_Source,
		"suspicious":        SecuritySasTokenEvidenceRoles_Suspicious,
		"unknown":           SecuritySasTokenEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySasTokenEvidenceRoles(input)
	return &out, nil
}
