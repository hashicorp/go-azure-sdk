package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MimeContent struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the content mime type.
	Type nullable.Type[string] `json:"type,omitempty"`

	// The byte array that contains the actual content.
	Value nullable.Type[string] `json:"value,omitempty"`
}
