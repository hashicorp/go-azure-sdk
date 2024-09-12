package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSLaunchItem struct {
	// Whether or not to hide the item from the Users and Groups List.
	Hide *bool `json:"hide,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Path to the launch item.
	Path *string `json:"path,omitempty"`
}
