package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessLocations struct {
	ExcludeLocations *[]string `json:"excludeLocations,omitempty"`
	IncludeLocations *[]string `json:"includeLocations,omitempty"`
	ODataType        *string   `json:"@odata.type,omitempty"`
}
