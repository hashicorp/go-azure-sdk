package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestorSettings struct {
	AcceptRequests    *bool      `json:"acceptRequests,omitempty"`
	AllowedRequestors *[]UserSet `json:"allowedRequestors,omitempty"`
	ODataType         *string    `json:"@odata.type,omitempty"`
	ScopeType         *string    `json:"scopeType,omitempty"`
}
