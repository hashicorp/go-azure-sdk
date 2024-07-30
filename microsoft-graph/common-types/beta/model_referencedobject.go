package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReferencedObject struct {
	ODataType            *string `json:"@odata.type,omitempty"`
	ReferencedObjectName *string `json:"referencedObjectName,omitempty"`
	ReferencedProperty   *string `json:"referencedProperty,omitempty"`
}
