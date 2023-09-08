package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentCustomization struct {
	AttributeCollection            *[]KeyValue `json:"attributeCollection,omitempty"`
	AttributeCollectionRelativeUrl *string     `json:"attributeCollectionRelativeUrl,omitempty"`
	ODataType                      *string     `json:"@odata.type,omitempty"`
}
