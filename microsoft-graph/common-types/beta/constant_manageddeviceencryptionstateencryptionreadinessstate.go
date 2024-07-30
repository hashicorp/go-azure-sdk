package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceEncryptionStateEncryptionReadinessState string

const (
	ManagedDeviceEncryptionStateEncryptionReadinessState_NotReady ManagedDeviceEncryptionStateEncryptionReadinessState = "notReady"
	ManagedDeviceEncryptionStateEncryptionReadinessState_Ready    ManagedDeviceEncryptionStateEncryptionReadinessState = "ready"
)

func PossibleValuesForManagedDeviceEncryptionStateEncryptionReadinessState() []string {
	return []string{
		string(ManagedDeviceEncryptionStateEncryptionReadinessState_NotReady),
		string(ManagedDeviceEncryptionStateEncryptionReadinessState_Ready),
	}
}

func (s *ManagedDeviceEncryptionStateEncryptionReadinessState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceEncryptionStateEncryptionReadinessState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceEncryptionStateEncryptionReadinessState(input string) (*ManagedDeviceEncryptionStateEncryptionReadinessState, error) {
	vals := map[string]ManagedDeviceEncryptionStateEncryptionReadinessState{
		"notready": ManagedDeviceEncryptionStateEncryptionReadinessState_NotReady,
		"ready":    ManagedDeviceEncryptionStateEncryptionReadinessState_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceEncryptionStateEncryptionReadinessState(input)
	return &out, nil
}
