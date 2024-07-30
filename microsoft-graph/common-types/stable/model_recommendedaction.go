package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedAction struct {
	ActionWebUrl         *string  `json:"actionWebUrl,omitempty"`
	ODataType            *string  `json:"@odata.type,omitempty"`
	PotentialScoreImpact *float64 `json:"potentialScoreImpact,omitempty"`
	Title                *string  `json:"title,omitempty"`
}
