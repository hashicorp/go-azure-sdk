package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ServiceAnnouncement{}

type ServiceAnnouncement struct {
	// A collection of service health information for tenant. This property is a contained navigation property, it is
	// nullable and readonly.
	HealthOverviews *[]ServiceHealth `json:"healthOverviews,omitempty"`

	// A collection of service issues for tenant. This property is a contained navigation property, it is nullable and
	// readonly.
	Issues *[]ServiceHealthIssue `json:"issues,omitempty"`

	// A collection of service messages for tenant. This property is a contained navigation property, it is nullable and
	// readonly.
	Messages *[]ServiceUpdateMessage `json:"messages,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ServiceAnnouncement) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ServiceAnnouncement{}

func (s ServiceAnnouncement) MarshalJSON() ([]byte, error) {
	type wrapper ServiceAnnouncement
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceAnnouncement: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceAnnouncement: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.serviceAnnouncement"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceAnnouncement: %+v", err)
	}

	return encoded, nil
}
