package restorables

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableLocationResource struct {
	CreationTime                      *string `json:"creationTime,omitempty"`
	DeletionTime                      *string `json:"deletionTime,omitempty"`
	LocationName                      *string `json:"locationName,omitempty"`
	RegionalDatabaseAccountInstanceId *string `json:"regionalDatabaseAccountInstanceId,omitempty"`
}
