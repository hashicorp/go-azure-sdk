package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesSafeguardProfile struct {
	Category  *WindowsUpdatesSafeguardProfileCategory `json:"category,omitempty"`
	ODataType *string                                 `json:"@odata.type,omitempty"`
}
