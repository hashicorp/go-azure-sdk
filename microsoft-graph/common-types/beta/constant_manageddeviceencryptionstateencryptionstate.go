package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceEncryptionStateEncryptionState string

const (
	ManagedDeviceEncryptionStateEncryptionState_Encrypted    ManagedDeviceEncryptionStateEncryptionState = "encrypted"
	ManagedDeviceEncryptionStateEncryptionState_NotEncrypted ManagedDeviceEncryptionStateEncryptionState = "notEncrypted"
)

func PossibleValuesForManagedDeviceEncryptionStateEncryptionState() []string {
	return []string{
		string(ManagedDeviceEncryptionStateEncryptionState_Encrypted),
		string(ManagedDeviceEncryptionStateEncryptionState_NotEncrypted),
	}
}

func (s *ManagedDeviceEncryptionStateEncryptionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceEncryptionStateEncryptionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceEncryptionStateEncryptionState(input string) (*ManagedDeviceEncryptionStateEncryptionState, error) {
	vals := map[string]ManagedDeviceEncryptionStateEncryptionState{
		"encrypted":    ManagedDeviceEncryptionStateEncryptionState_Encrypted,
		"notencrypted": ManagedDeviceEncryptionStateEncryptionState_NotEncrypted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceEncryptionStateEncryptionState(input)
	return &out, nil
}
