package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceEncryptionStateAdvancedBitLockerStates string

const (
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_FixedDriveEncryptionMethodMismatch      ManagedDeviceEncryptionStateAdvancedBitLockerStates = "fixedDriveEncryptionMethodMismatch"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_FixedDriveNotEncrypted                  ManagedDeviceEncryptionStateAdvancedBitLockerStates = "fixedDriveNotEncrypted"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_LoggedOnUserNonAdmin                    ManagedDeviceEncryptionStateAdvancedBitLockerStates = "loggedOnUserNonAdmin"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_NetworkError                            ManagedDeviceEncryptionStateAdvancedBitLockerStates = "networkError"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_NoUserConsent                           ManagedDeviceEncryptionStateAdvancedBitLockerStates = "noUserConsent"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeEncryptionMethodMismatch        ManagedDeviceEncryptionStateAdvancedBitLockerStates = "osVolumeEncryptionMethodMismatch"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmOnlyRequired                 ManagedDeviceEncryptionStateAdvancedBitLockerStates = "osVolumeTpmOnlyRequired"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmPinRequired                  ManagedDeviceEncryptionStateAdvancedBitLockerStates = "osVolumeTpmPinRequired"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmPinStartupKeyRequired        ManagedDeviceEncryptionStateAdvancedBitLockerStates = "osVolumeTpmPinStartupKeyRequired"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmRequired                     ManagedDeviceEncryptionStateAdvancedBitLockerStates = "osVolumeTpmRequired"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmStartupKeyRequired           ManagedDeviceEncryptionStateAdvancedBitLockerStates = "osVolumeTpmStartupKeyRequired"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeUnprotected                     ManagedDeviceEncryptionStateAdvancedBitLockerStates = "osVolumeUnprotected"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_RecoveryKeyBackupFailed                 ManagedDeviceEncryptionStateAdvancedBitLockerStates = "recoveryKeyBackupFailed"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_Success                                 ManagedDeviceEncryptionStateAdvancedBitLockerStates = "success"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_TpmNotAvailable                         ManagedDeviceEncryptionStateAdvancedBitLockerStates = "tpmNotAvailable"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_TpmNotReady                             ManagedDeviceEncryptionStateAdvancedBitLockerStates = "tpmNotReady"
	ManagedDeviceEncryptionStateAdvancedBitLockerStates_WindowsRecoveryEnvironmentNotConfigured ManagedDeviceEncryptionStateAdvancedBitLockerStates = "windowsRecoveryEnvironmentNotConfigured"
)

func PossibleValuesForManagedDeviceEncryptionStateAdvancedBitLockerStates() []string {
	return []string{
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_FixedDriveEncryptionMethodMismatch),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_FixedDriveNotEncrypted),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_LoggedOnUserNonAdmin),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_NetworkError),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_NoUserConsent),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeEncryptionMethodMismatch),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmOnlyRequired),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmPinRequired),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmPinStartupKeyRequired),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmRequired),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmStartupKeyRequired),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeUnprotected),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_RecoveryKeyBackupFailed),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_Success),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_TpmNotAvailable),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_TpmNotReady),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStates_WindowsRecoveryEnvironmentNotConfigured),
	}
}

func (s *ManagedDeviceEncryptionStateAdvancedBitLockerStates) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceEncryptionStateAdvancedBitLockerStates(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceEncryptionStateAdvancedBitLockerStates(input string) (*ManagedDeviceEncryptionStateAdvancedBitLockerStates, error) {
	vals := map[string]ManagedDeviceEncryptionStateAdvancedBitLockerStates{
		"fixeddriveencryptionmethodmismatch":      ManagedDeviceEncryptionStateAdvancedBitLockerStates_FixedDriveEncryptionMethodMismatch,
		"fixeddrivenotencrypted":                  ManagedDeviceEncryptionStateAdvancedBitLockerStates_FixedDriveNotEncrypted,
		"loggedonusernonadmin":                    ManagedDeviceEncryptionStateAdvancedBitLockerStates_LoggedOnUserNonAdmin,
		"networkerror":                            ManagedDeviceEncryptionStateAdvancedBitLockerStates_NetworkError,
		"nouserconsent":                           ManagedDeviceEncryptionStateAdvancedBitLockerStates_NoUserConsent,
		"osvolumeencryptionmethodmismatch":        ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeEncryptionMethodMismatch,
		"osvolumetpmonlyrequired":                 ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmOnlyRequired,
		"osvolumetpmpinrequired":                  ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmPinRequired,
		"osvolumetpmpinstartupkeyrequired":        ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmPinStartupKeyRequired,
		"osvolumetpmrequired":                     ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmRequired,
		"osvolumetpmstartupkeyrequired":           ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeTpmStartupKeyRequired,
		"osvolumeunprotected":                     ManagedDeviceEncryptionStateAdvancedBitLockerStates_OsVolumeUnprotected,
		"recoverykeybackupfailed":                 ManagedDeviceEncryptionStateAdvancedBitLockerStates_RecoveryKeyBackupFailed,
		"success":                                 ManagedDeviceEncryptionStateAdvancedBitLockerStates_Success,
		"tpmnotavailable":                         ManagedDeviceEncryptionStateAdvancedBitLockerStates_TpmNotAvailable,
		"tpmnotready":                             ManagedDeviceEncryptionStateAdvancedBitLockerStates_TpmNotReady,
		"windowsrecoveryenvironmentnotconfigured": ManagedDeviceEncryptionStateAdvancedBitLockerStates_WindowsRecoveryEnvironmentNotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceEncryptionStateAdvancedBitLockerStates(input)
	return &out, nil
}
