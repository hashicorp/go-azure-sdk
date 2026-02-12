package iotcentrals

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Destination = ExportDestination{}

type ExportDestination struct {
	Transform *string `json:"transform,omitempty"`

	// Fields inherited from Destination

	DisplayName    string             `json:"displayName"`
	Errors         *[]DataExportError `json:"errors,omitempty"`
	Id             *string            `json:"id,omitempty"`
	LastExportTime *string            `json:"lastExportTime,omitempty"`
	Status         *string            `json:"status,omitempty"`
	Type           string             `json:"type"`
}

func (s ExportDestination) Destination() BaseDestinationImpl {
	return BaseDestinationImpl{
		DisplayName:    s.DisplayName,
		Errors:         s.Errors,
		Id:             s.Id,
		LastExportTime: s.LastExportTime,
		Status:         s.Status,
		Type:           s.Type,
	}
}

func (o *ExportDestination) GetLastExportTimeAsTime() (*time.Time, error) {
	if o.LastExportTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastExportTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ExportDestination) SetLastExportTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastExportTime = &formatted
}

var _ json.Marshaler = ExportDestination{}

func (s ExportDestination) MarshalJSON() ([]byte, error) {
	type wrapper ExportDestination
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExportDestination: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExportDestination: %+v", err)
	}

	decoded["type"] = "ExportDestination"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExportDestination: %+v", err)
	}

	return encoded, nil
}
