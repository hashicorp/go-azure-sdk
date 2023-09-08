package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitiveContentLocation struct {
	Confidence *int64                      `json:"confidence,omitempty"`
	Evidences  *[]SensitiveContentEvidence `json:"evidences,omitempty"`
	IdMatch    *string                     `json:"idMatch,omitempty"`
	Length     *int64                      `json:"length,omitempty"`
	ODataType  *string                     `json:"@odata.type,omitempty"`
	Offset     *int64                      `json:"offset,omitempty"`
}
