package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateRolloutSettings struct {
	ODataType               *string `json:"@odata.type,omitempty"`
	OfferEndDateTimeInUTC   *string `json:"offerEndDateTimeInUTC,omitempty"`
	OfferIntervalInDays     *int64  `json:"offerIntervalInDays,omitempty"`
	OfferStartDateTimeInUTC *string `json:"offerStartDateTimeInUTC,omitempty"`
}
