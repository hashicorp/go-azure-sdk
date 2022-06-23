package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInsightsResultsMetadata struct {
	Errors     *[]GetInsightsError `json:"errors,omitempty"`
	TotalCount int64               `json:"totalCount"`
}
