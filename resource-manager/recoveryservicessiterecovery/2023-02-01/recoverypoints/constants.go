package recoverypoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointSyncType string

const (
	RecoveryPointSyncTypeMultiVMSyncRecoveryPoint RecoveryPointSyncType = "MultiVmSyncRecoveryPoint"
	RecoveryPointSyncTypePerVMRecoveryPoint       RecoveryPointSyncType = "PerVmRecoveryPoint"
)

func PossibleValuesForRecoveryPointSyncType() []string {
	return []string{
		string(RecoveryPointSyncTypeMultiVMSyncRecoveryPoint),
		string(RecoveryPointSyncTypePerVMRecoveryPoint),
	}
}
