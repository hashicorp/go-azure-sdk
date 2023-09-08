package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppRegistrationFlaggedReasons string

const (
	AndroidManagedAppRegistrationFlaggedReasonsnone         AndroidManagedAppRegistrationFlaggedReasons = "None"
	AndroidManagedAppRegistrationFlaggedReasonsrootedDevice AndroidManagedAppRegistrationFlaggedReasons = "RootedDevice"
)

func PossibleValuesForAndroidManagedAppRegistrationFlaggedReasons() []string {
	return []string{
		string(AndroidManagedAppRegistrationFlaggedReasonsnone),
		string(AndroidManagedAppRegistrationFlaggedReasonsrootedDevice),
	}
}

func parseAndroidManagedAppRegistrationFlaggedReasons(input string) (*AndroidManagedAppRegistrationFlaggedReasons, error) {
	vals := map[string]AndroidManagedAppRegistrationFlaggedReasons{
		"none":         AndroidManagedAppRegistrationFlaggedReasonsnone,
		"rooteddevice": AndroidManagedAppRegistrationFlaggedReasonsrootedDevice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppRegistrationFlaggedReasons(input)
	return &out, nil
}
