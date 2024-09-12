package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Phone struct {
	// The phone number.
	Number nullable.Type[string] `json:"number,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type of phone number. Possible values are: home, business, mobile, other, assistant, homeFax, businessFax,
	// otherFax, pager, radio.
	Type *PhoneType `json:"type,omitempty"`
}
