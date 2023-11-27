package managedinstanceoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpsertManagedServerOperationParameters struct {
	Family          *string `json:"family,omitempty"`
	StorageSizeInGB *int64  `json:"storageSizeInGB,omitempty"`
	Tier            *string `json:"tier,omitempty"`
	VCores          *int64  `json:"vCores,omitempty"`
}
