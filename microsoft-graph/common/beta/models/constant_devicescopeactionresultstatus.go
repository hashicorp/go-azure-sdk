package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceScopeActionResultStatus string

const (
	DeviceScopeActionResultStatusfailed    DeviceScopeActionResultStatus = "Failed"
	DeviceScopeActionResultStatussucceeded DeviceScopeActionResultStatus = "Succeeded"
)

func PossibleValuesForDeviceScopeActionResultStatus() []string {
	return []string{
		string(DeviceScopeActionResultStatusfailed),
		string(DeviceScopeActionResultStatussucceeded),
	}
}

func parseDeviceScopeActionResultStatus(input string) (*DeviceScopeActionResultStatus, error) {
	vals := map[string]DeviceScopeActionResultStatus{
		"failed":    DeviceScopeActionResultStatusfailed,
		"succeeded": DeviceScopeActionResultStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceScopeActionResultStatus(input)
	return &out, nil
}
