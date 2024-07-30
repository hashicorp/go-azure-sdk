package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementPartnerPartnerState string

const (
	DeviceManagementPartnerPartnerState_Enabled      DeviceManagementPartnerPartnerState = "enabled"
	DeviceManagementPartnerPartnerState_Rejected     DeviceManagementPartnerPartnerState = "rejected"
	DeviceManagementPartnerPartnerState_Terminated   DeviceManagementPartnerPartnerState = "terminated"
	DeviceManagementPartnerPartnerState_Unavailable  DeviceManagementPartnerPartnerState = "unavailable"
	DeviceManagementPartnerPartnerState_Unknown      DeviceManagementPartnerPartnerState = "unknown"
	DeviceManagementPartnerPartnerState_Unresponsive DeviceManagementPartnerPartnerState = "unresponsive"
)

func PossibleValuesForDeviceManagementPartnerPartnerState() []string {
	return []string{
		string(DeviceManagementPartnerPartnerState_Enabled),
		string(DeviceManagementPartnerPartnerState_Rejected),
		string(DeviceManagementPartnerPartnerState_Terminated),
		string(DeviceManagementPartnerPartnerState_Unavailable),
		string(DeviceManagementPartnerPartnerState_Unknown),
		string(DeviceManagementPartnerPartnerState_Unresponsive),
	}
}

func (s *DeviceManagementPartnerPartnerState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementPartnerPartnerState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementPartnerPartnerState(input string) (*DeviceManagementPartnerPartnerState, error) {
	vals := map[string]DeviceManagementPartnerPartnerState{
		"enabled":      DeviceManagementPartnerPartnerState_Enabled,
		"rejected":     DeviceManagementPartnerPartnerState_Rejected,
		"terminated":   DeviceManagementPartnerPartnerState_Terminated,
		"unavailable":  DeviceManagementPartnerPartnerState_Unavailable,
		"unknown":      DeviceManagementPartnerPartnerState_Unknown,
		"unresponsive": DeviceManagementPartnerPartnerState_Unresponsive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementPartnerPartnerState(input)
	return &out, nil
}
