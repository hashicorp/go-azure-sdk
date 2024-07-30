package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceActionResultActionState string

const (
	DeviceActionResultActionState_Active       DeviceActionResultActionState = "active"
	DeviceActionResultActionState_Canceled     DeviceActionResultActionState = "canceled"
	DeviceActionResultActionState_Done         DeviceActionResultActionState = "done"
	DeviceActionResultActionState_Failed       DeviceActionResultActionState = "failed"
	DeviceActionResultActionState_None         DeviceActionResultActionState = "none"
	DeviceActionResultActionState_NotSupported DeviceActionResultActionState = "notSupported"
	DeviceActionResultActionState_Pending      DeviceActionResultActionState = "pending"
)

func PossibleValuesForDeviceActionResultActionState() []string {
	return []string{
		string(DeviceActionResultActionState_Active),
		string(DeviceActionResultActionState_Canceled),
		string(DeviceActionResultActionState_Done),
		string(DeviceActionResultActionState_Failed),
		string(DeviceActionResultActionState_None),
		string(DeviceActionResultActionState_NotSupported),
		string(DeviceActionResultActionState_Pending),
	}
}

func (s *DeviceActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceActionResultActionState(input string) (*DeviceActionResultActionState, error) {
	vals := map[string]DeviceActionResultActionState{
		"active":       DeviceActionResultActionState_Active,
		"canceled":     DeviceActionResultActionState_Canceled,
		"done":         DeviceActionResultActionState_Done,
		"failed":       DeviceActionResultActionState_Failed,
		"none":         DeviceActionResultActionState_None,
		"notsupported": DeviceActionResultActionState_NotSupported,
		"pending":      DeviceActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceActionResultActionState(input)
	return &out, nil
}
