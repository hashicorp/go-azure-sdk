package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ItemActivity{}

type ItemActivity struct {
	// An item was accessed.
	Access *AccessAction `json:"access,omitempty"`

	// Details about when the activity took place. Read-only.
	ActivityDateTime nullable.Type[string] `json:"activityDateTime,omitempty"`

	// Identity of who performed the action. Read-only.
	Actor *IdentitySet `json:"actor,omitempty"`

	// Exposes the driveItem that was the target of this activity.
	DriveItem *DriveItem `json:"driveItem,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ItemActivity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ItemActivity{}

func (s ItemActivity) MarshalJSON() ([]byte, error) {
	type wrapper ItemActivity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ItemActivity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ItemActivity: %+v", err)
	}

	delete(decoded, "activityDateTime")
	delete(decoded, "actor")
	decoded["@odata.type"] = "#microsoft.graph.itemActivity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ItemActivity: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ItemActivity{}

func (s *ItemActivity) UnmarshalJSON(bytes []byte) error {

	var decoded struct {
		Access           *AccessAction         `json:"access,omitempty"`
		ActivityDateTime nullable.Type[string] `json:"activityDateTime,omitempty"`
		DriveItem        *DriveItem            `json:"driveItem,omitempty"`
		Id               *string               `json:"id,omitempty"`
		ODataId          *string               `json:"@odata.id,omitempty"`
		ODataType        *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Access = decoded.Access
	s.ActivityDateTime = decoded.ActivityDateTime
	s.DriveItem = decoded.DriveItem
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ItemActivity into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["actor"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Actor' for 'ItemActivity': %+v", err)
		}
		s.Actor = &impl
	}

	return nil
}
