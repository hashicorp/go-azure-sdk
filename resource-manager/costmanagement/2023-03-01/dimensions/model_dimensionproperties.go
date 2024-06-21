package dimensions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DimensionProperties struct {
	Category        *string   `json:"category,omitempty"`
	Data            *[]string `json:"data,omitempty"`
	Description     *string   `json:"description,omitempty"`
	FilterEnabled   *bool     `json:"filterEnabled,omitempty"`
	GroupingEnabled *bool     `json:"groupingEnabled,omitempty"`
	NextLink        *string   `json:"nextLink,omitempty"`
	Total           *int64    `json:"total,omitempty"`
	UsageEnd        *string   `json:"usageEnd,omitempty"`
	UsageStart      *string   `json:"usageStart,omitempty"`
}
