package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Property struct {
	Aliases       *[]string       `json:"aliases,omitempty"`
	IsQueryable   *bool           `json:"isQueryable,omitempty"`
	IsRefinable   *bool           `json:"isRefinable,omitempty"`
	IsRetrievable *bool           `json:"isRetrievable,omitempty"`
	IsSearchable  *bool           `json:"isSearchable,omitempty"`
	Labels        *PropertyLabels `json:"labels,omitempty"`
	Name          *string         `json:"name,omitempty"`
	ODataType     *string         `json:"@odata.type,omitempty"`
	Type          *PropertyType   `json:"type,omitempty"`
}
