package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CvssSummary struct {
	ODataType    *string              `json:"@odata.type,omitempty"`
	Severity     *CvssSummarySeverity `json:"severity,omitempty"`
	VectorString *string              `json:"vectorString,omitempty"`
}
