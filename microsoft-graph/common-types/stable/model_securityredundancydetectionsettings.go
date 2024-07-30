package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRedundancyDetectionSettings struct {
	IsEnabled           *bool   `json:"isEnabled,omitempty"`
	MaxWords            *int64  `json:"maxWords,omitempty"`
	MinWords            *int64  `json:"minWords,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	SimilarityThreshold *int64  `json:"similarityThreshold,omitempty"`
}
