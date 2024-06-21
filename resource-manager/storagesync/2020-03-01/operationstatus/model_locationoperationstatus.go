package operationstatus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationOperationStatus struct {
	EndTime         *string              `json:"endTime,omitempty"`
	Error           *StorageSyncApiError `json:"error,omitempty"`
	Id              *string              `json:"id,omitempty"`
	Name            *string              `json:"name,omitempty"`
	PercentComplete *int64               `json:"percentComplete,omitempty"`
	StartTime       *string              `json:"startTime,omitempty"`
	Status          *string              `json:"status,omitempty"`
}
