package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityOauthApplicationEvidenceRoles string

const (
	SecurityOauthApplicationEvidenceRoles_Added             SecurityOauthApplicationEvidenceRoles = "added"
	SecurityOauthApplicationEvidenceRoles_Attacked          SecurityOauthApplicationEvidenceRoles = "attacked"
	SecurityOauthApplicationEvidenceRoles_Attacker          SecurityOauthApplicationEvidenceRoles = "attacker"
	SecurityOauthApplicationEvidenceRoles_CommandAndControl SecurityOauthApplicationEvidenceRoles = "commandAndControl"
	SecurityOauthApplicationEvidenceRoles_Compromised       SecurityOauthApplicationEvidenceRoles = "compromised"
	SecurityOauthApplicationEvidenceRoles_Contextual        SecurityOauthApplicationEvidenceRoles = "contextual"
	SecurityOauthApplicationEvidenceRoles_Created           SecurityOauthApplicationEvidenceRoles = "created"
	SecurityOauthApplicationEvidenceRoles_Destination       SecurityOauthApplicationEvidenceRoles = "destination"
	SecurityOauthApplicationEvidenceRoles_Edited            SecurityOauthApplicationEvidenceRoles = "edited"
	SecurityOauthApplicationEvidenceRoles_Loaded            SecurityOauthApplicationEvidenceRoles = "loaded"
	SecurityOauthApplicationEvidenceRoles_PolicyViolator    SecurityOauthApplicationEvidenceRoles = "policyViolator"
	SecurityOauthApplicationEvidenceRoles_Scanned           SecurityOauthApplicationEvidenceRoles = "scanned"
	SecurityOauthApplicationEvidenceRoles_Source            SecurityOauthApplicationEvidenceRoles = "source"
	SecurityOauthApplicationEvidenceRoles_Suspicious        SecurityOauthApplicationEvidenceRoles = "suspicious"
	SecurityOauthApplicationEvidenceRoles_Unknown           SecurityOauthApplicationEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityOauthApplicationEvidenceRoles() []string {
	return []string{
		string(SecurityOauthApplicationEvidenceRoles_Added),
		string(SecurityOauthApplicationEvidenceRoles_Attacked),
		string(SecurityOauthApplicationEvidenceRoles_Attacker),
		string(SecurityOauthApplicationEvidenceRoles_CommandAndControl),
		string(SecurityOauthApplicationEvidenceRoles_Compromised),
		string(SecurityOauthApplicationEvidenceRoles_Contextual),
		string(SecurityOauthApplicationEvidenceRoles_Created),
		string(SecurityOauthApplicationEvidenceRoles_Destination),
		string(SecurityOauthApplicationEvidenceRoles_Edited),
		string(SecurityOauthApplicationEvidenceRoles_Loaded),
		string(SecurityOauthApplicationEvidenceRoles_PolicyViolator),
		string(SecurityOauthApplicationEvidenceRoles_Scanned),
		string(SecurityOauthApplicationEvidenceRoles_Source),
		string(SecurityOauthApplicationEvidenceRoles_Suspicious),
		string(SecurityOauthApplicationEvidenceRoles_Unknown),
	}
}

func (s *SecurityOauthApplicationEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityOauthApplicationEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityOauthApplicationEvidenceRoles(input string) (*SecurityOauthApplicationEvidenceRoles, error) {
	vals := map[string]SecurityOauthApplicationEvidenceRoles{
		"added":             SecurityOauthApplicationEvidenceRoles_Added,
		"attacked":          SecurityOauthApplicationEvidenceRoles_Attacked,
		"attacker":          SecurityOauthApplicationEvidenceRoles_Attacker,
		"commandandcontrol": SecurityOauthApplicationEvidenceRoles_CommandAndControl,
		"compromised":       SecurityOauthApplicationEvidenceRoles_Compromised,
		"contextual":        SecurityOauthApplicationEvidenceRoles_Contextual,
		"created":           SecurityOauthApplicationEvidenceRoles_Created,
		"destination":       SecurityOauthApplicationEvidenceRoles_Destination,
		"edited":            SecurityOauthApplicationEvidenceRoles_Edited,
		"loaded":            SecurityOauthApplicationEvidenceRoles_Loaded,
		"policyviolator":    SecurityOauthApplicationEvidenceRoles_PolicyViolator,
		"scanned":           SecurityOauthApplicationEvidenceRoles_Scanned,
		"source":            SecurityOauthApplicationEvidenceRoles_Source,
		"suspicious":        SecurityOauthApplicationEvidenceRoles_Suspicious,
		"unknown":           SecurityOauthApplicationEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityOauthApplicationEvidenceRoles(input)
	return &out, nil
}
