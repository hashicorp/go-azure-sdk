package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAmazonResourceEvidenceRoles string

const (
	SecurityAmazonResourceEvidenceRoles_Added             SecurityAmazonResourceEvidenceRoles = "added"
	SecurityAmazonResourceEvidenceRoles_Attacked          SecurityAmazonResourceEvidenceRoles = "attacked"
	SecurityAmazonResourceEvidenceRoles_Attacker          SecurityAmazonResourceEvidenceRoles = "attacker"
	SecurityAmazonResourceEvidenceRoles_CommandAndControl SecurityAmazonResourceEvidenceRoles = "commandAndControl"
	SecurityAmazonResourceEvidenceRoles_Compromised       SecurityAmazonResourceEvidenceRoles = "compromised"
	SecurityAmazonResourceEvidenceRoles_Contextual        SecurityAmazonResourceEvidenceRoles = "contextual"
	SecurityAmazonResourceEvidenceRoles_Created           SecurityAmazonResourceEvidenceRoles = "created"
	SecurityAmazonResourceEvidenceRoles_Destination       SecurityAmazonResourceEvidenceRoles = "destination"
	SecurityAmazonResourceEvidenceRoles_Edited            SecurityAmazonResourceEvidenceRoles = "edited"
	SecurityAmazonResourceEvidenceRoles_Loaded            SecurityAmazonResourceEvidenceRoles = "loaded"
	SecurityAmazonResourceEvidenceRoles_PolicyViolator    SecurityAmazonResourceEvidenceRoles = "policyViolator"
	SecurityAmazonResourceEvidenceRoles_Scanned           SecurityAmazonResourceEvidenceRoles = "scanned"
	SecurityAmazonResourceEvidenceRoles_Source            SecurityAmazonResourceEvidenceRoles = "source"
	SecurityAmazonResourceEvidenceRoles_Suspicious        SecurityAmazonResourceEvidenceRoles = "suspicious"
	SecurityAmazonResourceEvidenceRoles_Unknown           SecurityAmazonResourceEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityAmazonResourceEvidenceRoles() []string {
	return []string{
		string(SecurityAmazonResourceEvidenceRoles_Added),
		string(SecurityAmazonResourceEvidenceRoles_Attacked),
		string(SecurityAmazonResourceEvidenceRoles_Attacker),
		string(SecurityAmazonResourceEvidenceRoles_CommandAndControl),
		string(SecurityAmazonResourceEvidenceRoles_Compromised),
		string(SecurityAmazonResourceEvidenceRoles_Contextual),
		string(SecurityAmazonResourceEvidenceRoles_Created),
		string(SecurityAmazonResourceEvidenceRoles_Destination),
		string(SecurityAmazonResourceEvidenceRoles_Edited),
		string(SecurityAmazonResourceEvidenceRoles_Loaded),
		string(SecurityAmazonResourceEvidenceRoles_PolicyViolator),
		string(SecurityAmazonResourceEvidenceRoles_Scanned),
		string(SecurityAmazonResourceEvidenceRoles_Source),
		string(SecurityAmazonResourceEvidenceRoles_Suspicious),
		string(SecurityAmazonResourceEvidenceRoles_Unknown),
	}
}

func (s *SecurityAmazonResourceEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAmazonResourceEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAmazonResourceEvidenceRoles(input string) (*SecurityAmazonResourceEvidenceRoles, error) {
	vals := map[string]SecurityAmazonResourceEvidenceRoles{
		"added":             SecurityAmazonResourceEvidenceRoles_Added,
		"attacked":          SecurityAmazonResourceEvidenceRoles_Attacked,
		"attacker":          SecurityAmazonResourceEvidenceRoles_Attacker,
		"commandandcontrol": SecurityAmazonResourceEvidenceRoles_CommandAndControl,
		"compromised":       SecurityAmazonResourceEvidenceRoles_Compromised,
		"contextual":        SecurityAmazonResourceEvidenceRoles_Contextual,
		"created":           SecurityAmazonResourceEvidenceRoles_Created,
		"destination":       SecurityAmazonResourceEvidenceRoles_Destination,
		"edited":            SecurityAmazonResourceEvidenceRoles_Edited,
		"loaded":            SecurityAmazonResourceEvidenceRoles_Loaded,
		"policyviolator":    SecurityAmazonResourceEvidenceRoles_PolicyViolator,
		"scanned":           SecurityAmazonResourceEvidenceRoles_Scanned,
		"source":            SecurityAmazonResourceEvidenceRoles_Source,
		"suspicious":        SecurityAmazonResourceEvidenceRoles_Suspicious,
		"unknown":           SecurityAmazonResourceEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAmazonResourceEvidenceRoles(input)
	return &out, nil
}
