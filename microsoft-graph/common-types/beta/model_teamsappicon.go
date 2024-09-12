package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamsAppIcon{}

type TeamsAppIcon struct {
	// The contents of the app icon if the icon is hosted within the Teams infrastructure.
	HostedContent *TeamworkHostedContent `json:"hostedContent,omitempty"`

	// The web URL that can be used for downloading the image.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s TeamsAppIcon) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamsAppIcon{}

func (s TeamsAppIcon) MarshalJSON() ([]byte, error) {
	type wrapper TeamsAppIcon
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamsAppIcon: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamsAppIcon: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.teamsAppIcon"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamsAppIcon: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamsAppIcon{}

func (s *TeamsAppIcon) UnmarshalJSON(bytes []byte) error {
	type alias TeamsAppIcon
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into TeamsAppIcon: %+v", err)
	}

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.WebUrl = decoded.WebUrl

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamsAppIcon into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["hostedContent"]; ok {
		impl, err := UnmarshalTeamworkHostedContentImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'HostedContent' for 'TeamsAppIcon': %+v", err)
		}
		s.HostedContent = &impl
	}
	return nil
}
