package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIOSStoreAppAppAvailability string

const (
	ManagedIOSStoreAppAppAvailabilityglobal         ManagedIOSStoreAppAppAvailability = "Global"
	ManagedIOSStoreAppAppAvailabilitylineOfBusiness ManagedIOSStoreAppAppAvailability = "LineOfBusiness"
)

func PossibleValuesForManagedIOSStoreAppAppAvailability() []string {
	return []string{
		string(ManagedIOSStoreAppAppAvailabilityglobal),
		string(ManagedIOSStoreAppAppAvailabilitylineOfBusiness),
	}
}

func parseManagedIOSStoreAppAppAvailability(input string) (*ManagedIOSStoreAppAppAvailability, error) {
	vals := map[string]ManagedIOSStoreAppAppAvailability{
		"global":         ManagedIOSStoreAppAppAvailabilityglobal,
		"lineofbusiness": ManagedIOSStoreAppAppAvailabilitylineOfBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedIOSStoreAppAppAvailability(input)
	return &out, nil
}
