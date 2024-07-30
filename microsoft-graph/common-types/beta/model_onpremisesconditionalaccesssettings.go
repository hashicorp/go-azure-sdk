package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesConditionalAccessSettings struct {
	Enabled             *bool     `json:"enabled,omitempty"`
	ExcludedGroups      *[]string `json:"excludedGroups,omitempty"`
	Id                  *string   `json:"id,omitempty"`
	IncludedGroups      *[]string `json:"includedGroups,omitempty"`
	ODataType           *string   `json:"@odata.type,omitempty"`
	OverrideDefaultRule *bool     `json:"overrideDefaultRule,omitempty"`
}
