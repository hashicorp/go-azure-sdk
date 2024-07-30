package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperatingSystemVersionRange struct {
	Description    *string `json:"description,omitempty"`
	HighestVersion *string `json:"highestVersion,omitempty"`
	LowestVersion  *string `json:"lowestVersion,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
}
