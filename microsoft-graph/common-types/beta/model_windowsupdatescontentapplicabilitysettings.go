package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesContentApplicabilitySettings struct {
	ODataType               *string                          `json:"@odata.type,omitempty"`
	OfferWhileRecommendedBy *[]string                        `json:"offerWhileRecommendedBy,omitempty"`
	Safeguard               *WindowsUpdatesSafeguardSettings `json:"safeguard,omitempty"`
}
