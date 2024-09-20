package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = TeamsAppInstalledEventMessageDetail{}

type TeamsAppInstalledEventMessageDetail struct {
	// Initiator of the event.
	Initiator IdentitySet `json:"initiator"`

	// Display name of the teamsApp.
	TeamsAppDisplayName nullable.Type[string] `json:"teamsAppDisplayName,omitempty"`

	// Unique identifier of the teamsApp.
	TeamsAppId nullable.Type[string] `json:"teamsAppId,omitempty"`

	// Fields inherited from EventMessageDetail

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s TeamsAppInstalledEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamsAppInstalledEventMessageDetail{}

func (s TeamsAppInstalledEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper TeamsAppInstalledEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamsAppInstalledEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamsAppInstalledEventMessageDetail: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.teamsAppInstalledEventMessageDetail"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamsAppInstalledEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamsAppInstalledEventMessageDetail{}

func (s *TeamsAppInstalledEventMessageDetail) UnmarshalJSON(bytes []byte) error {

	var decoded struct {
		TeamsAppDisplayName nullable.Type[string] `json:"teamsAppDisplayName,omitempty"`
		TeamsAppId          nullable.Type[string] `json:"teamsAppId,omitempty"`
		ODataId             *string               `json:"@odata.id,omitempty"`
		ODataType           *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.TeamsAppDisplayName = decoded.TeamsAppDisplayName
	s.TeamsAppId = decoded.TeamsAppId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamsAppInstalledEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'TeamsAppInstalledEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}

	return nil
}
