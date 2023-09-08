package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceEncryptionStateAdvancedBitLockerStates string

const (
	ManagedDeviceEncryptionStateAdvancedBitLockerStatesfixedDriveEncryptionMethodMismatch      ManagedDeviceEncryptionStateAdvancedBitLockerStates = "FixedDriveEncryptionMethodMismatch"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatesfixedDriveNotEncrypted                  ManagedDeviceEncryptionStateAdvancedBitLockerStates = "FixedDriveNotEncrypted"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatesloggedOnUserNonAdmin                    ManagedDeviceEncryptionStateAdvancedBitLockerStates = "LoggedOnUserNonAdmin"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatesnetworkError                            ManagedDeviceEncryptionStateAdvancedBitLockerStates = "NetworkError"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatesnoUserConsent                           ManagedDeviceEncryptionStateAdvancedBitLockerStates = "NoUserConsent"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeEncryptionMethodMismatch        ManagedDeviceEncryptionStateAdvancedBitLockerStates = "OsVolumeEncryptionMethodMismatch"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmOnlyRequired                 ManagedDeviceEncryptionStateAdvancedBitLockerStates = "OsVolumeTpmOnlyRequired"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmPinRequired                  ManagedDeviceEncryptionStateAdvancedBitLockerStates = "OsVolumeTpmPinRequired"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmPinStartupKeyRequired        ManagedDeviceEncryptionStateAdvancedBitLockerStates = "OsVolumeTpmPinStartupKeyRequired"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmRequired                     ManagedDeviceEncryptionStateAdvancedBitLockerStates = "OsVolumeTpmRequired"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmStartupKeyRequired           ManagedDeviceEncryptionStateAdvancedBitLockerStates = "OsVolumeTpmStartupKeyRequired"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeUnprotected                     ManagedDeviceEncryptionStateAdvancedBitLockerStates = "OsVolumeUnprotected"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatesrecoveryKeyBackupFailed                 ManagedDeviceEncryptionStateAdvancedBitLockerStates = "RecoveryKeyBackupFailed"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatessuccess                                 ManagedDeviceEncryptionStateAdvancedBitLockerStates = "Success"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatestpmNotAvailable                         ManagedDeviceEncryptionStateAdvancedBitLockerStates = "TpmNotAvailable"
	ManagedDeviceEncryptionStateAdvancedBitLockerStatestpmNotReady                             ManagedDeviceEncryptionStateAdvancedBitLockerStates = "TpmNotReady"
	ManagedDeviceEncryptionStateAdvancedBitLockerStateswindowsRecoveryEnvironmentNotConfigured ManagedDeviceEncryptionStateAdvancedBitLockerStates = "WindowsRecoveryEnvironmentNotConfigured"
)

func PossibleValuesForManagedDeviceEncryptionStateAdvancedBitLockerStates() []string {
	return []string{
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatesfixedDriveEncryptionMethodMismatch),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatesfixedDriveNotEncrypted),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatesloggedOnUserNonAdmin),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatesnetworkError),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatesnoUserConsent),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeEncryptionMethodMismatch),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmOnlyRequired),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmPinRequired),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmPinStartupKeyRequired),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmRequired),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmStartupKeyRequired),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeUnprotected),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatesrecoveryKeyBackupFailed),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatessuccess),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatestpmNotAvailable),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStatestpmNotReady),
		string(ManagedDeviceEncryptionStateAdvancedBitLockerStateswindowsRecoveryEnvironmentNotConfigured),
	}
}

func parseManagedDeviceEncryptionStateAdvancedBitLockerStates(input string) (*ManagedDeviceEncryptionStateAdvancedBitLockerStates, error) {
	vals := map[string]ManagedDeviceEncryptionStateAdvancedBitLockerStates{
		"fixeddriveencryptionmethodmismatch":      ManagedDeviceEncryptionStateAdvancedBitLockerStatesfixedDriveEncryptionMethodMismatch,
		"fixeddrivenotencrypted":                  ManagedDeviceEncryptionStateAdvancedBitLockerStatesfixedDriveNotEncrypted,
		"loggedonusernonadmin":                    ManagedDeviceEncryptionStateAdvancedBitLockerStatesloggedOnUserNonAdmin,
		"networkerror":                            ManagedDeviceEncryptionStateAdvancedBitLockerStatesnetworkError,
		"nouserconsent":                           ManagedDeviceEncryptionStateAdvancedBitLockerStatesnoUserConsent,
		"osvolumeencryptionmethodmismatch":        ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeEncryptionMethodMismatch,
		"osvolumetpmonlyrequired":                 ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmOnlyRequired,
		"osvolumetpmpinrequired":                  ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmPinRequired,
		"osvolumetpmpinstartupkeyrequired":        ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmPinStartupKeyRequired,
		"osvolumetpmrequired":                     ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmRequired,
		"osvolumetpmstartupkeyrequired":           ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeTpmStartupKeyRequired,
		"osvolumeunprotected":                     ManagedDeviceEncryptionStateAdvancedBitLockerStatesosVolumeUnprotected,
		"recoverykeybackupfailed":                 ManagedDeviceEncryptionStateAdvancedBitLockerStatesrecoveryKeyBackupFailed,
		"success":                                 ManagedDeviceEncryptionStateAdvancedBitLockerStatessuccess,
		"tpmnotavailable":                         ManagedDeviceEncryptionStateAdvancedBitLockerStatestpmNotAvailable,
		"tpmnotready":                             ManagedDeviceEncryptionStateAdvancedBitLockerStatestpmNotReady,
		"windowsrecoveryenvironmentnotconfigured": ManagedDeviceEncryptionStateAdvancedBitLockerStateswindowsRecoveryEnvironmentNotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceEncryptionStateAdvancedBitLockerStates(input)
	return &out, nil
}
