package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingCustomer struct {
	Addresses    *[]PhysicalAddress `json:"addresses,omitempty"`
	DisplayName  *string            `json:"displayName,omitempty"`
	EmailAddress *string            `json:"emailAddress,omitempty"`
	Id           *string            `json:"id,omitempty"`
	ODataType    *string            `json:"@odata.type,omitempty"`
	Phones       *[]Phone           `json:"phones,omitempty"`
}
