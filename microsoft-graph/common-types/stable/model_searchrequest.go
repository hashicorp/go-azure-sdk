package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchRequest struct {
	AggregationFilters        *[]string                  `json:"aggregationFilters,omitempty"`
	Aggregations              *[]AggregationOption       `json:"aggregations,omitempty"`
	CollapseProperties        *[]CollapseProperty        `json:"collapseProperties,omitempty"`
	ContentSources            *[]string                  `json:"contentSources,omitempty"`
	EnableTopResults          *bool                      `json:"enableTopResults,omitempty"`
	EntityTypes               *SearchRequestEntityTypes  `json:"entityTypes,omitempty"`
	Fields                    *[]string                  `json:"fields,omitempty"`
	From                      *int64                     `json:"from,omitempty"`
	ODataType                 *string                    `json:"@odata.type,omitempty"`
	Query                     *SearchQuery               `json:"query,omitempty"`
	QueryAlterationOptions    *SearchAlterationOptions   `json:"queryAlterationOptions,omitempty"`
	Region                    *string                    `json:"region,omitempty"`
	ResultTemplateOptions     *ResultTemplateOption      `json:"resultTemplateOptions,omitempty"`
	SharePointOneDriveOptions *SharePointOneDriveOptions `json:"sharePointOneDriveOptions,omitempty"`
	Size                      *int64                     `json:"size,omitempty"`
	SortProperties            *[]SortProperty            `json:"sortProperties,omitempty"`
}
