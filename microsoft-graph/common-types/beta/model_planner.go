package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Planner{}

type Planner struct {
	// Read-only. Nullable. Returns a collection of the specified buckets
	Buckets *[]PlannerBucket `json:"buckets,omitempty"`

	// Read-only. Nullable. Returns a collection of the specified plans
	Plans *[]PlannerPlan `json:"plans,omitempty"`

	// Read-only. Nullable. Returns a collection of the specified rosters
	Rosters *[]PlannerRoster `json:"rosters,omitempty"`

	// Read-only. Nullable. Returns a collection of the specified tasks
	Tasks *[]PlannerTask `json:"tasks,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s Planner) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Planner{}

func (s Planner) MarshalJSON() ([]byte, error) {
	type wrapper Planner
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Planner: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Planner: %+v", err)
	}

	delete(decoded, "buckets")
	delete(decoded, "plans")
	delete(decoded, "rosters")
	delete(decoded, "tasks")
	decoded["@odata.type"] = "#microsoft.graph.planner"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Planner: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Planner{}

func (s *Planner) UnmarshalJSON(bytes []byte) error {
	type alias Planner
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into Planner: %+v", err)
	}

	s.Buckets = decoded.Buckets
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Plans = decoded.Plans
	s.Rosters = decoded.Rosters

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Planner into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["tasks"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Tasks into list []json.RawMessage: %+v", err)
		}

		output := make([]PlannerTask, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalPlannerTaskImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Tasks' for 'Planner': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Tasks = &output
	}
	return nil
}
