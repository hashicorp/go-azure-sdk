package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateProfileInventorySyncStatus struct {
	DriverInventorySyncState   *WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState `json:"driverInventorySyncState,omitempty"`
	LastSuccessfulSyncDateTime *string                                                                `json:"lastSuccessfulSyncDateTime,omitempty"`
	ODataType                  *string                                                                `json:"@odata.type,omitempty"`
}
