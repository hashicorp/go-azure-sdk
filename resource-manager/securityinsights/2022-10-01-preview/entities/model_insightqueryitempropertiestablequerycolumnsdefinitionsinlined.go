package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightQueryItemPropertiesTableQueryColumnsDefinitionsInlined struct {
	Header          *string     `json:"header,omitempty"`
	OutputType      *OutputType `json:"outputType,omitempty"`
	SupportDeepLink *bool       `json:"supportDeepLink,omitempty"`
}
