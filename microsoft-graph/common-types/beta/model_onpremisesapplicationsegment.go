package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesApplicationSegment struct {
	AlternateUrl       *string              `json:"alternateUrl,omitempty"`
	CorsConfigurations *[]CorsConfiguration `json:"corsConfigurations,omitempty"`
	ExternalUrl        *string              `json:"externalUrl,omitempty"`
	InternalUrl        *string              `json:"internalUrl,omitempty"`
	ODataType          *string              `json:"@odata.type,omitempty"`
}
