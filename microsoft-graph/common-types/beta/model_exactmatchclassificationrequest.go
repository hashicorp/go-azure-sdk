package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactMatchClassificationRequest struct {
	ContentClassifications *[]ContentClassification `json:"contentClassifications,omitempty"`
	ODataType              *string                  `json:"@odata.type,omitempty"`
	SensitiveTypeIds       *[]string                `json:"sensitiveTypeIds,omitempty"`
	Text                   *string                  `json:"text,omitempty"`
	TimeoutInMs            *int64                   `json:"timeoutInMs,omitempty"`
}
