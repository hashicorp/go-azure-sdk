package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactMatchDetectedSensitiveContent struct {
	Confidence            *int64                      `json:"confidence,omitempty"`
	DisplayName           *string                     `json:"displayName,omitempty"`
	Id                    *string                     `json:"id,omitempty"`
	Matches               *[]SensitiveContentLocation `json:"matches,omitempty"`
	ODataType             *string                     `json:"@odata.type,omitempty"`
	RecommendedConfidence *int64                      `json:"recommendedConfidence,omitempty"`
	UniqueCount           *int64                      `json:"uniqueCount,omitempty"`
}
