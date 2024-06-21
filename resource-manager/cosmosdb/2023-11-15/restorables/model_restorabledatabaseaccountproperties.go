package restorables

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableDatabaseAccountProperties struct {
	AccountName          *string                       `json:"accountName,omitempty"`
	ApiType              *ApiType                      `json:"apiType,omitempty"`
	CreationTime         *string                       `json:"creationTime,omitempty"`
	DeletionTime         *string                       `json:"deletionTime,omitempty"`
	OldestRestorableTime *string                       `json:"oldestRestorableTime,omitempty"`
	RestorableLocations  *[]RestorableLocationResource `json:"restorableLocations,omitempty"`
}
