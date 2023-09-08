package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementPartnerPartnerState string

const (
	DeviceManagementPartnerPartnerStateenabled      DeviceManagementPartnerPartnerState = "Enabled"
	DeviceManagementPartnerPartnerStaterejected     DeviceManagementPartnerPartnerState = "Rejected"
	DeviceManagementPartnerPartnerStateterminated   DeviceManagementPartnerPartnerState = "Terminated"
	DeviceManagementPartnerPartnerStateunavailable  DeviceManagementPartnerPartnerState = "Unavailable"
	DeviceManagementPartnerPartnerStateunknown      DeviceManagementPartnerPartnerState = "Unknown"
	DeviceManagementPartnerPartnerStateunresponsive DeviceManagementPartnerPartnerState = "Unresponsive"
)

func PossibleValuesForDeviceManagementPartnerPartnerState() []string {
	return []string{
		string(DeviceManagementPartnerPartnerStateenabled),
		string(DeviceManagementPartnerPartnerStaterejected),
		string(DeviceManagementPartnerPartnerStateterminated),
		string(DeviceManagementPartnerPartnerStateunavailable),
		string(DeviceManagementPartnerPartnerStateunknown),
		string(DeviceManagementPartnerPartnerStateunresponsive),
	}
}

func parseDeviceManagementPartnerPartnerState(input string) (*DeviceManagementPartnerPartnerState, error) {
	vals := map[string]DeviceManagementPartnerPartnerState{
		"enabled":      DeviceManagementPartnerPartnerStateenabled,
		"rejected":     DeviceManagementPartnerPartnerStaterejected,
		"terminated":   DeviceManagementPartnerPartnerStateterminated,
		"unavailable":  DeviceManagementPartnerPartnerStateunavailable,
		"unknown":      DeviceManagementPartnerPartnerStateunknown,
		"unresponsive": DeviceManagementPartnerPartnerStateunresponsive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementPartnerPartnerState(input)
	return &out, nil
}
