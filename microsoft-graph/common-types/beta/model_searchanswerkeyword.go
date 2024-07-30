package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAnswerKeyword struct {
	Keywords             *[]string `json:"keywords,omitempty"`
	MatchSimilarKeywords *bool     `json:"matchSimilarKeywords,omitempty"`
	ODataType            *string   `json:"@odata.type,omitempty"`
	ReservedKeywords     *[]string `json:"reservedKeywords,omitempty"`
}
