package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceLostModeState string

const (
	WindowsManagedDeviceLostModeState_Disabled WindowsManagedDeviceLostModeState = "disabled"
	WindowsManagedDeviceLostModeState_Enabled  WindowsManagedDeviceLostModeState = "enabled"
)

func PossibleValuesForWindowsManagedDeviceLostModeState() []string {
	return []string{
		string(WindowsManagedDeviceLostModeState_Disabled),
		string(WindowsManagedDeviceLostModeState_Enabled),
	}
}

func (s *WindowsManagedDeviceLostModeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceLostModeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceLostModeState(input string) (*WindowsManagedDeviceLostModeState, error) {
	vals := map[string]WindowsManagedDeviceLostModeState{
		"disabled": WindowsManagedDeviceLostModeState_Disabled,
		"enabled":  WindowsManagedDeviceLostModeState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceLostModeState(input)
	return &out, nil
}
