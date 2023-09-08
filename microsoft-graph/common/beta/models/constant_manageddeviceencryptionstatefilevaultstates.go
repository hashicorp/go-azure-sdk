package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceEncryptionStateFileVaultStates string

const (
	ManagedDeviceEncryptionStateFileVaultStatesdriveEncryptedByUser   ManagedDeviceEncryptionStateFileVaultStates = "DriveEncryptedByUser"
	ManagedDeviceEncryptionStateFileVaultStatesescrowNotEnabled       ManagedDeviceEncryptionStateFileVaultStates = "EscrowNotEnabled"
	ManagedDeviceEncryptionStateFileVaultStatessuccess                ManagedDeviceEncryptionStateFileVaultStates = "Success"
	ManagedDeviceEncryptionStateFileVaultStatesuserDeferredEncryption ManagedDeviceEncryptionStateFileVaultStates = "UserDeferredEncryption"
)

func PossibleValuesForManagedDeviceEncryptionStateFileVaultStates() []string {
	return []string{
		string(ManagedDeviceEncryptionStateFileVaultStatesdriveEncryptedByUser),
		string(ManagedDeviceEncryptionStateFileVaultStatesescrowNotEnabled),
		string(ManagedDeviceEncryptionStateFileVaultStatessuccess),
		string(ManagedDeviceEncryptionStateFileVaultStatesuserDeferredEncryption),
	}
}

func parseManagedDeviceEncryptionStateFileVaultStates(input string) (*ManagedDeviceEncryptionStateFileVaultStates, error) {
	vals := map[string]ManagedDeviceEncryptionStateFileVaultStates{
		"driveencryptedbyuser":   ManagedDeviceEncryptionStateFileVaultStatesdriveEncryptedByUser,
		"escrownotenabled":       ManagedDeviceEncryptionStateFileVaultStatesescrowNotEnabled,
		"success":                ManagedDeviceEncryptionStateFileVaultStatessuccess,
		"userdeferredencryption": ManagedDeviceEncryptionStateFileVaultStatesuserDeferredEncryption,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceEncryptionStateFileVaultStates(input)
	return &out, nil
}
