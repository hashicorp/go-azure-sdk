package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchHit struct {
	ContentSource    *string `json:"contentSource,omitempty"`
	HitId            *string `json:"hitId,omitempty"`
	IsCollapsed      *bool   `json:"isCollapsed,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	Rank             *int64  `json:"rank,omitempty"`
	Resource         *Entity `json:"resource,omitempty"`
	ResultTemplateId *string `json:"resultTemplateId,omitempty"`
	Summary          *string `json:"summary,omitempty"`
}
