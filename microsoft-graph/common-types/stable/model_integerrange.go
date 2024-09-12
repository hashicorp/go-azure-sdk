package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegerRange struct {
	// The inclusive upper bound of the integer range.
	End nullable.Type[int64] `json:"end,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The inclusive lower bound of the integer range.
	Start nullable.Type[int64] `json:"start,omitempty"`
}
