package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditProperty struct {
	// Display name.
	DisplayName *string `json:"displayName,omitempty"`

	// New value.
	NewValue *string `json:"newValue,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Old value.
	OldValue nullable.Type[string] `json:"oldValue,omitempty"`
}
