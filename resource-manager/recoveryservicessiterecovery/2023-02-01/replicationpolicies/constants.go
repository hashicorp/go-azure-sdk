package replicationpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetMultiVMSyncStatus string

const (
	SetMultiVMSyncStatusDisable SetMultiVMSyncStatus = "Disable"
	SetMultiVMSyncStatusEnable  SetMultiVMSyncStatus = "Enable"
)

func PossibleValuesForSetMultiVMSyncStatus() []string {
	return []string{
		string(SetMultiVMSyncStatusDisable),
		string(SetMultiVMSyncStatusEnable),
	}
}
