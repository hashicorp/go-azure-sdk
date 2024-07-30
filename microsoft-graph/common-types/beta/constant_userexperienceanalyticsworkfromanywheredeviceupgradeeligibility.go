package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility string

const (
	UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility_Capable    UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility = "capable"
	UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility_NotCapable UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility = "notCapable"
	UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility_Unknown    UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility = "unknown"
	UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility_Upgraded   UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility = "upgraded"
)

func PossibleValuesForUserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility() []string {
	return []string{
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility_Capable),
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility_NotCapable),
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility_Unknown),
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility_Upgraded),
	}
}

func (s *UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility(input string) (*UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility, error) {
	vals := map[string]UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility{
		"capable":    UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility_Capable,
		"notcapable": UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility_NotCapable,
		"unknown":    UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility_Unknown,
		"upgraded":   UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility_Upgraded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsWorkFromAnywhereDeviceUpgradeEligibility(input)
	return &out, nil
}
