package iotcentrals

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Export struct {
	Destinations   []DestinationReference `json:"destinations"`
	DisplayName    string                 `json:"displayName"`
	Enabled        bool                   `json:"enabled"`
	Enrichments    *map[string]Enrichment `json:"enrichments,omitempty"`
	Errors         *[]DataExportError     `json:"errors,omitempty"`
	Filter         *string                `json:"filter,omitempty"`
	Id             *string                `json:"id,omitempty"`
	LastExportTime *string                `json:"lastExportTime,omitempty"`
	Source         DestinationSource      `json:"source"`
	Status         *string                `json:"status,omitempty"`
}

func (o *Export) GetLastExportTimeAsTime() (*time.Time, error) {
	if o.LastExportTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastExportTime, "2006-01-02T15:04:05Z07:00")
}

func (o *Export) SetLastExportTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastExportTime = &formatted
}
