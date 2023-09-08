package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget string

const (
	AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTargetall                              AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget = "All"
	AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTargetnone                             AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget = "None"
	AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTargettargeted                         AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget = "Targeted"
	AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTargettargetedAsEnrollmentRestrictions AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget = "TargetedAsEnrollmentRestrictions"
)

func PossibleValuesForAndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget() []string {
	return []string{
		string(AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTargetall),
		string(AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTargetnone),
		string(AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTargettargeted),
		string(AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTargettargetedAsEnrollmentRestrictions),
	}
}

func parseAndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget(input string) (*AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget, error) {
	vals := map[string]AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget{
		"all":                              AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTargetall,
		"none":                             AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTargetnone,
		"targeted":                         AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTargettargeted,
		"targetedasenrollmentrestrictions": AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTargettargetedAsEnrollmentRestrictions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAccountEnterpriseSettingsEnrollmentTarget(input)
	return &out, nil
}
