package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntentUserStateSummary struct {
	ConflictCount      *int64  `json:"conflictCount,omitempty"`
	ErrorCount         *int64  `json:"errorCount,omitempty"`
	FailedCount        *int64  `json:"failedCount,omitempty"`
	Id                 *string `json:"id,omitempty"`
	NotApplicableCount *int64  `json:"notApplicableCount,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
	SuccessCount       *int64  `json:"successCount,omitempty"`
}
