package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatRestrictions struct {
	AllowTextOnly *bool   `json:"allowTextOnly,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
}
