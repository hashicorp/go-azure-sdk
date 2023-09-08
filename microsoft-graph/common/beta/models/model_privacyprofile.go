package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivacyProfile struct {
	ContactEmail *string `json:"contactEmail,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	StatementUrl *string `json:"statementUrl,omitempty"`
}
