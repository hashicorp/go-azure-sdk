package openaiintegration

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenAIIntegrationProperties struct {
	Key                    *string `json:"key,omitempty"`
	LastRefreshAt          *string `json:"lastRefreshAt,omitempty"`
	OpenAIResourceEndpoint *string `json:"openAIResourceEndpoint,omitempty"`
	OpenAIResourceId       *string `json:"openAIResourceId,omitempty"`
}

func (o *OpenAIIntegrationProperties) GetLastRefreshAtAsTime() (*time.Time, error) {
	if o.LastRefreshAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastRefreshAt, "2006-01-02T15:04:05Z07:00")
}

func (o *OpenAIIntegrationProperties) SetLastRefreshAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastRefreshAt = &formatted
}
