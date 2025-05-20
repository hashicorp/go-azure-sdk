package returnsrestoreoperationstatus

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupRestoreOperationStatus string

const (
	BackupRestoreOperationStatusCancelled  BackupRestoreOperationStatus = "Cancelled"
	BackupRestoreOperationStatusFailed     BackupRestoreOperationStatus = "Failed"
	BackupRestoreOperationStatusInProgress BackupRestoreOperationStatus = "InProgress"
	BackupRestoreOperationStatusSucceeded  BackupRestoreOperationStatus = "Succeeded"
)

func PossibleValuesForBackupRestoreOperationStatus() []string {
	return []string{
		string(BackupRestoreOperationStatusCancelled),
		string(BackupRestoreOperationStatusFailed),
		string(BackupRestoreOperationStatusInProgress),
		string(BackupRestoreOperationStatusSucceeded),
	}
}

func (s *BackupRestoreOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBackupRestoreOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBackupRestoreOperationStatus(input string) (*BackupRestoreOperationStatus, error) {
	vals := map[string]BackupRestoreOperationStatus{
		"cancelled":  BackupRestoreOperationStatusCancelled,
		"failed":     BackupRestoreOperationStatusFailed,
		"inprogress": BackupRestoreOperationStatusInProgress,
		"succeeded":  BackupRestoreOperationStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupRestoreOperationStatus(input)
	return &out, nil
}
