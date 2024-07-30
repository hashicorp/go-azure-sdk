package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FreeBusyError struct {
	Message      *string `json:"message,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
}
