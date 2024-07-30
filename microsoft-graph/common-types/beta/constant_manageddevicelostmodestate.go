package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceLostModeState string

const (
	ManagedDeviceLostModeState_Disabled ManagedDeviceLostModeState = "disabled"
	ManagedDeviceLostModeState_Enabled  ManagedDeviceLostModeState = "enabled"
)

func PossibleValuesForManagedDeviceLostModeState() []string {
	return []string{
		string(ManagedDeviceLostModeState_Disabled),
		string(ManagedDeviceLostModeState_Enabled),
	}
}

func (s *ManagedDeviceLostModeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceLostModeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceLostModeState(input string) (*ManagedDeviceLostModeState, error) {
	vals := map[string]ManagedDeviceLostModeState{
		"disabled": ManagedDeviceLostModeState_Disabled,
		"enabled":  ManagedDeviceLostModeState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceLostModeState(input)
	return &out, nil
}
