package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGitHubRepoEvidenceRemediationStatus string

const (
	SecurityGitHubRepoEvidenceRemediationStatus_Blocked    SecurityGitHubRepoEvidenceRemediationStatus = "blocked"
	SecurityGitHubRepoEvidenceRemediationStatus_None       SecurityGitHubRepoEvidenceRemediationStatus = "none"
	SecurityGitHubRepoEvidenceRemediationStatus_NotFound   SecurityGitHubRepoEvidenceRemediationStatus = "notFound"
	SecurityGitHubRepoEvidenceRemediationStatus_Prevented  SecurityGitHubRepoEvidenceRemediationStatus = "prevented"
	SecurityGitHubRepoEvidenceRemediationStatus_Remediated SecurityGitHubRepoEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityGitHubRepoEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityGitHubRepoEvidenceRemediationStatus_Blocked),
		string(SecurityGitHubRepoEvidenceRemediationStatus_None),
		string(SecurityGitHubRepoEvidenceRemediationStatus_NotFound),
		string(SecurityGitHubRepoEvidenceRemediationStatus_Prevented),
		string(SecurityGitHubRepoEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityGitHubRepoEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityGitHubRepoEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityGitHubRepoEvidenceRemediationStatus(input string) (*SecurityGitHubRepoEvidenceRemediationStatus, error) {
	vals := map[string]SecurityGitHubRepoEvidenceRemediationStatus{
		"blocked":    SecurityGitHubRepoEvidenceRemediationStatus_Blocked,
		"none":       SecurityGitHubRepoEvidenceRemediationStatus_None,
		"notfound":   SecurityGitHubRepoEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityGitHubRepoEvidenceRemediationStatus_Prevented,
		"remediated": SecurityGitHubRepoEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGitHubRepoEvidenceRemediationStatus(input)
	return &out, nil
}
