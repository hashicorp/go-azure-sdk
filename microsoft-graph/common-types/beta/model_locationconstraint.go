package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationConstraint struct {
	IsRequired      *bool                     `json:"isRequired,omitempty"`
	Locations       *[]LocationConstraintItem `json:"locations,omitempty"`
	ODataType       *string                   `json:"@odata.type,omitempty"`
	SuggestLocation *bool                     `json:"suggestLocation,omitempty"`
}
