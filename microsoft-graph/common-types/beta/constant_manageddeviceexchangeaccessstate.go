package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceExchangeAccessState string

const (
	ManagedDeviceExchangeAccessState_Allowed     ManagedDeviceExchangeAccessState = "allowed"
	ManagedDeviceExchangeAccessState_Blocked     ManagedDeviceExchangeAccessState = "blocked"
	ManagedDeviceExchangeAccessState_None        ManagedDeviceExchangeAccessState = "none"
	ManagedDeviceExchangeAccessState_Quarantined ManagedDeviceExchangeAccessState = "quarantined"
	ManagedDeviceExchangeAccessState_Unknown     ManagedDeviceExchangeAccessState = "unknown"
)

func PossibleValuesForManagedDeviceExchangeAccessState() []string {
	return []string{
		string(ManagedDeviceExchangeAccessState_Allowed),
		string(ManagedDeviceExchangeAccessState_Blocked),
		string(ManagedDeviceExchangeAccessState_None),
		string(ManagedDeviceExchangeAccessState_Quarantined),
		string(ManagedDeviceExchangeAccessState_Unknown),
	}
}

func (s *ManagedDeviceExchangeAccessState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceExchangeAccessState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceExchangeAccessState(input string) (*ManagedDeviceExchangeAccessState, error) {
	vals := map[string]ManagedDeviceExchangeAccessState{
		"allowed":     ManagedDeviceExchangeAccessState_Allowed,
		"blocked":     ManagedDeviceExchangeAccessState_Blocked,
		"none":        ManagedDeviceExchangeAccessState_None,
		"quarantined": ManagedDeviceExchangeAccessState_Quarantined,
		"unknown":     ManagedDeviceExchangeAccessState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceExchangeAccessState(input)
	return &out, nil
}
