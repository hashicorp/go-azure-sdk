package triggerruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunFilterParameters struct {
	ContinuationToken *string            `json:"continuationToken,omitempty"`
	Filters           *[]RunQueryFilter  `json:"filters,omitempty"`
	LastUpdatedAfter  string             `json:"lastUpdatedAfter"`
	LastUpdatedBefore string             `json:"lastUpdatedBefore"`
	OrderBy           *[]RunQueryOrderBy `json:"orderBy,omitempty"`
}
