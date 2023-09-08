package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchQuery struct {
	ODataType     *string `json:"@odata.type,omitempty"`
	QueryString   *string `json:"queryString,omitempty"`
	QueryTemplate *string `json:"queryTemplate,omitempty"`
}
