package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedMobileLobAppAppAvailability string

const (
	ManagedMobileLobAppAppAvailabilityglobal         ManagedMobileLobAppAppAvailability = "Global"
	ManagedMobileLobAppAppAvailabilitylineOfBusiness ManagedMobileLobAppAppAvailability = "LineOfBusiness"
)

func PossibleValuesForManagedMobileLobAppAppAvailability() []string {
	return []string{
		string(ManagedMobileLobAppAppAvailabilityglobal),
		string(ManagedMobileLobAppAppAvailabilitylineOfBusiness),
	}
}

func parseManagedMobileLobAppAppAvailability(input string) (*ManagedMobileLobAppAppAvailability, error) {
	vals := map[string]ManagedMobileLobAppAppAvailability{
		"global":         ManagedMobileLobAppAppAvailabilityglobal,
		"lineofbusiness": ManagedMobileLobAppAppAvailabilitylineOfBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedMobileLobAppAppAvailability(input)
	return &out, nil
}
