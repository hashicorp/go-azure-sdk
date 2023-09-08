package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppRegistrationFlaggedReasons string

const (
	IosManagedAppRegistrationFlaggedReasonsnone         IosManagedAppRegistrationFlaggedReasons = "None"
	IosManagedAppRegistrationFlaggedReasonsrootedDevice IosManagedAppRegistrationFlaggedReasons = "RootedDevice"
)

func PossibleValuesForIosManagedAppRegistrationFlaggedReasons() []string {
	return []string{
		string(IosManagedAppRegistrationFlaggedReasonsnone),
		string(IosManagedAppRegistrationFlaggedReasonsrootedDevice),
	}
}

func parseIosManagedAppRegistrationFlaggedReasons(input string) (*IosManagedAppRegistrationFlaggedReasons, error) {
	vals := map[string]IosManagedAppRegistrationFlaggedReasons{
		"none":         IosManagedAppRegistrationFlaggedReasonsnone,
		"rooteddevice": IosManagedAppRegistrationFlaggedReasonsrootedDevice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppRegistrationFlaggedReasons(input)
	return &out, nil
}
