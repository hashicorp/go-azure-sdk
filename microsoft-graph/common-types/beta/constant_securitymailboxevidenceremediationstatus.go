package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailboxEvidenceRemediationStatus string

const (
	SecurityMailboxEvidenceRemediationStatus_Blocked    SecurityMailboxEvidenceRemediationStatus = "blocked"
	SecurityMailboxEvidenceRemediationStatus_None       SecurityMailboxEvidenceRemediationStatus = "none"
	SecurityMailboxEvidenceRemediationStatus_NotFound   SecurityMailboxEvidenceRemediationStatus = "notFound"
	SecurityMailboxEvidenceRemediationStatus_Prevented  SecurityMailboxEvidenceRemediationStatus = "prevented"
	SecurityMailboxEvidenceRemediationStatus_Remediated SecurityMailboxEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityMailboxEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityMailboxEvidenceRemediationStatus_Blocked),
		string(SecurityMailboxEvidenceRemediationStatus_None),
		string(SecurityMailboxEvidenceRemediationStatus_NotFound),
		string(SecurityMailboxEvidenceRemediationStatus_Prevented),
		string(SecurityMailboxEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityMailboxEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityMailboxEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityMailboxEvidenceRemediationStatus(input string) (*SecurityMailboxEvidenceRemediationStatus, error) {
	vals := map[string]SecurityMailboxEvidenceRemediationStatus{
		"blocked":    SecurityMailboxEvidenceRemediationStatus_Blocked,
		"none":       SecurityMailboxEvidenceRemediationStatus_None,
		"notfound":   SecurityMailboxEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityMailboxEvidenceRemediationStatus_Prevented,
		"remediated": SecurityMailboxEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMailboxEvidenceRemediationStatus(input)
	return &out, nil
}
