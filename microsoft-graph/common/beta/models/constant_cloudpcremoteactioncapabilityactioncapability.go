package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRemoteActionCapabilityActionCapability string

const (
	CloudPCRemoteActionCapabilityActionCapabilitydisabled CloudPCRemoteActionCapabilityActionCapability = "Disabled"
	CloudPCRemoteActionCapabilityActionCapabilityenabled  CloudPCRemoteActionCapabilityActionCapability = "Enabled"
)

func PossibleValuesForCloudPCRemoteActionCapabilityActionCapability() []string {
	return []string{
		string(CloudPCRemoteActionCapabilityActionCapabilitydisabled),
		string(CloudPCRemoteActionCapabilityActionCapabilityenabled),
	}
}

func parseCloudPCRemoteActionCapabilityActionCapability(input string) (*CloudPCRemoteActionCapabilityActionCapability, error) {
	vals := map[string]CloudPCRemoteActionCapabilityActionCapability{
		"disabled": CloudPCRemoteActionCapabilityActionCapabilitydisabled,
		"enabled":  CloudPCRemoteActionCapabilityActionCapabilityenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCRemoteActionCapabilityActionCapability(input)
	return &out, nil
}
