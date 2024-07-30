package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesGradualRolloutSettings struct {
	DurationBetweenOffers *string `json:"durationBetweenOffers,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
}
