package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MfaDetail struct {
	AuthDetail *string `json:"authDetail,omitempty"`
	AuthMethod *string `json:"authMethod,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
}
