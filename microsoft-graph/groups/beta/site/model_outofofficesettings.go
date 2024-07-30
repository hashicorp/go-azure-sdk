package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutOfOfficeSettings struct {
	IsOutOfOffice *bool   `json:"isOutOfOffice,omitempty"`
	Message       *string `json:"message,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
}
