package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentClassification struct {
	Confidence      *int64           `json:"confidence,omitempty"`
	Matches         *[]MatchLocation `json:"matches,omitempty"`
	ODataType       *string          `json:"@odata.type,omitempty"`
	SensitiveTypeId *string          `json:"sensitiveTypeId,omitempty"`
	UniqueCount     *int64           `json:"uniqueCount,omitempty"`
}
