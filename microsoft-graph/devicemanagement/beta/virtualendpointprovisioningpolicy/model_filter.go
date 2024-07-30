package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Filter struct {
	CategoryFilterGroups *[]FilterGroup `json:"categoryFilterGroups,omitempty"`
	Groups               *[]FilterGroup `json:"groups,omitempty"`
	InputFilterGroups    *[]FilterGroup `json:"inputFilterGroups,omitempty"`
	ODataType            *string        `json:"@odata.type,omitempty"`
}
