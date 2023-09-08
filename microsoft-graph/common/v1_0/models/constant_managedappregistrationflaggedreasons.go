package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppRegistrationFlaggedReasons string

const (
	ManagedAppRegistrationFlaggedReasonsnone         ManagedAppRegistrationFlaggedReasons = "None"
	ManagedAppRegistrationFlaggedReasonsrootedDevice ManagedAppRegistrationFlaggedReasons = "RootedDevice"
)

func PossibleValuesForManagedAppRegistrationFlaggedReasons() []string {
	return []string{
		string(ManagedAppRegistrationFlaggedReasonsnone),
		string(ManagedAppRegistrationFlaggedReasonsrootedDevice),
	}
}

func parseManagedAppRegistrationFlaggedReasons(input string) (*ManagedAppRegistrationFlaggedReasons, error) {
	vals := map[string]ManagedAppRegistrationFlaggedReasons{
		"none":         ManagedAppRegistrationFlaggedReasonsnone,
		"rooteddevice": ManagedAppRegistrationFlaggedReasonsrootedDevice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppRegistrationFlaggedReasons(input)
	return &out, nil
}
