package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedMessageEvidenceRoles string

const (
	SecurityAnalyzedMessageEvidenceRoles_Added             SecurityAnalyzedMessageEvidenceRoles = "added"
	SecurityAnalyzedMessageEvidenceRoles_Attacked          SecurityAnalyzedMessageEvidenceRoles = "attacked"
	SecurityAnalyzedMessageEvidenceRoles_Attacker          SecurityAnalyzedMessageEvidenceRoles = "attacker"
	SecurityAnalyzedMessageEvidenceRoles_CommandAndControl SecurityAnalyzedMessageEvidenceRoles = "commandAndControl"
	SecurityAnalyzedMessageEvidenceRoles_Compromised       SecurityAnalyzedMessageEvidenceRoles = "compromised"
	SecurityAnalyzedMessageEvidenceRoles_Contextual        SecurityAnalyzedMessageEvidenceRoles = "contextual"
	SecurityAnalyzedMessageEvidenceRoles_Created           SecurityAnalyzedMessageEvidenceRoles = "created"
	SecurityAnalyzedMessageEvidenceRoles_Destination       SecurityAnalyzedMessageEvidenceRoles = "destination"
	SecurityAnalyzedMessageEvidenceRoles_Edited            SecurityAnalyzedMessageEvidenceRoles = "edited"
	SecurityAnalyzedMessageEvidenceRoles_Loaded            SecurityAnalyzedMessageEvidenceRoles = "loaded"
	SecurityAnalyzedMessageEvidenceRoles_PolicyViolator    SecurityAnalyzedMessageEvidenceRoles = "policyViolator"
	SecurityAnalyzedMessageEvidenceRoles_Scanned           SecurityAnalyzedMessageEvidenceRoles = "scanned"
	SecurityAnalyzedMessageEvidenceRoles_Source            SecurityAnalyzedMessageEvidenceRoles = "source"
	SecurityAnalyzedMessageEvidenceRoles_Suspicious        SecurityAnalyzedMessageEvidenceRoles = "suspicious"
	SecurityAnalyzedMessageEvidenceRoles_Unknown           SecurityAnalyzedMessageEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityAnalyzedMessageEvidenceRoles() []string {
	return []string{
		string(SecurityAnalyzedMessageEvidenceRoles_Added),
		string(SecurityAnalyzedMessageEvidenceRoles_Attacked),
		string(SecurityAnalyzedMessageEvidenceRoles_Attacker),
		string(SecurityAnalyzedMessageEvidenceRoles_CommandAndControl),
		string(SecurityAnalyzedMessageEvidenceRoles_Compromised),
		string(SecurityAnalyzedMessageEvidenceRoles_Contextual),
		string(SecurityAnalyzedMessageEvidenceRoles_Created),
		string(SecurityAnalyzedMessageEvidenceRoles_Destination),
		string(SecurityAnalyzedMessageEvidenceRoles_Edited),
		string(SecurityAnalyzedMessageEvidenceRoles_Loaded),
		string(SecurityAnalyzedMessageEvidenceRoles_PolicyViolator),
		string(SecurityAnalyzedMessageEvidenceRoles_Scanned),
		string(SecurityAnalyzedMessageEvidenceRoles_Source),
		string(SecurityAnalyzedMessageEvidenceRoles_Suspicious),
		string(SecurityAnalyzedMessageEvidenceRoles_Unknown),
	}
}

func (s *SecurityAnalyzedMessageEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAnalyzedMessageEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAnalyzedMessageEvidenceRoles(input string) (*SecurityAnalyzedMessageEvidenceRoles, error) {
	vals := map[string]SecurityAnalyzedMessageEvidenceRoles{
		"added":             SecurityAnalyzedMessageEvidenceRoles_Added,
		"attacked":          SecurityAnalyzedMessageEvidenceRoles_Attacked,
		"attacker":          SecurityAnalyzedMessageEvidenceRoles_Attacker,
		"commandandcontrol": SecurityAnalyzedMessageEvidenceRoles_CommandAndControl,
		"compromised":       SecurityAnalyzedMessageEvidenceRoles_Compromised,
		"contextual":        SecurityAnalyzedMessageEvidenceRoles_Contextual,
		"created":           SecurityAnalyzedMessageEvidenceRoles_Created,
		"destination":       SecurityAnalyzedMessageEvidenceRoles_Destination,
		"edited":            SecurityAnalyzedMessageEvidenceRoles_Edited,
		"loaded":            SecurityAnalyzedMessageEvidenceRoles_Loaded,
		"policyviolator":    SecurityAnalyzedMessageEvidenceRoles_PolicyViolator,
		"scanned":           SecurityAnalyzedMessageEvidenceRoles_Scanned,
		"source":            SecurityAnalyzedMessageEvidenceRoles_Source,
		"suspicious":        SecurityAnalyzedMessageEvidenceRoles_Suspicious,
		"unknown":           SecurityAnalyzedMessageEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAnalyzedMessageEvidenceRoles(input)
	return &out, nil
}
