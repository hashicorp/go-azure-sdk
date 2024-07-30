package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegionalFormatOverrides struct {
	Calendar        *string `json:"calendar,omitempty"`
	FirstDayOfWeek  *string `json:"firstDayOfWeek,omitempty"`
	LongDateFormat  *string `json:"longDateFormat,omitempty"`
	LongTimeFormat  *string `json:"longTimeFormat,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	ShortDateFormat *string `json:"shortDateFormat,omitempty"`
	ShortTimeFormat *string `json:"shortTimeFormat,omitempty"`
	TimeZone        *string `json:"timeZone,omitempty"`
}
