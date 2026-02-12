package iotcentrals

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Destination = WebhookV1Destination{}

type WebhookV1Destination struct {
	Authorization        WebhookV1DestinationAuth                      `json:"authorization"`
	HeaderCustomizations *map[string]WebhookV1DestinationCustomization `json:"headerCustomizations,omitempty"`
	QueryCustomizations  *map[string]WebhookV1DestinationCustomization `json:"queryCustomizations,omitempty"`
	Url                  string                                        `json:"url"`

	// Fields inherited from Destination

	DisplayName    string             `json:"displayName"`
	Errors         *[]DataExportError `json:"errors,omitempty"`
	Id             *string            `json:"id,omitempty"`
	LastExportTime *string            `json:"lastExportTime,omitempty"`
	Status         *string            `json:"status,omitempty"`
	Type           string             `json:"type"`
}

func (s WebhookV1Destination) Destination() BaseDestinationImpl {
	return BaseDestinationImpl{
		DisplayName:    s.DisplayName,
		Errors:         s.Errors,
		Id:             s.Id,
		LastExportTime: s.LastExportTime,
		Status:         s.Status,
		Type:           s.Type,
	}
}

func (o *WebhookV1Destination) GetLastExportTimeAsTime() (*time.Time, error) {
	if o.LastExportTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastExportTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WebhookV1Destination) SetLastExportTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastExportTime = &formatted
}

var _ json.Marshaler = WebhookV1Destination{}

func (s WebhookV1Destination) MarshalJSON() ([]byte, error) {
	type wrapper WebhookV1Destination
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WebhookV1Destination: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WebhookV1Destination: %+v", err)
	}

	decoded["type"] = "webhook@v1"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WebhookV1Destination: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WebhookV1Destination{}

func (s *WebhookV1Destination) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		HeaderCustomizations *map[string]WebhookV1DestinationCustomization `json:"headerCustomizations,omitempty"`
		QueryCustomizations  *map[string]WebhookV1DestinationCustomization `json:"queryCustomizations,omitempty"`
		Url                  string                                        `json:"url"`
		DisplayName          string                                        `json:"displayName"`
		Errors               *[]DataExportError                            `json:"errors,omitempty"`
		Id                   *string                                       `json:"id,omitempty"`
		LastExportTime       *string                                       `json:"lastExportTime,omitempty"`
		Status               *string                                       `json:"status,omitempty"`
		Type                 string                                        `json:"type"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.HeaderCustomizations = decoded.HeaderCustomizations
	s.QueryCustomizations = decoded.QueryCustomizations
	s.Url = decoded.Url
	s.DisplayName = decoded.DisplayName
	s.Errors = decoded.Errors
	s.Id = decoded.Id
	s.LastExportTime = decoded.LastExportTime
	s.Status = decoded.Status
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WebhookV1Destination into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authorization"]; ok {
		impl, err := UnmarshalWebhookV1DestinationAuthImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Authorization' for 'WebhookV1Destination': %+v", err)
		}
		s.Authorization = impl
	}

	return nil
}
