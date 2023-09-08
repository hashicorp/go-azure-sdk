package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitiveContentEvidence struct {
	Length    *int64  `json:"length,omitempty"`
	Match     *string `json:"match,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Offset    *int64  `json:"offset,omitempty"`
}
