package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkSettingsEnrollmentTarget string

const (
	AndroidForWorkSettingsEnrollmentTargetall                              AndroidForWorkSettingsEnrollmentTarget = "All"
	AndroidForWorkSettingsEnrollmentTargetnone                             AndroidForWorkSettingsEnrollmentTarget = "None"
	AndroidForWorkSettingsEnrollmentTargettargeted                         AndroidForWorkSettingsEnrollmentTarget = "Targeted"
	AndroidForWorkSettingsEnrollmentTargettargetedAsEnrollmentRestrictions AndroidForWorkSettingsEnrollmentTarget = "TargetedAsEnrollmentRestrictions"
)

func PossibleValuesForAndroidForWorkSettingsEnrollmentTarget() []string {
	return []string{
		string(AndroidForWorkSettingsEnrollmentTargetall),
		string(AndroidForWorkSettingsEnrollmentTargetnone),
		string(AndroidForWorkSettingsEnrollmentTargettargeted),
		string(AndroidForWorkSettingsEnrollmentTargettargetedAsEnrollmentRestrictions),
	}
}

func parseAndroidForWorkSettingsEnrollmentTarget(input string) (*AndroidForWorkSettingsEnrollmentTarget, error) {
	vals := map[string]AndroidForWorkSettingsEnrollmentTarget{
		"all":                              AndroidForWorkSettingsEnrollmentTargetall,
		"none":                             AndroidForWorkSettingsEnrollmentTargetnone,
		"targeted":                         AndroidForWorkSettingsEnrollmentTargettargeted,
		"targetedasenrollmentrestrictions": AndroidForWorkSettingsEnrollmentTargettargetedAsEnrollmentRestrictions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkSettingsEnrollmentTarget(input)
	return &out, nil
}
