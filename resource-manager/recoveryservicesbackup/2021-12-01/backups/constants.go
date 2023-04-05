package backups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupType string

const (
	BackupTypeCopyOnlyFull BackupType = "CopyOnlyFull"
	BackupTypeDifferential BackupType = "Differential"
	BackupTypeFull         BackupType = "Full"
	BackupTypeIncremental  BackupType = "Incremental"
	BackupTypeInvalid      BackupType = "Invalid"
	BackupTypeLog          BackupType = "Log"
)

func PossibleValuesForBackupType() []string {
	return []string{
		string(BackupTypeCopyOnlyFull),
		string(BackupTypeDifferential),
		string(BackupTypeFull),
		string(BackupTypeIncremental),
		string(BackupTypeInvalid),
		string(BackupTypeLog),
	}
}
