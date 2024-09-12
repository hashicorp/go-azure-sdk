package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudClipboardItemPayload struct {
	// The formatName version of the value of a cloud clipboard encoded in base64.
	Content *string `json:"content,omitempty"`

	// For a list of possible values see formatName values.
	FormatName *string `json:"formatName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
