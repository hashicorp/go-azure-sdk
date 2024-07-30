package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordSingleSignOnSettings struct {
	Fields    *[]PasswordSingleSignOnField `json:"fields,omitempty"`
	ODataType *string                      `json:"@odata.type,omitempty"`
}
