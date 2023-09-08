package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchHitsContainer struct {
	Aggregations         *[]SearchAggregation `json:"aggregations,omitempty"`
	Hits                 *[]SearchHit         `json:"hits,omitempty"`
	MoreResultsAvailable *bool                `json:"moreResultsAvailable,omitempty"`
	ODataType            *string              `json:"@odata.type,omitempty"`
	Total                *int64               `json:"total,omitempty"`
}
