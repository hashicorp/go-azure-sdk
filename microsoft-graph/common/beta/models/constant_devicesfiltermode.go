package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DevicesFilterMode string

const (
	DevicesFilterModeallowed DevicesFilterMode = "Allowed"
	DevicesFilterModeblocked DevicesFilterMode = "Blocked"
)

func PossibleValuesForDevicesFilterMode() []string {
	return []string{
		string(DevicesFilterModeallowed),
		string(DevicesFilterModeblocked),
	}
}

func parseDevicesFilterMode(input string) (*DevicesFilterMode, error) {
	vals := map[string]DevicesFilterMode{
		"allowed": DevicesFilterModeallowed,
		"blocked": DevicesFilterModeblocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DevicesFilterMode(input)
	return &out, nil
}
