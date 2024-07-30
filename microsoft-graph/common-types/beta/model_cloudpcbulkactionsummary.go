package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCBulkActionSummary struct {
	FailedCount       *int64  `json:"failedCount,omitempty"`
	InProgressCount   *int64  `json:"inProgressCount,omitempty"`
	NotSupportedCount *int64  `json:"notSupportedCount,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	PendingCount      *int64  `json:"pendingCount,omitempty"`
	SuccessfulCount   *int64  `json:"successfulCount,omitempty"`
}
