package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPowerState string

const (
	CloudPCPowerStatepoweredOff CloudPCPowerState = "PoweredOff"
	CloudPCPowerStaterunning    CloudPCPowerState = "Running"
)

func PossibleValuesForCloudPCPowerState() []string {
	return []string{
		string(CloudPCPowerStatepoweredOff),
		string(CloudPCPowerStaterunning),
	}
}

func parseCloudPCPowerState(input string) (*CloudPCPowerState, error) {
	vals := map[string]CloudPCPowerState{
		"poweredoff": CloudPCPowerStatepoweredOff,
		"running":    CloudPCPowerStaterunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPowerState(input)
	return &out, nil
}
