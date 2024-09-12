package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OfferShiftRequest = SwapShiftsChangeRequest{}

type SwapShiftsChangeRequest struct {
	// Shift ID for the recipient user with whom the request is to swap.
	RecipientShiftId nullable.Type[string] `json:"recipientShiftId,omitempty"`

	// Fields inherited from OfferShiftRequest

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	RecipientActionDateTime nullable.Type[string] `json:"recipientActionDateTime,omitempty"`

	// Custom message sent by recipient of the offer shift request.
	RecipientActionMessage nullable.Type[string] `json:"recipientActionMessage,omitempty"`

	// User id of the recipient of the offer shift request.
	RecipientUserId nullable.Type[string] `json:"recipientUserId,omitempty"`

	// User id of the sender of the offer shift request.
	SenderShiftId nullable.Type[string] `json:"senderShiftId,omitempty"`

	// Fields inherited from ScheduleChangeRequest

	AssignedTo            *ScheduleChangeRequestActor `json:"assignedTo,omitempty"`
	ManagerActionDateTime nullable.Type[string]       `json:"managerActionDateTime,omitempty"`
	ManagerActionMessage  nullable.Type[string]       `json:"managerActionMessage,omitempty"`
	ManagerUserId         nullable.Type[string]       `json:"managerUserId,omitempty"`
	SenderDateTime        nullable.Type[string]       `json:"senderDateTime,omitempty"`
	SenderMessage         nullable.Type[string]       `json:"senderMessage,omitempty"`
	SenderUserId          nullable.Type[string]       `json:"senderUserId,omitempty"`
	State                 *ScheduleChangeState        `json:"state,omitempty"`

	// Fields inherited from ChangeTrackedEntity

	// Identity of the user who created the entity.
	CreatedBy IdentitySet `json:"createdBy"`

	// The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Identity of the user who last modified the entity.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s SwapShiftsChangeRequest) OfferShiftRequest() BaseOfferShiftRequestImpl {
	return BaseOfferShiftRequestImpl{
		RecipientActionDateTime: s.RecipientActionDateTime,
		RecipientActionMessage:  s.RecipientActionMessage,
		RecipientUserId:         s.RecipientUserId,
		SenderShiftId:           s.SenderShiftId,
		AssignedTo:              s.AssignedTo,
		ManagerActionDateTime:   s.ManagerActionDateTime,
		ManagerActionMessage:    s.ManagerActionMessage,
		ManagerUserId:           s.ManagerUserId,
		SenderDateTime:          s.SenderDateTime,
		SenderMessage:           s.SenderMessage,
		SenderUserId:            s.SenderUserId,
		State:                   s.State,
		CreatedBy:               s.CreatedBy,
		CreatedDateTime:         s.CreatedDateTime,
		LastModifiedBy:          s.LastModifiedBy,
		LastModifiedDateTime:    s.LastModifiedDateTime,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s SwapShiftsChangeRequest) ScheduleChangeRequest() BaseScheduleChangeRequestImpl {
	return BaseScheduleChangeRequestImpl{
		AssignedTo:            s.AssignedTo,
		ManagerActionDateTime: s.ManagerActionDateTime,
		ManagerActionMessage:  s.ManagerActionMessage,
		ManagerUserId:         s.ManagerUserId,
		SenderDateTime:        s.SenderDateTime,
		SenderMessage:         s.SenderMessage,
		SenderUserId:          s.SenderUserId,
		State:                 s.State,
		CreatedBy:             s.CreatedBy,
		CreatedDateTime:       s.CreatedDateTime,
		LastModifiedBy:        s.LastModifiedBy,
		LastModifiedDateTime:  s.LastModifiedDateTime,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s SwapShiftsChangeRequest) ChangeTrackedEntity() BaseChangeTrackedEntityImpl {
	return BaseChangeTrackedEntityImpl{
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SwapShiftsChangeRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SwapShiftsChangeRequest{}

func (s SwapShiftsChangeRequest) MarshalJSON() ([]byte, error) {
	type wrapper SwapShiftsChangeRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SwapShiftsChangeRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SwapShiftsChangeRequest: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.swapShiftsChangeRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SwapShiftsChangeRequest: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SwapShiftsChangeRequest{}

func (s *SwapShiftsChangeRequest) UnmarshalJSON(bytes []byte) error {
	type alias SwapShiftsChangeRequest
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into SwapShiftsChangeRequest: %+v", err)
	}

	s.AssignedTo = decoded.AssignedTo
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ManagerActionDateTime = decoded.ManagerActionDateTime
	s.ManagerActionMessage = decoded.ManagerActionMessage
	s.ManagerUserId = decoded.ManagerUserId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RecipientActionDateTime = decoded.RecipientActionDateTime
	s.RecipientActionMessage = decoded.RecipientActionMessage
	s.RecipientShiftId = decoded.RecipientShiftId
	s.RecipientUserId = decoded.RecipientUserId
	s.SenderDateTime = decoded.SenderDateTime
	s.SenderMessage = decoded.SenderMessage
	s.SenderShiftId = decoded.SenderShiftId
	s.SenderUserId = decoded.SenderUserId
	s.State = decoded.State

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SwapShiftsChangeRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SwapShiftsChangeRequest': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SwapShiftsChangeRequest': %+v", err)
		}
		s.LastModifiedBy = &impl
	}
	return nil
}
