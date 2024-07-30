package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGitHubUserEvidenceRemediationStatus string

const (
	SecurityGitHubUserEvidenceRemediationStatus_Blocked    SecurityGitHubUserEvidenceRemediationStatus = "blocked"
	SecurityGitHubUserEvidenceRemediationStatus_None       SecurityGitHubUserEvidenceRemediationStatus = "none"
	SecurityGitHubUserEvidenceRemediationStatus_NotFound   SecurityGitHubUserEvidenceRemediationStatus = "notFound"
	SecurityGitHubUserEvidenceRemediationStatus_Prevented  SecurityGitHubUserEvidenceRemediationStatus = "prevented"
	SecurityGitHubUserEvidenceRemediationStatus_Remediated SecurityGitHubUserEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityGitHubUserEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityGitHubUserEvidenceRemediationStatus_Blocked),
		string(SecurityGitHubUserEvidenceRemediationStatus_None),
		string(SecurityGitHubUserEvidenceRemediationStatus_NotFound),
		string(SecurityGitHubUserEvidenceRemediationStatus_Prevented),
		string(SecurityGitHubUserEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityGitHubUserEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityGitHubUserEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityGitHubUserEvidenceRemediationStatus(input string) (*SecurityGitHubUserEvidenceRemediationStatus, error) {
	vals := map[string]SecurityGitHubUserEvidenceRemediationStatus{
		"blocked":    SecurityGitHubUserEvidenceRemediationStatus_Blocked,
		"none":       SecurityGitHubUserEvidenceRemediationStatus_None,
		"notfound":   SecurityGitHubUserEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityGitHubUserEvidenceRemediationStatus_Prevented,
		"remediated": SecurityGitHubUserEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGitHubUserEvidenceRemediationStatus(input)
	return &out, nil
}
