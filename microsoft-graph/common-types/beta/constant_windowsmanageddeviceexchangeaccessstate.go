package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceExchangeAccessState string

const (
	WindowsManagedDeviceExchangeAccessState_Allowed     WindowsManagedDeviceExchangeAccessState = "allowed"
	WindowsManagedDeviceExchangeAccessState_Blocked     WindowsManagedDeviceExchangeAccessState = "blocked"
	WindowsManagedDeviceExchangeAccessState_None        WindowsManagedDeviceExchangeAccessState = "none"
	WindowsManagedDeviceExchangeAccessState_Quarantined WindowsManagedDeviceExchangeAccessState = "quarantined"
	WindowsManagedDeviceExchangeAccessState_Unknown     WindowsManagedDeviceExchangeAccessState = "unknown"
)

func PossibleValuesForWindowsManagedDeviceExchangeAccessState() []string {
	return []string{
		string(WindowsManagedDeviceExchangeAccessState_Allowed),
		string(WindowsManagedDeviceExchangeAccessState_Blocked),
		string(WindowsManagedDeviceExchangeAccessState_None),
		string(WindowsManagedDeviceExchangeAccessState_Quarantined),
		string(WindowsManagedDeviceExchangeAccessState_Unknown),
	}
}

func (s *WindowsManagedDeviceExchangeAccessState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceExchangeAccessState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceExchangeAccessState(input string) (*WindowsManagedDeviceExchangeAccessState, error) {
	vals := map[string]WindowsManagedDeviceExchangeAccessState{
		"allowed":     WindowsManagedDeviceExchangeAccessState_Allowed,
		"blocked":     WindowsManagedDeviceExchangeAccessState_Blocked,
		"none":        WindowsManagedDeviceExchangeAccessState_None,
		"quarantined": WindowsManagedDeviceExchangeAccessState_Quarantined,
		"unknown":     WindowsManagedDeviceExchangeAccessState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceExchangeAccessState(input)
	return &out, nil
}
