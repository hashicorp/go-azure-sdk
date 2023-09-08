package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDefenderScanActionResult struct {
	ActionName          *string                                     `json:"actionName,omitempty"`
	ActionState         *WindowsDefenderScanActionResultActionState `json:"actionState,omitempty"`
	LastUpdatedDateTime *string                                     `json:"lastUpdatedDateTime,omitempty"`
	ODataType           *string                                     `json:"@odata.type,omitempty"`
	ScanType            *string                                     `json:"scanType,omitempty"`
	StartDateTime       *string                                     `json:"startDateTime,omitempty"`
}
