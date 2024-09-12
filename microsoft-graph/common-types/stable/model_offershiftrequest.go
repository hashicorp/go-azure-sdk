package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfferShiftRequest interface {
	Entity
	ChangeTrackedEntity
	ScheduleChangeRequest
	OfferShiftRequest() BaseOfferShiftRequestImpl
}

var _ OfferShiftRequest = BaseOfferShiftRequestImpl{}

type BaseOfferShiftRequestImpl struct {
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	RecipientActionDateTime nullable.Type[string] `json:"recipientActionDateTime,omitempty"`

	// Custom message sent by recipient of the offer shift request.
	RecipientActionMessage nullable.Type[string] `json:"recipientActionMessage,omitempty"`

	// User ID of the recipient of the offer shift request.
	RecipientUserId nullable.Type[string] `json:"recipientUserId,omitempty"`

	// User ID of the sender of the offer shift request.
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

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Identity of the person who last modified the entity.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseOfferShiftRequestImpl) OfferShiftRequest() BaseOfferShiftRequestImpl {
	return s
}

func (s BaseOfferShiftRequestImpl) ScheduleChangeRequest() BaseScheduleChangeRequestImpl {
	return BaseScheduleChangeRequestImpl{
		AssignedTo:            s.AssignedTo,
		ManagerActionDateTime: s.ManagerActionDateTime,
		ManagerActionMessage:  s.ManagerActionMessage,
		ManagerUserId:         s.ManagerUserId,
		SenderDateTime:        s.SenderDateTime,
		SenderMessage:         s.SenderMessage,
		SenderUserId:          s.SenderUserId,
		State:                 s.State,
		CreatedDateTime:       s.CreatedDateTime,
		LastModifiedBy:        s.LastModifiedBy,
		LastModifiedDateTime:  s.LastModifiedDateTime,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s BaseOfferShiftRequestImpl) ChangeTrackedEntity() BaseChangeTrackedEntityImpl {
	return BaseChangeTrackedEntityImpl{
		CreatedDateTime:      s.CreatedDateTime,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s BaseOfferShiftRequestImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ OfferShiftRequest = RawOfferShiftRequestImpl{}

// RawOfferShiftRequestImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOfferShiftRequestImpl struct {
	offerShiftRequest BaseOfferShiftRequestImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawOfferShiftRequestImpl) OfferShiftRequest() BaseOfferShiftRequestImpl {
	return s.offerShiftRequest
}

func (s RawOfferShiftRequestImpl) ScheduleChangeRequest() BaseScheduleChangeRequestImpl {
	return s.offerShiftRequest.ScheduleChangeRequest()
}

func (s RawOfferShiftRequestImpl) ChangeTrackedEntity() BaseChangeTrackedEntityImpl {
	return s.offerShiftRequest.ChangeTrackedEntity()
}

func (s RawOfferShiftRequestImpl) Entity() BaseEntityImpl {
	return s.offerShiftRequest.Entity()
}

var _ json.Marshaler = BaseOfferShiftRequestImpl{}

func (s BaseOfferShiftRequestImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseOfferShiftRequestImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseOfferShiftRequestImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseOfferShiftRequestImpl: %+v", err)
	}

	delete(decoded, "recipientActionDateTime")
	decoded["@odata.type"] = "#microsoft.graph.offerShiftRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseOfferShiftRequestImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseOfferShiftRequestImpl{}

func (s *BaseOfferShiftRequestImpl) UnmarshalJSON(bytes []byte) error {
	type alias BaseOfferShiftRequestImpl
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into BaseOfferShiftRequestImpl: %+v", err)
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
	s.RecipientUserId = decoded.RecipientUserId
	s.SenderDateTime = decoded.SenderDateTime
	s.SenderMessage = decoded.SenderMessage
	s.SenderShiftId = decoded.SenderShiftId
	s.SenderUserId = decoded.SenderUserId
	s.State = decoded.State

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseOfferShiftRequestImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseOfferShiftRequestImpl': %+v", err)
		}
		s.LastModifiedBy = &impl
	}
	return nil
}

func UnmarshalOfferShiftRequestImplementation(input []byte) (OfferShiftRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OfferShiftRequest into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.swapShiftsChangeRequest") {
		var out SwapShiftsChangeRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SwapShiftsChangeRequest: %+v", err)
		}
		return out, nil
	}

	var parent BaseOfferShiftRequestImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOfferShiftRequestImpl: %+v", err)
	}

	return RawOfferShiftRequestImpl{
		offerShiftRequest: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
