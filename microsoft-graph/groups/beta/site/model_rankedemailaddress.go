package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RankedEmailAddress struct {
	Address   *string  `json:"address,omitempty"`
	ODataType *string  `json:"@odata.type,omitempty"`
	Rank      *float64 `json:"rank,omitempty"`
}
