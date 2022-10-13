package recoverypoints

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointSyncType string

const (
	RecoveryPointSyncTypeMultiVmSyncRecoveryPoint RecoveryPointSyncType = "MultiVmSyncRecoveryPoint"
	RecoveryPointSyncTypePerVmRecoveryPoint       RecoveryPointSyncType = "PerVmRecoveryPoint"
)

func PossibleValuesForRecoveryPointSyncType() []string {
	return []string{
		string(RecoveryPointSyncTypeMultiVmSyncRecoveryPoint),
		string(RecoveryPointSyncTypePerVmRecoveryPoint),
	}
}

func parseRecoveryPointSyncType(input string) (*RecoveryPointSyncType, error) {
	vals := map[string]RecoveryPointSyncType{
		"multivmsyncrecoverypoint": RecoveryPointSyncTypeMultiVmSyncRecoveryPoint,
		"pervmrecoverypoint":       RecoveryPointSyncTypePerVmRecoveryPoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecoveryPointSyncType(input)
	return &out, nil
}
