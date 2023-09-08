package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResultTemplateOption struct {
	EnableResultTemplate *bool   `json:"enableResultTemplate,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
}
