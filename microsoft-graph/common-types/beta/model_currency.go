package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Currency struct {
	AmountDecimalPlaces     *string  `json:"amountDecimalPlaces,omitempty"`
	AmountRoundingPrecision *float64 `json:"amountRoundingPrecision,omitempty"`
	Code                    *string  `json:"code,omitempty"`
	DisplayName             *string  `json:"displayName,omitempty"`
	Id                      *string  `json:"id,omitempty"`
	LastModifiedDateTime    *string  `json:"lastModifiedDateTime,omitempty"`
	ODataType               *string  `json:"@odata.type,omitempty"`
	Symbol                  *string  `json:"symbol,omitempty"`
}
