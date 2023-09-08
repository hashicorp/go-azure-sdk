package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkMobileAppConfigurationProfileApplicability string

const (
	AndroidForWorkMobileAppConfigurationProfileApplicabilityandroidDeviceOwner AndroidForWorkMobileAppConfigurationProfileApplicability = "AndroidDeviceOwner"
	AndroidForWorkMobileAppConfigurationProfileApplicabilityandroidWorkProfile AndroidForWorkMobileAppConfigurationProfileApplicability = "AndroidWorkProfile"
	AndroidForWorkMobileAppConfigurationProfileApplicabilitydefault            AndroidForWorkMobileAppConfigurationProfileApplicability = "Default"
)

func PossibleValuesForAndroidForWorkMobileAppConfigurationProfileApplicability() []string {
	return []string{
		string(AndroidForWorkMobileAppConfigurationProfileApplicabilityandroidDeviceOwner),
		string(AndroidForWorkMobileAppConfigurationProfileApplicabilityandroidWorkProfile),
		string(AndroidForWorkMobileAppConfigurationProfileApplicabilitydefault),
	}
}

func parseAndroidForWorkMobileAppConfigurationProfileApplicability(input string) (*AndroidForWorkMobileAppConfigurationProfileApplicability, error) {
	vals := map[string]AndroidForWorkMobileAppConfigurationProfileApplicability{
		"androiddeviceowner": AndroidForWorkMobileAppConfigurationProfileApplicabilityandroidDeviceOwner,
		"androidworkprofile": AndroidForWorkMobileAppConfigurationProfileApplicabilityandroidWorkProfile,
		"default":            AndroidForWorkMobileAppConfigurationProfileApplicabilitydefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkMobileAppConfigurationProfileApplicability(input)
	return &out, nil
}
