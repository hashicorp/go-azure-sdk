package connectordefinitions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomizableConnectorDefinitionProperties struct {
	ConnectionsConfig *CustomizableConnectionsConfig `json:"connectionsConfig,omitempty"`
	ConnectorUiConfig CustomizableConnectorUiConfig  `json:"connectorUiConfig"`
	CreatedTimeUtc    *string                        `json:"createdTimeUtc,omitempty"`
	LastModifiedUtc   *string                        `json:"lastModifiedUtc,omitempty"`
}

func (o *CustomizableConnectorDefinitionProperties) GetCreatedTimeUtcAsTime() (*time.Time, error) {
	if o.CreatedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *CustomizableConnectorDefinitionProperties) SetCreatedTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedTimeUtc = &formatted
}

func (o *CustomizableConnectorDefinitionProperties) GetLastModifiedUtcAsTime() (*time.Time, error) {
	if o.LastModifiedUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *CustomizableConnectorDefinitionProperties) SetLastModifiedUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModifiedUtc = &formatted
}
