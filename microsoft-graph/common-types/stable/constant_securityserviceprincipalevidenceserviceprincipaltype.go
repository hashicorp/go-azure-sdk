package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityServicePrincipalEvidenceServicePrincipalType string

const (
	SecurityServicePrincipalEvidenceServicePrincipalType_Application     SecurityServicePrincipalEvidenceServicePrincipalType = "application"
	SecurityServicePrincipalEvidenceServicePrincipalType_Legacy          SecurityServicePrincipalEvidenceServicePrincipalType = "legacy"
	SecurityServicePrincipalEvidenceServicePrincipalType_ManagedIdentity SecurityServicePrincipalEvidenceServicePrincipalType = "managedIdentity"
	SecurityServicePrincipalEvidenceServicePrincipalType_Unknown         SecurityServicePrincipalEvidenceServicePrincipalType = "unknown"
)

func PossibleValuesForSecurityServicePrincipalEvidenceServicePrincipalType() []string {
	return []string{
		string(SecurityServicePrincipalEvidenceServicePrincipalType_Application),
		string(SecurityServicePrincipalEvidenceServicePrincipalType_Legacy),
		string(SecurityServicePrincipalEvidenceServicePrincipalType_ManagedIdentity),
		string(SecurityServicePrincipalEvidenceServicePrincipalType_Unknown),
	}
}

func (s *SecurityServicePrincipalEvidenceServicePrincipalType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityServicePrincipalEvidenceServicePrincipalType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityServicePrincipalEvidenceServicePrincipalType(input string) (*SecurityServicePrincipalEvidenceServicePrincipalType, error) {
	vals := map[string]SecurityServicePrincipalEvidenceServicePrincipalType{
		"application":     SecurityServicePrincipalEvidenceServicePrincipalType_Application,
		"legacy":          SecurityServicePrincipalEvidenceServicePrincipalType_Legacy,
		"managedidentity": SecurityServicePrincipalEvidenceServicePrincipalType_ManagedIdentity,
		"unknown":         SecurityServicePrincipalEvidenceServicePrincipalType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityServicePrincipalEvidenceServicePrincipalType(input)
	return &out, nil
}
