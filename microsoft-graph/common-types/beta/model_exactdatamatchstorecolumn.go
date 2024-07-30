package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactDataMatchStoreColumn struct {
	IgnoredDelimiters *[]string `json:"ignoredDelimiters,omitempty"`
	IsCaseInsensitive *bool     `json:"isCaseInsensitive,omitempty"`
	IsSearchable      *bool     `json:"isSearchable,omitempty"`
	Name              *string   `json:"name,omitempty"`
	ODataType         *string   `json:"@odata.type,omitempty"`
}
