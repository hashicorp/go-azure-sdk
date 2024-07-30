package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionLabelDefaultRecordBehavior string

const (
	SecurityRetentionLabelDefaultRecordBehavior_StartLocked   SecurityRetentionLabelDefaultRecordBehavior = "startLocked"
	SecurityRetentionLabelDefaultRecordBehavior_StartUnlocked SecurityRetentionLabelDefaultRecordBehavior = "startUnlocked"
)

func PossibleValuesForSecurityRetentionLabelDefaultRecordBehavior() []string {
	return []string{
		string(SecurityRetentionLabelDefaultRecordBehavior_StartLocked),
		string(SecurityRetentionLabelDefaultRecordBehavior_StartUnlocked),
	}
}

func (s *SecurityRetentionLabelDefaultRecordBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRetentionLabelDefaultRecordBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRetentionLabelDefaultRecordBehavior(input string) (*SecurityRetentionLabelDefaultRecordBehavior, error) {
	vals := map[string]SecurityRetentionLabelDefaultRecordBehavior{
		"startlocked":   SecurityRetentionLabelDefaultRecordBehavior_StartLocked,
		"startunlocked": SecurityRetentionLabelDefaultRecordBehavior_StartUnlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRetentionLabelDefaultRecordBehavior(input)
	return &out, nil
}
