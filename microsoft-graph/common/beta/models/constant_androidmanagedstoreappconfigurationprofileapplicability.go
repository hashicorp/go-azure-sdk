package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppConfigurationProfileApplicability string

const (
	AndroidManagedStoreAppConfigurationProfileApplicabilityandroidDeviceOwner AndroidManagedStoreAppConfigurationProfileApplicability = "AndroidDeviceOwner"
	AndroidManagedStoreAppConfigurationProfileApplicabilityandroidWorkProfile AndroidManagedStoreAppConfigurationProfileApplicability = "AndroidWorkProfile"
	AndroidManagedStoreAppConfigurationProfileApplicabilitydefault            AndroidManagedStoreAppConfigurationProfileApplicability = "Default"
)

func PossibleValuesForAndroidManagedStoreAppConfigurationProfileApplicability() []string {
	return []string{
		string(AndroidManagedStoreAppConfigurationProfileApplicabilityandroidDeviceOwner),
		string(AndroidManagedStoreAppConfigurationProfileApplicabilityandroidWorkProfile),
		string(AndroidManagedStoreAppConfigurationProfileApplicabilitydefault),
	}
}

func parseAndroidManagedStoreAppConfigurationProfileApplicability(input string) (*AndroidManagedStoreAppConfigurationProfileApplicability, error) {
	vals := map[string]AndroidManagedStoreAppConfigurationProfileApplicability{
		"androiddeviceowner": AndroidManagedStoreAppConfigurationProfileApplicabilityandroidDeviceOwner,
		"androidworkprofile": AndroidManagedStoreAppConfigurationProfileApplicabilityandroidWorkProfile,
		"default":            AndroidManagedStoreAppConfigurationProfileApplicabilitydefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAppConfigurationProfileApplicability(input)
	return &out, nil
}
