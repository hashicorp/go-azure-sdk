package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementTroubleshootingEvent = MobileAppTroubleshootingEvent{}

type MobileAppTroubleshootingEvent struct {
	// Indicates collection of App Log Upload Request.
	AppLogCollectionRequests *[]AppLogCollectionRequest `json:"appLogCollectionRequests,omitempty"`

	// Intune application identifier.
	ApplicationId nullable.Type[string] `json:"applicationId,omitempty"`

	// Device identifier created or collected by Intune.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// Intune Mobile Application Troubleshooting History Item
	History *[]MobileAppTroubleshootingHistoryItem `json:"history,omitempty"`

	// Device identifier created or collected by Intune.
	ManagedDeviceIdentifier nullable.Type[string] `json:"managedDeviceIdentifier,omitempty"`

	// Identifier for the user that tried to enroll the device.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// Fields inherited from DeviceManagementTroubleshootingEvent

	// A set of string key and string value pairs which provides additional information on the Troubleshooting event
	AdditionalInformation *[]KeyValuePair `json:"additionalInformation,omitempty"`

	// Id used for tracing the failure in the service.
	CorrelationId nullable.Type[string] `json:"correlationId,omitempty"`

	// Time when the event occurred .
	EventDateTime *string `json:"eventDateTime,omitempty"`

	// Event Name corresponding to the Troubleshooting Event. It is an Optional field
	EventName nullable.Type[string] `json:"eventName,omitempty"`

	// Object containing detailed information about the error and its remediation.
	TroubleshootingErrorDetails *DeviceManagementTroubleshootingErrorDetails `json:"troubleshootingErrorDetails,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s MobileAppTroubleshootingEvent) DeviceManagementTroubleshootingEvent() BaseDeviceManagementTroubleshootingEventImpl {
	return BaseDeviceManagementTroubleshootingEventImpl{
		AdditionalInformation:       s.AdditionalInformation,
		CorrelationId:               s.CorrelationId,
		EventDateTime:               s.EventDateTime,
		EventName:                   s.EventName,
		TroubleshootingErrorDetails: s.TroubleshootingErrorDetails,
		Id:                          s.Id,
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
	}
}

func (s MobileAppTroubleshootingEvent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MobileAppTroubleshootingEvent{}

func (s MobileAppTroubleshootingEvent) MarshalJSON() ([]byte, error) {
	type wrapper MobileAppTroubleshootingEvent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobileAppTroubleshootingEvent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppTroubleshootingEvent: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.mobileAppTroubleshootingEvent"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobileAppTroubleshootingEvent: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MobileAppTroubleshootingEvent{}

func (s *MobileAppTroubleshootingEvent) UnmarshalJSON(bytes []byte) error {
	type alias MobileAppTroubleshootingEvent
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into MobileAppTroubleshootingEvent: %+v", err)
	}

	s.AdditionalInformation = decoded.AdditionalInformation
	s.AppLogCollectionRequests = decoded.AppLogCollectionRequests
	s.ApplicationId = decoded.ApplicationId
	s.CorrelationId = decoded.CorrelationId
	s.DeviceId = decoded.DeviceId
	s.EventDateTime = decoded.EventDateTime
	s.EventName = decoded.EventName
	s.Id = decoded.Id
	s.ManagedDeviceIdentifier = decoded.ManagedDeviceIdentifier
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.TroubleshootingErrorDetails = decoded.TroubleshootingErrorDetails
	s.UserId = decoded.UserId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MobileAppTroubleshootingEvent into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["history"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling History into list []json.RawMessage: %+v", err)
		}

		output := make([]MobileAppTroubleshootingHistoryItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMobileAppTroubleshootingHistoryItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'History' for 'MobileAppTroubleshootingEvent': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.History = &output
	}
	return nil
}
