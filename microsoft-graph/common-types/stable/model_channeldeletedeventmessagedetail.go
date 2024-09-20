package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = ChannelDeletedEventMessageDetail{}

type ChannelDeletedEventMessageDetail struct {
	// Display name of the channel.
	ChannelDisplayName nullable.Type[string] `json:"channelDisplayName,omitempty"`

	// Unique identifier of the channel.
	ChannelId nullable.Type[string] `json:"channelId,omitempty"`

	// Initiator of the event.
	Initiator IdentitySet `json:"initiator"`

	// Fields inherited from EventMessageDetail

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ChannelDeletedEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ChannelDeletedEventMessageDetail{}

func (s ChannelDeletedEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper ChannelDeletedEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChannelDeletedEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChannelDeletedEventMessageDetail: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.channelDeletedEventMessageDetail"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChannelDeletedEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ChannelDeletedEventMessageDetail{}

func (s *ChannelDeletedEventMessageDetail) UnmarshalJSON(bytes []byte) error {

	var decoded struct {
		ChannelDisplayName nullable.Type[string] `json:"channelDisplayName,omitempty"`
		ChannelId          nullable.Type[string] `json:"channelId,omitempty"`
		ODataId            *string               `json:"@odata.id,omitempty"`
		ODataType          *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ChannelDisplayName = decoded.ChannelDisplayName
	s.ChannelId = decoded.ChannelId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChannelDeletedEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'ChannelDeletedEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}

	return nil
}
