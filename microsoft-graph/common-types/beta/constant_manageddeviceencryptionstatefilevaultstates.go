package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceEncryptionStateFileVaultStates string

const (
	ManagedDeviceEncryptionStateFileVaultStates_DriveEncryptedByUser   ManagedDeviceEncryptionStateFileVaultStates = "driveEncryptedByUser"
	ManagedDeviceEncryptionStateFileVaultStates_EscrowNotEnabled       ManagedDeviceEncryptionStateFileVaultStates = "escrowNotEnabled"
	ManagedDeviceEncryptionStateFileVaultStates_Success                ManagedDeviceEncryptionStateFileVaultStates = "success"
	ManagedDeviceEncryptionStateFileVaultStates_UserDeferredEncryption ManagedDeviceEncryptionStateFileVaultStates = "userDeferredEncryption"
)

func PossibleValuesForManagedDeviceEncryptionStateFileVaultStates() []string {
	return []string{
		string(ManagedDeviceEncryptionStateFileVaultStates_DriveEncryptedByUser),
		string(ManagedDeviceEncryptionStateFileVaultStates_EscrowNotEnabled),
		string(ManagedDeviceEncryptionStateFileVaultStates_Success),
		string(ManagedDeviceEncryptionStateFileVaultStates_UserDeferredEncryption),
	}
}

func (s *ManagedDeviceEncryptionStateFileVaultStates) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceEncryptionStateFileVaultStates(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceEncryptionStateFileVaultStates(input string) (*ManagedDeviceEncryptionStateFileVaultStates, error) {
	vals := map[string]ManagedDeviceEncryptionStateFileVaultStates{
		"driveencryptedbyuser":   ManagedDeviceEncryptionStateFileVaultStates_DriveEncryptedByUser,
		"escrownotenabled":       ManagedDeviceEncryptionStateFileVaultStates_EscrowNotEnabled,
		"success":                ManagedDeviceEncryptionStateFileVaultStates_Success,
		"userdeferredencryption": ManagedDeviceEncryptionStateFileVaultStates_UserDeferredEncryption,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceEncryptionStateFileVaultStates(input)
	return &out, nil
}
