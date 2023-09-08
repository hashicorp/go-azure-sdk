package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionCalloutResult struct {
	CalloutDateTime   *string `json:"calloutDateTime,omitempty"`
	CustomExtensionId *string `json:"customExtensionId,omitempty"`
	ErrorCode         *int64  `json:"errorCode,omitempty"`
	HttpStatus        *int64  `json:"httpStatus,omitempty"`
	NumberOfAttempts  *int64  `json:"numberOfAttempts,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
}
