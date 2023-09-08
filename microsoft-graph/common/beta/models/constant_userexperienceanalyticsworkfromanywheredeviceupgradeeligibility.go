package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility string

const (
	UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibilitycapable    UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility = "Capable"
	UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibilitynotCapable UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility = "NotCapable"
	UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibilityunknown    UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility = "Unknown"
	UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibilityupgraded   UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility = "Upgraded"
)

func PossibleValuesForUserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility() []string {
	return []string{
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibilitycapable),
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibilitynotCapable),
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibilityunknown),
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibilityupgraded),
	}
}

func parseUserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility(input string) (*UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility, error) {
	vals := map[string]UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility{
		"capable":    UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibilitycapable,
		"notcapable": UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibilitynotCapable,
		"unknown":    UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibilityunknown,
		"upgraded":   UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibilityupgraded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility(input)
	return &out, nil
}
