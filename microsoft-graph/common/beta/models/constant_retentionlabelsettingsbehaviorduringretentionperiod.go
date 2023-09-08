package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetentionLabelSettingsBehaviorDuringRetentionPeriod string

const (
	RetentionLabelSettingsBehaviorDuringRetentionPerioddoNotRetain              RetentionLabelSettingsBehaviorDuringRetentionPeriod = "DoNotRetain"
	RetentionLabelSettingsBehaviorDuringRetentionPeriodretain                   RetentionLabelSettingsBehaviorDuringRetentionPeriod = "Retain"
	RetentionLabelSettingsBehaviorDuringRetentionPeriodretainAsRecord           RetentionLabelSettingsBehaviorDuringRetentionPeriod = "RetainAsRecord"
	RetentionLabelSettingsBehaviorDuringRetentionPeriodretainAsRegulatoryRecord RetentionLabelSettingsBehaviorDuringRetentionPeriod = "RetainAsRegulatoryRecord"
)

func PossibleValuesForRetentionLabelSettingsBehaviorDuringRetentionPeriod() []string {
	return []string{
		string(RetentionLabelSettingsBehaviorDuringRetentionPerioddoNotRetain),
		string(RetentionLabelSettingsBehaviorDuringRetentionPeriodretain),
		string(RetentionLabelSettingsBehaviorDuringRetentionPeriodretainAsRecord),
		string(RetentionLabelSettingsBehaviorDuringRetentionPeriodretainAsRegulatoryRecord),
	}
}

func parseRetentionLabelSettingsBehaviorDuringRetentionPeriod(input string) (*RetentionLabelSettingsBehaviorDuringRetentionPeriod, error) {
	vals := map[string]RetentionLabelSettingsBehaviorDuringRetentionPeriod{
		"donotretain":              RetentionLabelSettingsBehaviorDuringRetentionPerioddoNotRetain,
		"retain":                   RetentionLabelSettingsBehaviorDuringRetentionPeriodretain,
		"retainasrecord":           RetentionLabelSettingsBehaviorDuringRetentionPeriodretainAsRecord,
		"retainasregulatoryrecord": RetentionLabelSettingsBehaviorDuringRetentionPeriodretainAsRegulatoryRecord,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RetentionLabelSettingsBehaviorDuringRetentionPeriod(input)
	return &out, nil
}
