package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ParticipantLeftNotification{}

type ParticipantLeftNotification struct {
	Call *Call `json:"call,omitempty"`

	// ID of the participant under the policy who has left the meeting.
	ParticipantId *string `json:"participantId,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ParticipantLeftNotification) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ParticipantLeftNotification{}

func (s ParticipantLeftNotification) MarshalJSON() ([]byte, error) {
	type wrapper ParticipantLeftNotification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ParticipantLeftNotification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ParticipantLeftNotification: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.participantLeftNotification"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ParticipantLeftNotification: %+v", err)
	}

	return encoded, nil
}
