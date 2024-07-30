package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityNetworkConnectionEvidenceRoles string

const (
	SecurityNetworkConnectionEvidenceRoles_Added             SecurityNetworkConnectionEvidenceRoles = "added"
	SecurityNetworkConnectionEvidenceRoles_Attacked          SecurityNetworkConnectionEvidenceRoles = "attacked"
	SecurityNetworkConnectionEvidenceRoles_Attacker          SecurityNetworkConnectionEvidenceRoles = "attacker"
	SecurityNetworkConnectionEvidenceRoles_CommandAndControl SecurityNetworkConnectionEvidenceRoles = "commandAndControl"
	SecurityNetworkConnectionEvidenceRoles_Compromised       SecurityNetworkConnectionEvidenceRoles = "compromised"
	SecurityNetworkConnectionEvidenceRoles_Contextual        SecurityNetworkConnectionEvidenceRoles = "contextual"
	SecurityNetworkConnectionEvidenceRoles_Created           SecurityNetworkConnectionEvidenceRoles = "created"
	SecurityNetworkConnectionEvidenceRoles_Destination       SecurityNetworkConnectionEvidenceRoles = "destination"
	SecurityNetworkConnectionEvidenceRoles_Edited            SecurityNetworkConnectionEvidenceRoles = "edited"
	SecurityNetworkConnectionEvidenceRoles_Loaded            SecurityNetworkConnectionEvidenceRoles = "loaded"
	SecurityNetworkConnectionEvidenceRoles_PolicyViolator    SecurityNetworkConnectionEvidenceRoles = "policyViolator"
	SecurityNetworkConnectionEvidenceRoles_Scanned           SecurityNetworkConnectionEvidenceRoles = "scanned"
	SecurityNetworkConnectionEvidenceRoles_Source            SecurityNetworkConnectionEvidenceRoles = "source"
	SecurityNetworkConnectionEvidenceRoles_Suspicious        SecurityNetworkConnectionEvidenceRoles = "suspicious"
	SecurityNetworkConnectionEvidenceRoles_Unknown           SecurityNetworkConnectionEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityNetworkConnectionEvidenceRoles() []string {
	return []string{
		string(SecurityNetworkConnectionEvidenceRoles_Added),
		string(SecurityNetworkConnectionEvidenceRoles_Attacked),
		string(SecurityNetworkConnectionEvidenceRoles_Attacker),
		string(SecurityNetworkConnectionEvidenceRoles_CommandAndControl),
		string(SecurityNetworkConnectionEvidenceRoles_Compromised),
		string(SecurityNetworkConnectionEvidenceRoles_Contextual),
		string(SecurityNetworkConnectionEvidenceRoles_Created),
		string(SecurityNetworkConnectionEvidenceRoles_Destination),
		string(SecurityNetworkConnectionEvidenceRoles_Edited),
		string(SecurityNetworkConnectionEvidenceRoles_Loaded),
		string(SecurityNetworkConnectionEvidenceRoles_PolicyViolator),
		string(SecurityNetworkConnectionEvidenceRoles_Scanned),
		string(SecurityNetworkConnectionEvidenceRoles_Source),
		string(SecurityNetworkConnectionEvidenceRoles_Suspicious),
		string(SecurityNetworkConnectionEvidenceRoles_Unknown),
	}
}

func (s *SecurityNetworkConnectionEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityNetworkConnectionEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityNetworkConnectionEvidenceRoles(input string) (*SecurityNetworkConnectionEvidenceRoles, error) {
	vals := map[string]SecurityNetworkConnectionEvidenceRoles{
		"added":             SecurityNetworkConnectionEvidenceRoles_Added,
		"attacked":          SecurityNetworkConnectionEvidenceRoles_Attacked,
		"attacker":          SecurityNetworkConnectionEvidenceRoles_Attacker,
		"commandandcontrol": SecurityNetworkConnectionEvidenceRoles_CommandAndControl,
		"compromised":       SecurityNetworkConnectionEvidenceRoles_Compromised,
		"contextual":        SecurityNetworkConnectionEvidenceRoles_Contextual,
		"created":           SecurityNetworkConnectionEvidenceRoles_Created,
		"destination":       SecurityNetworkConnectionEvidenceRoles_Destination,
		"edited":            SecurityNetworkConnectionEvidenceRoles_Edited,
		"loaded":            SecurityNetworkConnectionEvidenceRoles_Loaded,
		"policyviolator":    SecurityNetworkConnectionEvidenceRoles_PolicyViolator,
		"scanned":           SecurityNetworkConnectionEvidenceRoles_Scanned,
		"source":            SecurityNetworkConnectionEvidenceRoles_Source,
		"suspicious":        SecurityNetworkConnectionEvidenceRoles_Suspicious,
		"unknown":           SecurityNetworkConnectionEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityNetworkConnectionEvidenceRoles(input)
	return &out, nil
}
