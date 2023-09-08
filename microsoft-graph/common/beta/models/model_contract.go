package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Contract struct {
	ContractType      *string `json:"contractType,omitempty"`
	CustomerId        *string `json:"customerId,omitempty"`
	DefaultDomainName *string `json:"defaultDomainName,omitempty"`
	DeletedDateTime   *string `json:"deletedDateTime,omitempty"`
	DisplayName       *string `json:"displayName,omitempty"`
	Id                *string `json:"id,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
}
