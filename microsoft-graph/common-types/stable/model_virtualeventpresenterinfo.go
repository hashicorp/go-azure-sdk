package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MeetingParticipantInfo = VirtualEventPresenterInfo{}

type VirtualEventPresenterInfo struct {
	PresenterDetails *VirtualEventPresenterDetails `json:"presenterDetails,omitempty"`

	// Fields inherited from MeetingParticipantInfo

	// Identity information of the participant.
	Identity IdentitySet `json:"identity"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the participant's role in the meeting.
	Role *OnlineMeetingRole `json:"role,omitempty"`

	// User principal name of the participant.
	Upn nullable.Type[string] `json:"upn,omitempty"`
}

func (s VirtualEventPresenterInfo) MeetingParticipantInfo() BaseMeetingParticipantInfoImpl {
	return BaseMeetingParticipantInfoImpl{
		Identity:  s.Identity,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		Role:      s.Role,
		Upn:       s.Upn,
	}
}

var _ json.Marshaler = VirtualEventPresenterInfo{}

func (s VirtualEventPresenterInfo) MarshalJSON() ([]byte, error) {
	type wrapper VirtualEventPresenterInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualEventPresenterInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualEventPresenterInfo: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.virtualEventPresenterInfo"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualEventPresenterInfo: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &VirtualEventPresenterInfo{}

func (s *VirtualEventPresenterInfo) UnmarshalJSON(bytes []byte) error {
	type alias VirtualEventPresenterInfo
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into VirtualEventPresenterInfo: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PresenterDetails = decoded.PresenterDetails
	s.Role = decoded.Role
	s.Upn = decoded.Upn

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling VirtualEventPresenterInfo into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'VirtualEventPresenterInfo': %+v", err)
		}
		s.Identity = impl
	}
	return nil
}
