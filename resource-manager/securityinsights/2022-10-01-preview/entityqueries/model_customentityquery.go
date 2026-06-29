package entityqueries

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomEntityQuery interface {
	CustomEntityQuery() BaseCustomEntityQueryImpl
}

var _ CustomEntityQuery = BaseCustomEntityQueryImpl{}

type BaseCustomEntityQueryImpl struct {
	Etag       *string                `json:"etag,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Kind       CustomEntityQueryKind  `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s BaseCustomEntityQueryImpl) CustomEntityQuery() BaseCustomEntityQueryImpl {
	return s
}

var _ CustomEntityQuery = RawCustomEntityQueryImpl{}

// RawCustomEntityQueryImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawCustomEntityQueryImpl struct {
	customEntityQuery BaseCustomEntityQueryImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawCustomEntityQueryImpl) CustomEntityQuery() BaseCustomEntityQueryImpl {
	return s.customEntityQuery
}

func (s RawCustomEntityQueryImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalCustomEntityQueryImplementation(input []byte) (CustomEntityQuery, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomEntityQuery into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Activity") {
		var out ActivityCustomEntityQuery
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActivityCustomEntityQuery: %+v", err)
		}
		return out, nil
	}

	var parent BaseCustomEntityQueryImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomEntityQueryImpl: %+v", err)
	}

	return RawCustomEntityQueryImpl{
		customEntityQuery: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
