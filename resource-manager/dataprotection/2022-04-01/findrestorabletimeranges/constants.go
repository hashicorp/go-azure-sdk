package findrestorabletimeranges

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreSourceDataStoreType string

const (
	RestoreSourceDataStoreTypeArchiveStore     RestoreSourceDataStoreType = "ArchiveStore"
	RestoreSourceDataStoreTypeOperationalStore RestoreSourceDataStoreType = "OperationalStore"
	RestoreSourceDataStoreTypeVaultStore       RestoreSourceDataStoreType = "VaultStore"
)

func PossibleValuesForRestoreSourceDataStoreType() []string {
	return []string{
		string(RestoreSourceDataStoreTypeArchiveStore),
		string(RestoreSourceDataStoreTypeOperationalStore),
		string(RestoreSourceDataStoreTypeVaultStore),
	}
}
