package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGoogleCloudResourceEvidenceRemediationStatus string

const (
	SecurityGoogleCloudResourceEvidenceRemediationStatus_Blocked    SecurityGoogleCloudResourceEvidenceRemediationStatus = "blocked"
	SecurityGoogleCloudResourceEvidenceRemediationStatus_None       SecurityGoogleCloudResourceEvidenceRemediationStatus = "none"
	SecurityGoogleCloudResourceEvidenceRemediationStatus_NotFound   SecurityGoogleCloudResourceEvidenceRemediationStatus = "notFound"
	SecurityGoogleCloudResourceEvidenceRemediationStatus_Prevented  SecurityGoogleCloudResourceEvidenceRemediationStatus = "prevented"
	SecurityGoogleCloudResourceEvidenceRemediationStatus_Remediated SecurityGoogleCloudResourceEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityGoogleCloudResourceEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityGoogleCloudResourceEvidenceRemediationStatus_Blocked),
		string(SecurityGoogleCloudResourceEvidenceRemediationStatus_None),
		string(SecurityGoogleCloudResourceEvidenceRemediationStatus_NotFound),
		string(SecurityGoogleCloudResourceEvidenceRemediationStatus_Prevented),
		string(SecurityGoogleCloudResourceEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityGoogleCloudResourceEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityGoogleCloudResourceEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityGoogleCloudResourceEvidenceRemediationStatus(input string) (*SecurityGoogleCloudResourceEvidenceRemediationStatus, error) {
	vals := map[string]SecurityGoogleCloudResourceEvidenceRemediationStatus{
		"blocked":    SecurityGoogleCloudResourceEvidenceRemediationStatus_Blocked,
		"none":       SecurityGoogleCloudResourceEvidenceRemediationStatus_None,
		"notfound":   SecurityGoogleCloudResourceEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityGoogleCloudResourceEvidenceRemediationStatus_Prevented,
		"remediated": SecurityGoogleCloudResourceEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGoogleCloudResourceEvidenceRemediationStatus(input)
	return &out, nil
}
