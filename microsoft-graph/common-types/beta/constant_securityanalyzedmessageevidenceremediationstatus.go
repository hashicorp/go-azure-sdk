package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedMessageEvidenceRemediationStatus string

const (
	SecurityAnalyzedMessageEvidenceRemediationStatus_Blocked    SecurityAnalyzedMessageEvidenceRemediationStatus = "blocked"
	SecurityAnalyzedMessageEvidenceRemediationStatus_None       SecurityAnalyzedMessageEvidenceRemediationStatus = "none"
	SecurityAnalyzedMessageEvidenceRemediationStatus_NotFound   SecurityAnalyzedMessageEvidenceRemediationStatus = "notFound"
	SecurityAnalyzedMessageEvidenceRemediationStatus_Prevented  SecurityAnalyzedMessageEvidenceRemediationStatus = "prevented"
	SecurityAnalyzedMessageEvidenceRemediationStatus_Remediated SecurityAnalyzedMessageEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityAnalyzedMessageEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityAnalyzedMessageEvidenceRemediationStatus_Blocked),
		string(SecurityAnalyzedMessageEvidenceRemediationStatus_None),
		string(SecurityAnalyzedMessageEvidenceRemediationStatus_NotFound),
		string(SecurityAnalyzedMessageEvidenceRemediationStatus_Prevented),
		string(SecurityAnalyzedMessageEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityAnalyzedMessageEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAnalyzedMessageEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAnalyzedMessageEvidenceRemediationStatus(input string) (*SecurityAnalyzedMessageEvidenceRemediationStatus, error) {
	vals := map[string]SecurityAnalyzedMessageEvidenceRemediationStatus{
		"blocked":    SecurityAnalyzedMessageEvidenceRemediationStatus_Blocked,
		"none":       SecurityAnalyzedMessageEvidenceRemediationStatus_None,
		"notfound":   SecurityAnalyzedMessageEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityAnalyzedMessageEvidenceRemediationStatus_Prevented,
		"remediated": SecurityAnalyzedMessageEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAnalyzedMessageEvidenceRemediationStatus(input)
	return &out, nil
}
