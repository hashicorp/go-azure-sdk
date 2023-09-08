package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeGraphInsights struct {
	Id        *string          `json:"id,omitempty"`
	ODataType *string          `json:"@odata.type,omitempty"`
	Shared    *[]SharedInsight `json:"shared,omitempty"`
	Trending  *[]Trending      `json:"trending,omitempty"`
	Used      *[]UsedInsight   `json:"used,omitempty"`
}
