package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingCurrency struct {
	Id        *string `json:"id,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Symbol    *string `json:"symbol,omitempty"`
}
