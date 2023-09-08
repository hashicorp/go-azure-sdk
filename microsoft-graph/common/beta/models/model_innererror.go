package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InnerError struct {
	ClientRequestId *string `json:"client-request-id,omitempty"`
	Date            *string `json:"Date,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	RequestId       *string `json:"request-id,omitempty"`
}
