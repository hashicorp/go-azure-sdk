package storageaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataPolicy string

const (
	DataPolicyCloud DataPolicy = "Cloud"
	DataPolicyLocal DataPolicy = "Local"
)

func PossibleValuesForDataPolicy() []string {
	return []string{
		string(DataPolicyCloud),
		string(DataPolicyLocal),
	}
}

type StorageAccountStatus string

const (
	StorageAccountStatusNeedsAttention StorageAccountStatus = "NeedsAttention"
	StorageAccountStatusOK             StorageAccountStatus = "OK"
	StorageAccountStatusOffline        StorageAccountStatus = "Offline"
	StorageAccountStatusUnknown        StorageAccountStatus = "Unknown"
	StorageAccountStatusUpdating       StorageAccountStatus = "Updating"
)

func PossibleValuesForStorageAccountStatus() []string {
	return []string{
		string(StorageAccountStatusNeedsAttention),
		string(StorageAccountStatusOK),
		string(StorageAccountStatusOffline),
		string(StorageAccountStatusUnknown),
		string(StorageAccountStatusUpdating),
	}
}
