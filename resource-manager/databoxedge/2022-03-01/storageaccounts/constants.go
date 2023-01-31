package storageaccounts

import "strings"

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

func parseDataPolicy(input string) (*DataPolicy, error) {
	vals := map[string]DataPolicy{
		"cloud": DataPolicyCloud,
		"local": DataPolicyLocal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataPolicy(input)
	return &out, nil
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

func parseStorageAccountStatus(input string) (*StorageAccountStatus, error) {
	vals := map[string]StorageAccountStatus{
		"needsattention": StorageAccountStatusNeedsAttention,
		"ok":             StorageAccountStatusOK,
		"offline":        StorageAccountStatusOffline,
		"unknown":        StorageAccountStatusUnknown,
		"updating":       StorageAccountStatusUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageAccountStatus(input)
	return &out, nil
}
