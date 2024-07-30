package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchResponse struct {
	HitsContainers          *[]SearchHitsContainer    `json:"hitsContainers,omitempty"`
	ODataType               *string                   `json:"@odata.type,omitempty"`
	QueryAlterationResponse *AlterationResponse       `json:"queryAlterationResponse,omitempty"`
	ResultTemplates         *ResultTemplateDictionary `json:"resultTemplates,omitempty"`
	SearchTerms             *[]string                 `json:"searchTerms,omitempty"`
}
