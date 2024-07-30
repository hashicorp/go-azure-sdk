package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCvssSummary struct {
	ODataType    *string                      `json:"@odata.type,omitempty"`
	Score        *float64                     `json:"score,omitempty"`
	Severity     *SecurityCvssSummarySeverity `json:"severity,omitempty"`
	VectorString *string                      `json:"vectorString,omitempty"`
}
