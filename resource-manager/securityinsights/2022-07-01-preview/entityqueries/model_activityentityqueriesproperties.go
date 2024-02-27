package entityqueries

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityEntityQueriesProperties struct {
	Content                 *string                                          `json:"content,omitempty"`
	CreatedTimeUtc          *string                                          `json:"createdTimeUtc,omitempty"`
	Description             *string                                          `json:"description,omitempty"`
	Enabled                 *bool                                            `json:"enabled,omitempty"`
	EntitiesFilter          *map[string][]string                             `json:"entitiesFilter,omitempty"`
	InputEntityType         *EntityType                                      `json:"inputEntityType,omitempty"`
	LastModifiedTimeUtc     *string                                          `json:"lastModifiedTimeUtc,omitempty"`
	QueryDefinitions        *ActivityEntityQueriesPropertiesQueryDefinitions `json:"queryDefinitions,omitempty"`
	RequiredInputFieldsSets *[][]string                                      `json:"requiredInputFieldsSets,omitempty"`
	TemplateName            *string                                          `json:"templateName,omitempty"`
	Title                   *string                                          `json:"title,omitempty"`
}

func (o *ActivityEntityQueriesProperties) GetCreatedTimeUtcAsTime() (*time.Time, error) {
	if o.CreatedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *ActivityEntityQueriesProperties) GetLastModifiedTimeUtcAsTime() (*time.Time, error) {
	if o.LastModifiedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedTimeUtc, "2006-01-02T15:04:05Z07:00")
}
