package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCloudApplicationEvidenceRoles string

const (
	SecurityCloudApplicationEvidenceRoles_Added             SecurityCloudApplicationEvidenceRoles = "added"
	SecurityCloudApplicationEvidenceRoles_Attacked          SecurityCloudApplicationEvidenceRoles = "attacked"
	SecurityCloudApplicationEvidenceRoles_Attacker          SecurityCloudApplicationEvidenceRoles = "attacker"
	SecurityCloudApplicationEvidenceRoles_CommandAndControl SecurityCloudApplicationEvidenceRoles = "commandAndControl"
	SecurityCloudApplicationEvidenceRoles_Compromised       SecurityCloudApplicationEvidenceRoles = "compromised"
	SecurityCloudApplicationEvidenceRoles_Contextual        SecurityCloudApplicationEvidenceRoles = "contextual"
	SecurityCloudApplicationEvidenceRoles_Created           SecurityCloudApplicationEvidenceRoles = "created"
	SecurityCloudApplicationEvidenceRoles_Destination       SecurityCloudApplicationEvidenceRoles = "destination"
	SecurityCloudApplicationEvidenceRoles_Edited            SecurityCloudApplicationEvidenceRoles = "edited"
	SecurityCloudApplicationEvidenceRoles_Loaded            SecurityCloudApplicationEvidenceRoles = "loaded"
	SecurityCloudApplicationEvidenceRoles_PolicyViolator    SecurityCloudApplicationEvidenceRoles = "policyViolator"
	SecurityCloudApplicationEvidenceRoles_Scanned           SecurityCloudApplicationEvidenceRoles = "scanned"
	SecurityCloudApplicationEvidenceRoles_Source            SecurityCloudApplicationEvidenceRoles = "source"
	SecurityCloudApplicationEvidenceRoles_Suspicious        SecurityCloudApplicationEvidenceRoles = "suspicious"
	SecurityCloudApplicationEvidenceRoles_Unknown           SecurityCloudApplicationEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityCloudApplicationEvidenceRoles() []string {
	return []string{
		string(SecurityCloudApplicationEvidenceRoles_Added),
		string(SecurityCloudApplicationEvidenceRoles_Attacked),
		string(SecurityCloudApplicationEvidenceRoles_Attacker),
		string(SecurityCloudApplicationEvidenceRoles_CommandAndControl),
		string(SecurityCloudApplicationEvidenceRoles_Compromised),
		string(SecurityCloudApplicationEvidenceRoles_Contextual),
		string(SecurityCloudApplicationEvidenceRoles_Created),
		string(SecurityCloudApplicationEvidenceRoles_Destination),
		string(SecurityCloudApplicationEvidenceRoles_Edited),
		string(SecurityCloudApplicationEvidenceRoles_Loaded),
		string(SecurityCloudApplicationEvidenceRoles_PolicyViolator),
		string(SecurityCloudApplicationEvidenceRoles_Scanned),
		string(SecurityCloudApplicationEvidenceRoles_Source),
		string(SecurityCloudApplicationEvidenceRoles_Suspicious),
		string(SecurityCloudApplicationEvidenceRoles_Unknown),
	}
}

func (s *SecurityCloudApplicationEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityCloudApplicationEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityCloudApplicationEvidenceRoles(input string) (*SecurityCloudApplicationEvidenceRoles, error) {
	vals := map[string]SecurityCloudApplicationEvidenceRoles{
		"added":             SecurityCloudApplicationEvidenceRoles_Added,
		"attacked":          SecurityCloudApplicationEvidenceRoles_Attacked,
		"attacker":          SecurityCloudApplicationEvidenceRoles_Attacker,
		"commandandcontrol": SecurityCloudApplicationEvidenceRoles_CommandAndControl,
		"compromised":       SecurityCloudApplicationEvidenceRoles_Compromised,
		"contextual":        SecurityCloudApplicationEvidenceRoles_Contextual,
		"created":           SecurityCloudApplicationEvidenceRoles_Created,
		"destination":       SecurityCloudApplicationEvidenceRoles_Destination,
		"edited":            SecurityCloudApplicationEvidenceRoles_Edited,
		"loaded":            SecurityCloudApplicationEvidenceRoles_Loaded,
		"policyviolator":    SecurityCloudApplicationEvidenceRoles_PolicyViolator,
		"scanned":           SecurityCloudApplicationEvidenceRoles_Scanned,
		"source":            SecurityCloudApplicationEvidenceRoles_Source,
		"suspicious":        SecurityCloudApplicationEvidenceRoles_Suspicious,
		"unknown":           SecurityCloudApplicationEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCloudApplicationEvidenceRoles(input)
	return &out, nil
}
