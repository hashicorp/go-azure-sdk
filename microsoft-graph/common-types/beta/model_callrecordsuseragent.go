package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsUserAgent struct {
	ApplicationVersion *string `json:"applicationVersion,omitempty"`
	HeaderValue        *string `json:"headerValue,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
}
