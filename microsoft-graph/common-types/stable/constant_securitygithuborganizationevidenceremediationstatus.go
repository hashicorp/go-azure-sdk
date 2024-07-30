package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGitHubOrganizationEvidenceRemediationStatus string

const (
	SecurityGitHubOrganizationEvidenceRemediationStatus_Blocked    SecurityGitHubOrganizationEvidenceRemediationStatus = "blocked"
	SecurityGitHubOrganizationEvidenceRemediationStatus_None       SecurityGitHubOrganizationEvidenceRemediationStatus = "none"
	SecurityGitHubOrganizationEvidenceRemediationStatus_NotFound   SecurityGitHubOrganizationEvidenceRemediationStatus = "notFound"
	SecurityGitHubOrganizationEvidenceRemediationStatus_Prevented  SecurityGitHubOrganizationEvidenceRemediationStatus = "prevented"
	SecurityGitHubOrganizationEvidenceRemediationStatus_Remediated SecurityGitHubOrganizationEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityGitHubOrganizationEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityGitHubOrganizationEvidenceRemediationStatus_Blocked),
		string(SecurityGitHubOrganizationEvidenceRemediationStatus_None),
		string(SecurityGitHubOrganizationEvidenceRemediationStatus_NotFound),
		string(SecurityGitHubOrganizationEvidenceRemediationStatus_Prevented),
		string(SecurityGitHubOrganizationEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityGitHubOrganizationEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityGitHubOrganizationEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityGitHubOrganizationEvidenceRemediationStatus(input string) (*SecurityGitHubOrganizationEvidenceRemediationStatus, error) {
	vals := map[string]SecurityGitHubOrganizationEvidenceRemediationStatus{
		"blocked":    SecurityGitHubOrganizationEvidenceRemediationStatus_Blocked,
		"none":       SecurityGitHubOrganizationEvidenceRemediationStatus_None,
		"notfound":   SecurityGitHubOrganizationEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityGitHubOrganizationEvidenceRemediationStatus_Prevented,
		"remediated": SecurityGitHubOrganizationEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGitHubOrganizationEvidenceRemediationStatus(input)
	return &out, nil
}
