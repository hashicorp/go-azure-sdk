package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineStateSummary struct {
	ConflictCount      *int64  `json:"conflictCount,omitempty"`
	ErrorCount         *int64  `json:"errorCount,omitempty"`
	Id                 *string `json:"id,omitempty"`
	NotApplicableCount *int64  `json:"notApplicableCount,omitempty"`
	NotSecureCount     *int64  `json:"notSecureCount,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
	SecureCount        *int64  `json:"secureCount,omitempty"`
	UnknownCount       *int64  `json:"unknownCount,omitempty"`
}
