package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTenantAllowBlockListEntryResultStatus string

const (
	SecurityTenantAllowBlockListEntryResultStatus_Failed     SecurityTenantAllowBlockListEntryResultStatus = "failed"
	SecurityTenantAllowBlockListEntryResultStatus_NotStarted SecurityTenantAllowBlockListEntryResultStatus = "notStarted"
	SecurityTenantAllowBlockListEntryResultStatus_Running    SecurityTenantAllowBlockListEntryResultStatus = "running"
	SecurityTenantAllowBlockListEntryResultStatus_Skipped    SecurityTenantAllowBlockListEntryResultStatus = "skipped"
	SecurityTenantAllowBlockListEntryResultStatus_Succeeded  SecurityTenantAllowBlockListEntryResultStatus = "succeeded"
)

func PossibleValuesForSecurityTenantAllowBlockListEntryResultStatus() []string {
	return []string{
		string(SecurityTenantAllowBlockListEntryResultStatus_Failed),
		string(SecurityTenantAllowBlockListEntryResultStatus_NotStarted),
		string(SecurityTenantAllowBlockListEntryResultStatus_Running),
		string(SecurityTenantAllowBlockListEntryResultStatus_Skipped),
		string(SecurityTenantAllowBlockListEntryResultStatus_Succeeded),
	}
}

func (s *SecurityTenantAllowBlockListEntryResultStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityTenantAllowBlockListEntryResultStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityTenantAllowBlockListEntryResultStatus(input string) (*SecurityTenantAllowBlockListEntryResultStatus, error) {
	vals := map[string]SecurityTenantAllowBlockListEntryResultStatus{
		"failed":     SecurityTenantAllowBlockListEntryResultStatus_Failed,
		"notstarted": SecurityTenantAllowBlockListEntryResultStatus_NotStarted,
		"running":    SecurityTenantAllowBlockListEntryResultStatus_Running,
		"skipped":    SecurityTenantAllowBlockListEntryResultStatus_Skipped,
		"succeeded":  SecurityTenantAllowBlockListEntryResultStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityTenantAllowBlockListEntryResultStatus(input)
	return &out, nil
}
