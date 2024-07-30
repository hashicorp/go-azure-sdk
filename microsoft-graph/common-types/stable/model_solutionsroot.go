package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SolutionsRoot struct {
	BookingBusinesses *[]BookingBusiness `json:"bookingBusinesses,omitempty"`
	BookingCurrencies *[]BookingCurrency `json:"bookingCurrencies,omitempty"`
	ODataType         *string            `json:"@odata.type,omitempty"`
	VirtualEvents     *VirtualEventsRoot `json:"virtualEvents,omitempty"`
}
