package autoprovisioningsettings

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoProvision string

const (
	AutoProvisionOff AutoProvision = "Off"
	AutoProvisionOn  AutoProvision = "On"
)

func PossibleValuesForAutoProvision() []string {
	return []string{
		string(AutoProvisionOff),
		string(AutoProvisionOn),
	}
}

func parseAutoProvision(input string) (*AutoProvision, error) {
	vals := map[string]AutoProvision{
		"off": AutoProvisionOff,
		"on":  AutoProvisionOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoProvision(input)
	return &out, nil
}
