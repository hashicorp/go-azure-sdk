package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAndroidStoreAppAppAvailability string

const (
	ManagedAndroidStoreAppAppAvailabilityglobal         ManagedAndroidStoreAppAppAvailability = "Global"
	ManagedAndroidStoreAppAppAvailabilitylineOfBusiness ManagedAndroidStoreAppAppAvailability = "LineOfBusiness"
)

func PossibleValuesForManagedAndroidStoreAppAppAvailability() []string {
	return []string{
		string(ManagedAndroidStoreAppAppAvailabilityglobal),
		string(ManagedAndroidStoreAppAppAvailabilitylineOfBusiness),
	}
}

func parseManagedAndroidStoreAppAppAvailability(input string) (*ManagedAndroidStoreAppAppAvailability, error) {
	vals := map[string]ManagedAndroidStoreAppAppAvailability{
		"global":         ManagedAndroidStoreAppAppAvailabilityglobal,
		"lineofbusiness": ManagedAndroidStoreAppAppAvailabilitylineOfBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAndroidStoreAppAppAvailability(input)
	return &out, nil
}
