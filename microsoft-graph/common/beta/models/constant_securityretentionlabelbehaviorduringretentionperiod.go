package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionLabelBehaviorDuringRetentionPeriod string

const (
	SecurityRetentionLabelBehaviorDuringRetentionPerioddoNotRetain              SecurityRetentionLabelBehaviorDuringRetentionPeriod = "DoNotRetain"
	SecurityRetentionLabelBehaviorDuringRetentionPeriodretain                   SecurityRetentionLabelBehaviorDuringRetentionPeriod = "Retain"
	SecurityRetentionLabelBehaviorDuringRetentionPeriodretainAsRecord           SecurityRetentionLabelBehaviorDuringRetentionPeriod = "RetainAsRecord"
	SecurityRetentionLabelBehaviorDuringRetentionPeriodretainAsRegulatoryRecord SecurityRetentionLabelBehaviorDuringRetentionPeriod = "RetainAsRegulatoryRecord"
)

func PossibleValuesForSecurityRetentionLabelBehaviorDuringRetentionPeriod() []string {
	return []string{
		string(SecurityRetentionLabelBehaviorDuringRetentionPerioddoNotRetain),
		string(SecurityRetentionLabelBehaviorDuringRetentionPeriodretain),
		string(SecurityRetentionLabelBehaviorDuringRetentionPeriodretainAsRecord),
		string(SecurityRetentionLabelBehaviorDuringRetentionPeriodretainAsRegulatoryRecord),
	}
}

func parseSecurityRetentionLabelBehaviorDuringRetentionPeriod(input string) (*SecurityRetentionLabelBehaviorDuringRetentionPeriod, error) {
	vals := map[string]SecurityRetentionLabelBehaviorDuringRetentionPeriod{
		"donotretain":              SecurityRetentionLabelBehaviorDuringRetentionPerioddoNotRetain,
		"retain":                   SecurityRetentionLabelBehaviorDuringRetentionPeriodretain,
		"retainasrecord":           SecurityRetentionLabelBehaviorDuringRetentionPeriodretainAsRecord,
		"retainasregulatoryrecord": SecurityRetentionLabelBehaviorDuringRetentionPeriodretainAsRegulatoryRecord,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRetentionLabelBehaviorDuringRetentionPeriod(input)
	return &out, nil
}
