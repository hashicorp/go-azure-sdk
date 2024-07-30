package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityOauthApplicationEvidenceRemediationStatus string

const (
	SecurityOauthApplicationEvidenceRemediationStatus_Blocked    SecurityOauthApplicationEvidenceRemediationStatus = "blocked"
	SecurityOauthApplicationEvidenceRemediationStatus_None       SecurityOauthApplicationEvidenceRemediationStatus = "none"
	SecurityOauthApplicationEvidenceRemediationStatus_NotFound   SecurityOauthApplicationEvidenceRemediationStatus = "notFound"
	SecurityOauthApplicationEvidenceRemediationStatus_Prevented  SecurityOauthApplicationEvidenceRemediationStatus = "prevented"
	SecurityOauthApplicationEvidenceRemediationStatus_Remediated SecurityOauthApplicationEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityOauthApplicationEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityOauthApplicationEvidenceRemediationStatus_Blocked),
		string(SecurityOauthApplicationEvidenceRemediationStatus_None),
		string(SecurityOauthApplicationEvidenceRemediationStatus_NotFound),
		string(SecurityOauthApplicationEvidenceRemediationStatus_Prevented),
		string(SecurityOauthApplicationEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityOauthApplicationEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityOauthApplicationEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityOauthApplicationEvidenceRemediationStatus(input string) (*SecurityOauthApplicationEvidenceRemediationStatus, error) {
	vals := map[string]SecurityOauthApplicationEvidenceRemediationStatus{
		"blocked":    SecurityOauthApplicationEvidenceRemediationStatus_Blocked,
		"none":       SecurityOauthApplicationEvidenceRemediationStatus_None,
		"notfound":   SecurityOauthApplicationEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityOauthApplicationEvidenceRemediationStatus_Prevented,
		"remediated": SecurityOauthApplicationEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityOauthApplicationEvidenceRemediationStatus(input)
	return &out, nil
}
