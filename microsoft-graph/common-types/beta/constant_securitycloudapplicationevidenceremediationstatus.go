package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCloudApplicationEvidenceRemediationStatus string

const (
	SecurityCloudApplicationEvidenceRemediationStatus_Blocked    SecurityCloudApplicationEvidenceRemediationStatus = "blocked"
	SecurityCloudApplicationEvidenceRemediationStatus_None       SecurityCloudApplicationEvidenceRemediationStatus = "none"
	SecurityCloudApplicationEvidenceRemediationStatus_NotFound   SecurityCloudApplicationEvidenceRemediationStatus = "notFound"
	SecurityCloudApplicationEvidenceRemediationStatus_Prevented  SecurityCloudApplicationEvidenceRemediationStatus = "prevented"
	SecurityCloudApplicationEvidenceRemediationStatus_Remediated SecurityCloudApplicationEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityCloudApplicationEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityCloudApplicationEvidenceRemediationStatus_Blocked),
		string(SecurityCloudApplicationEvidenceRemediationStatus_None),
		string(SecurityCloudApplicationEvidenceRemediationStatus_NotFound),
		string(SecurityCloudApplicationEvidenceRemediationStatus_Prevented),
		string(SecurityCloudApplicationEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityCloudApplicationEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityCloudApplicationEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityCloudApplicationEvidenceRemediationStatus(input string) (*SecurityCloudApplicationEvidenceRemediationStatus, error) {
	vals := map[string]SecurityCloudApplicationEvidenceRemediationStatus{
		"blocked":    SecurityCloudApplicationEvidenceRemediationStatus_Blocked,
		"none":       SecurityCloudApplicationEvidenceRemediationStatus_None,
		"notfound":   SecurityCloudApplicationEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityCloudApplicationEvidenceRemediationStatus_Prevented,
		"remediated": SecurityCloudApplicationEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCloudApplicationEvidenceRemediationStatus(input)
	return &out, nil
}
