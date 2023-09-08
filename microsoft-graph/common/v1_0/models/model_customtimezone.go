package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomTimeZone struct {
	Bias           *int64                  `json:"bias,omitempty"`
	DaylightOffset *DaylightTimeZoneOffset `json:"daylightOffset,omitempty"`
	Name           *string                 `json:"name,omitempty"`
	ODataType      *string                 `json:"@odata.type,omitempty"`
	StandardOffset *StandardTimeZoneOffset `json:"standardOffset,omitempty"`
}
