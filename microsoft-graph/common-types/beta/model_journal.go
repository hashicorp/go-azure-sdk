package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Journal struct {
	Account                *Account              `json:"account,omitempty"`
	BalancingAccountId     nullable.Type[string] `json:"balancingAccountId,omitempty"`
	BalancingAccountNumber nullable.Type[string] `json:"balancingAccountNumber,omitempty"`
	Code                   nullable.Type[string] `json:"code,omitempty"`
	DisplayName            nullable.Type[string] `json:"displayName,omitempty"`
	Id                     *string               `json:"id,omitempty"`
	JournalLines           *[]JournalLine        `json:"journalLines,omitempty"`
	LastModifiedDateTime   nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
