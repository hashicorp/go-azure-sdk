package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionLabelBehaviorDuringRetentionPeriod string

const (
	SecurityRetentionLabelBehaviorDuringRetentionPeriod_DoNotRetain              SecurityRetentionLabelBehaviorDuringRetentionPeriod = "doNotRetain"
	SecurityRetentionLabelBehaviorDuringRetentionPeriod_Retain                   SecurityRetentionLabelBehaviorDuringRetentionPeriod = "retain"
	SecurityRetentionLabelBehaviorDuringRetentionPeriod_RetainAsRecord           SecurityRetentionLabelBehaviorDuringRetentionPeriod = "retainAsRecord"
	SecurityRetentionLabelBehaviorDuringRetentionPeriod_RetainAsRegulatoryRecord SecurityRetentionLabelBehaviorDuringRetentionPeriod = "retainAsRegulatoryRecord"
)

func PossibleValuesForSecurityRetentionLabelBehaviorDuringRetentionPeriod() []string {
	return []string{
		string(SecurityRetentionLabelBehaviorDuringRetentionPeriod_DoNotRetain),
		string(SecurityRetentionLabelBehaviorDuringRetentionPeriod_Retain),
		string(SecurityRetentionLabelBehaviorDuringRetentionPeriod_RetainAsRecord),
		string(SecurityRetentionLabelBehaviorDuringRetentionPeriod_RetainAsRegulatoryRecord),
	}
}

func (s *SecurityRetentionLabelBehaviorDuringRetentionPeriod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRetentionLabelBehaviorDuringRetentionPeriod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRetentionLabelBehaviorDuringRetentionPeriod(input string) (*SecurityRetentionLabelBehaviorDuringRetentionPeriod, error) {
	vals := map[string]SecurityRetentionLabelBehaviorDuringRetentionPeriod{
		"donotretain":              SecurityRetentionLabelBehaviorDuringRetentionPeriod_DoNotRetain,
		"retain":                   SecurityRetentionLabelBehaviorDuringRetentionPeriod_Retain,
		"retainasrecord":           SecurityRetentionLabelBehaviorDuringRetentionPeriod_RetainAsRecord,
		"retainasregulatoryrecord": SecurityRetentionLabelBehaviorDuringRetentionPeriod_RetainAsRegulatoryRecord,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRetentionLabelBehaviorDuringRetentionPeriod(input)
	return &out, nil
}
