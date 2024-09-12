package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeGraphInsights interface {
	Entity
	OfficeGraphInsights() BaseOfficeGraphInsightsImpl
}

var _ OfficeGraphInsights = BaseOfficeGraphInsightsImpl{}

type BaseOfficeGraphInsightsImpl struct {
	// Calculated relationship that identifies documents shared with or by the user. This includes URLs, file attachments,
	// and reference attachments to OneDrive for work or school and SharePoint files found in Outlook messages and meetings.
	// This also includes URLs and reference attachments to Teams conversations. Ordered by recency of share.
	Shared *[]SharedInsight `json:"shared,omitempty"`

	// Calculated relationship that identifies documents trending around a user. Trending documents are calculated based on
	// activity of the user's closest network of people and include files stored in OneDrive for work or school and
	// SharePoint. Trending insights help the user to discover potentially useful content that the user has access to, but
	// has never viewed before.
	Trending *[]Trending `json:"trending,omitempty"`

	// Calculated relationship that identifies the latest documents viewed or modified by a user, including OneDrive for
	// work or school and SharePoint documents, ranked by recency of use.
	Used *[]UsedInsight `json:"used,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseOfficeGraphInsightsImpl) OfficeGraphInsights() BaseOfficeGraphInsightsImpl {
	return s
}

func (s BaseOfficeGraphInsightsImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ OfficeGraphInsights = RawOfficeGraphInsightsImpl{}

// RawOfficeGraphInsightsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOfficeGraphInsightsImpl struct {
	officeGraphInsights BaseOfficeGraphInsightsImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawOfficeGraphInsightsImpl) OfficeGraphInsights() BaseOfficeGraphInsightsImpl {
	return s.officeGraphInsights
}

func (s RawOfficeGraphInsightsImpl) Entity() BaseEntityImpl {
	return s.officeGraphInsights.Entity()
}

var _ json.Marshaler = BaseOfficeGraphInsightsImpl{}

func (s BaseOfficeGraphInsightsImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseOfficeGraphInsightsImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseOfficeGraphInsightsImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseOfficeGraphInsightsImpl: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.officeGraphInsights"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseOfficeGraphInsightsImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalOfficeGraphInsightsImplementation(input []byte) (OfficeGraphInsights, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OfficeGraphInsights into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemInsights") {
		var out ItemInsights
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemInsights: %+v", err)
		}
		return out, nil
	}

	var parent BaseOfficeGraphInsightsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOfficeGraphInsightsImpl: %+v", err)
	}

	return RawOfficeGraphInsightsImpl{
		officeGraphInsights: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
