package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAlteration struct {
	AlteredHighlightedQueryString *string              `json:"alteredHighlightedQueryString,omitempty"`
	AlteredQueryString            *string              `json:"alteredQueryString,omitempty"`
	AlteredQueryTokens            *[]AlteredQueryToken `json:"alteredQueryTokens,omitempty"`
	ODataType                     *string              `json:"@odata.type,omitempty"`
}
