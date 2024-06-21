package openaiintegration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenAIIntegrationProperties struct {
	Key                    *string `json:"key,omitempty"`
	LastRefreshAt          *string `json:"lastRefreshAt,omitempty"`
	OpenAIResourceEndpoint *string `json:"openAIResourceEndpoint,omitempty"`
	OpenAIResourceId       *string `json:"openAIResourceId,omitempty"`
}
