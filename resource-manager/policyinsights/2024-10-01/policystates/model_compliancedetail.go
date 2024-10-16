package policystates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceDetail struct {
	ComplianceState *string `json:"complianceState,omitempty"`
	Count           *int64  `json:"count,omitempty"`
}
