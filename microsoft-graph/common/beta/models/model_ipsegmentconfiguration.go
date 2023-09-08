package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IpSegmentConfiguration struct {
	ApplicationSegments *[]IpApplicationSegment `json:"applicationSegments,omitempty"`
	ODataType           *string                 `json:"@odata.type,omitempty"`
}
