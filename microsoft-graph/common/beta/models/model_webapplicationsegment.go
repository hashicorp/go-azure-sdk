package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebApplicationSegment struct {
	AlternateUrl       *string                `json:"alternateUrl,omitempty"`
	CorsConfigurations *[]CorsConfigurationv2 `json:"corsConfigurations,omitempty"`
	ExternalUrl        *string                `json:"externalUrl,omitempty"`
	Id                 *string                `json:"id,omitempty"`
	InternalUrl        *string                `json:"internalUrl,omitempty"`
	ODataType          *string                `json:"@odata.type,omitempty"`
}
