package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAndroidLobAppAppAvailability string

const (
	ManagedAndroidLobAppAppAvailabilityglobal         ManagedAndroidLobAppAppAvailability = "Global"
	ManagedAndroidLobAppAppAvailabilitylineOfBusiness ManagedAndroidLobAppAppAvailability = "LineOfBusiness"
)

func PossibleValuesForManagedAndroidLobAppAppAvailability() []string {
	return []string{
		string(ManagedAndroidLobAppAppAvailabilityglobal),
		string(ManagedAndroidLobAppAppAvailabilitylineOfBusiness),
	}
}

func parseManagedAndroidLobAppAppAvailability(input string) (*ManagedAndroidLobAppAppAvailability, error) {
	vals := map[string]ManagedAndroidLobAppAppAvailability{
		"global":         ManagedAndroidLobAppAppAvailabilityglobal,
		"lineofbusiness": ManagedAndroidLobAppAppAvailabilitylineOfBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAndroidLobAppAppAvailability(input)
	return &out, nil
}
