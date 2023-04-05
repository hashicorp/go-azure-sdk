package storageinsights

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageInsightState string

const (
	StorageInsightStateERROR StorageInsightState = "ERROR"
	StorageInsightStateOK    StorageInsightState = "OK"
)

func PossibleValuesForStorageInsightState() []string {
	return []string{
		string(StorageInsightStateERROR),
		string(StorageInsightStateOK),
	}
}
