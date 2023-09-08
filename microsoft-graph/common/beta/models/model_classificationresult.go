package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClassificationResult struct {
	ConfidenceLevel *int64  `json:"confidenceLevel,omitempty"`
	Count           *int64  `json:"count,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	SensitiveTypeId *string `json:"sensitiveTypeId,omitempty"`
}
