package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityGetInsightsResponse struct {
	MetaData *GetInsightsResultsMetadata `json:"metaData"`
	Value    *[]EntityInsightItem        `json:"value,omitempty"`
}
