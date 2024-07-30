package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesRateDrivenRolloutSettings struct {
	DevicesPerOffer       *int64  `json:"devicesPerOffer,omitempty"`
	DurationBetweenOffers *string `json:"durationBetweenOffers,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
}
