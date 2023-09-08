package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListInfo struct {
	ContentTypesEnabled *bool   `json:"contentTypesEnabled,omitempty"`
	Hidden              *bool   `json:"hidden,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	Template            *string `json:"template,omitempty"`
}
