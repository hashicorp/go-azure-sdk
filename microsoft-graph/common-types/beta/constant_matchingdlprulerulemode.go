package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MatchingDlpRuleRuleMode string

const (
	MatchingDlpRuleRuleMode_Audit           MatchingDlpRuleRuleMode = "audit"
	MatchingDlpRuleRuleMode_AuditAndNotify  MatchingDlpRuleRuleMode = "auditAndNotify"
	MatchingDlpRuleRuleMode_Enforce         MatchingDlpRuleRuleMode = "enforce"
	MatchingDlpRuleRuleMode_PendingDeletion MatchingDlpRuleRuleMode = "pendingDeletion"
	MatchingDlpRuleRuleMode_Test            MatchingDlpRuleRuleMode = "test"
)

func PossibleValuesForMatchingDlpRuleRuleMode() []string {
	return []string{
		string(MatchingDlpRuleRuleMode_Audit),
		string(MatchingDlpRuleRuleMode_AuditAndNotify),
		string(MatchingDlpRuleRuleMode_Enforce),
		string(MatchingDlpRuleRuleMode_PendingDeletion),
		string(MatchingDlpRuleRuleMode_Test),
	}
}

func (s *MatchingDlpRuleRuleMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMatchingDlpRuleRuleMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMatchingDlpRuleRuleMode(input string) (*MatchingDlpRuleRuleMode, error) {
	vals := map[string]MatchingDlpRuleRuleMode{
		"audit":           MatchingDlpRuleRuleMode_Audit,
		"auditandnotify":  MatchingDlpRuleRuleMode_AuditAndNotify,
		"enforce":         MatchingDlpRuleRuleMode_Enforce,
		"pendingdeletion": MatchingDlpRuleRuleMode_PendingDeletion,
		"test":            MatchingDlpRuleRuleMode_Test,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MatchingDlpRuleRuleMode(input)
	return &out, nil
}
