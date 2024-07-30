package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetentionLabelSettingsBehaviorDuringRetentionPeriod string

const (
	RetentionLabelSettingsBehaviorDuringRetentionPeriod_DoNotRetain              RetentionLabelSettingsBehaviorDuringRetentionPeriod = "doNotRetain"
	RetentionLabelSettingsBehaviorDuringRetentionPeriod_Retain                   RetentionLabelSettingsBehaviorDuringRetentionPeriod = "retain"
	RetentionLabelSettingsBehaviorDuringRetentionPeriod_RetainAsRecord           RetentionLabelSettingsBehaviorDuringRetentionPeriod = "retainAsRecord"
	RetentionLabelSettingsBehaviorDuringRetentionPeriod_RetainAsRegulatoryRecord RetentionLabelSettingsBehaviorDuringRetentionPeriod = "retainAsRegulatoryRecord"
)

func PossibleValuesForRetentionLabelSettingsBehaviorDuringRetentionPeriod() []string {
	return []string{
		string(RetentionLabelSettingsBehaviorDuringRetentionPeriod_DoNotRetain),
		string(RetentionLabelSettingsBehaviorDuringRetentionPeriod_Retain),
		string(RetentionLabelSettingsBehaviorDuringRetentionPeriod_RetainAsRecord),
		string(RetentionLabelSettingsBehaviorDuringRetentionPeriod_RetainAsRegulatoryRecord),
	}
}

func (s *RetentionLabelSettingsBehaviorDuringRetentionPeriod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRetentionLabelSettingsBehaviorDuringRetentionPeriod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRetentionLabelSettingsBehaviorDuringRetentionPeriod(input string) (*RetentionLabelSettingsBehaviorDuringRetentionPeriod, error) {
	vals := map[string]RetentionLabelSettingsBehaviorDuringRetentionPeriod{
		"donotretain":              RetentionLabelSettingsBehaviorDuringRetentionPeriod_DoNotRetain,
		"retain":                   RetentionLabelSettingsBehaviorDuringRetentionPeriod_Retain,
		"retainasrecord":           RetentionLabelSettingsBehaviorDuringRetentionPeriod_RetainAsRecord,
		"retainasregulatoryrecord": RetentionLabelSettingsBehaviorDuringRetentionPeriod_RetainAsRegulatoryRecord,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RetentionLabelSettingsBehaviorDuringRetentionPeriod(input)
	return &out, nil
}
