package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundFlow interface {
	Entity
	IndustryDataIndustryDataActivity
	IndustryDataInboundFlow() BaseIndustryDataInboundFlowImpl
}

var _ IndustryDataInboundFlow = BaseIndustryDataInboundFlowImpl{}

type BaseIndustryDataInboundFlowImpl struct {
	DataConnector *IndustryDataIndustryDataConnector `json:"dataConnector,omitempty"`
	DataDomain    *IndustryDataInboundDomain         `json:"dataDomain,omitempty"`

	// The start of the time window when the flow is allowed to run. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`

	// The end of the time window when the flow is allowed to run. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	Year *IndustryDataYearTimePeriodDefinition `json:"year,omitempty"`

	// Fields inherited from IndustryDataIndustryDataActivity

	// The name of the activity. Maximum supported length is 100 characters.
	DisplayName *string `json:"displayName,omitempty"`

	ReadinessStatus *IndustryDataReadinessStatus `json:"readinessStatus,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseIndustryDataInboundFlowImpl) IndustryDataInboundFlow() BaseIndustryDataInboundFlowImpl {
	return s
}

func (s BaseIndustryDataInboundFlowImpl) IndustryDataIndustryDataActivity() BaseIndustryDataIndustryDataActivityImpl {
	return BaseIndustryDataIndustryDataActivityImpl{
		DisplayName:     s.DisplayName,
		ReadinessStatus: s.ReadinessStatus,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BaseIndustryDataInboundFlowImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IndustryDataInboundFlow = RawIndustryDataInboundFlowImpl{}

// RawIndustryDataInboundFlowImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIndustryDataInboundFlowImpl struct {
	industryDataInboundFlow BaseIndustryDataInboundFlowImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawIndustryDataInboundFlowImpl) IndustryDataInboundFlow() BaseIndustryDataInboundFlowImpl {
	return s.industryDataInboundFlow
}

func (s RawIndustryDataInboundFlowImpl) IndustryDataIndustryDataActivity() BaseIndustryDataIndustryDataActivityImpl {
	return s.industryDataInboundFlow.IndustryDataIndustryDataActivity()
}

func (s RawIndustryDataInboundFlowImpl) Entity() BaseEntityImpl {
	return s.industryDataInboundFlow.Entity()
}

var _ json.Marshaler = BaseIndustryDataInboundFlowImpl{}

func (s BaseIndustryDataInboundFlowImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIndustryDataInboundFlowImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIndustryDataInboundFlowImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIndustryDataInboundFlowImpl: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.industryData.inboundFlow"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIndustryDataInboundFlowImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseIndustryDataInboundFlowImpl{}

func (s *BaseIndustryDataInboundFlowImpl) UnmarshalJSON(bytes []byte) error {

	var decoded struct {
		DataDomain         *IndustryDataInboundDomain            `json:"dataDomain,omitempty"`
		EffectiveDateTime  *string                               `json:"effectiveDateTime,omitempty"`
		ExpirationDateTime nullable.Type[string]                 `json:"expirationDateTime,omitempty"`
		Year               *IndustryDataYearTimePeriodDefinition `json:"year,omitempty"`
		DisplayName        *string                               `json:"displayName,omitempty"`
		ReadinessStatus    *IndustryDataReadinessStatus          `json:"readinessStatus,omitempty"`
		Id                 *string                               `json:"id,omitempty"`
		ODataId            *string                               `json:"@odata.id,omitempty"`
		ODataType          *string                               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DataDomain = decoded.DataDomain
	s.EffectiveDateTime = decoded.EffectiveDateTime
	s.ExpirationDateTime = decoded.ExpirationDateTime
	s.Year = decoded.Year
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ReadinessStatus = decoded.ReadinessStatus

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseIndustryDataInboundFlowImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["dataConnector"]; ok {
		impl, err := UnmarshalIndustryDataIndustryDataConnectorImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DataConnector' for 'BaseIndustryDataInboundFlowImpl': %+v", err)
		}
		s.DataConnector = &impl
	}

	return nil
}

func UnmarshalIndustryDataInboundFlowImplementation(input []byte) (IndustryDataInboundFlow, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataInboundFlow into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.inboundApiFlow") {
		var out IndustryDataInboundApiFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataInboundApiFlow: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.inboundFileFlow") {
		var out IndustryDataInboundFileFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataInboundFileFlow: %+v", err)
		}
		return out, nil
	}

	var parent BaseIndustryDataInboundFlowImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIndustryDataInboundFlowImpl: %+v", err)
	}

	return RawIndustryDataInboundFlowImpl{
		industryDataInboundFlow: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
