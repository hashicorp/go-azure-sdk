package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppRegistrationFlaggedReasons string

const (
	ManagedAppRegistrationFlaggedReasonsandroidBootloaderUnlocked ManagedAppRegistrationFlaggedReasons = "AndroidBootloaderUnlocked"
	ManagedAppRegistrationFlaggedReasonsandroidFactoryRomModified ManagedAppRegistrationFlaggedReasons = "AndroidFactoryRomModified"
	ManagedAppRegistrationFlaggedReasonsnone                      ManagedAppRegistrationFlaggedReasons = "None"
	ManagedAppRegistrationFlaggedReasonsrootedDevice              ManagedAppRegistrationFlaggedReasons = "RootedDevice"
)

func PossibleValuesForManagedAppRegistrationFlaggedReasons() []string {
	return []string{
		string(ManagedAppRegistrationFlaggedReasonsandroidBootloaderUnlocked),
		string(ManagedAppRegistrationFlaggedReasonsandroidFactoryRomModified),
		string(ManagedAppRegistrationFlaggedReasonsnone),
		string(ManagedAppRegistrationFlaggedReasonsrootedDevice),
	}
}

func parseManagedAppRegistrationFlaggedReasons(input string) (*ManagedAppRegistrationFlaggedReasons, error) {
	vals := map[string]ManagedAppRegistrationFlaggedReasons{
		"androidbootloaderunlocked": ManagedAppRegistrationFlaggedReasonsandroidBootloaderUnlocked,
		"androidfactoryrommodified": ManagedAppRegistrationFlaggedReasonsandroidFactoryRomModified,
		"none":                      ManagedAppRegistrationFlaggedReasonsnone,
		"rooteddevice":              ManagedAppRegistrationFlaggedReasonsrootedDevice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppRegistrationFlaggedReasons(input)
	return &out, nil
}
