package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CorsConfiguration struct {
	AllowedHeaders  *[]string `json:"allowedHeaders,omitempty"`
	AllowedMethods  *[]string `json:"allowedMethods,omitempty"`
	AllowedOrigins  *[]string `json:"allowedOrigins,omitempty"`
	MaxAgeInSeconds *int64    `json:"maxAgeInSeconds,omitempty"`
	ODataType       *string   `json:"@odata.type,omitempty"`
	Resource        *string   `json:"resource,omitempty"`
}
