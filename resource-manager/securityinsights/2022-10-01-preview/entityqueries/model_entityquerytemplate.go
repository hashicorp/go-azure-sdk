package entityqueries

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityQueryTemplate interface {
	EntityQueryTemplate() BaseEntityQueryTemplateImpl
}

var _ EntityQueryTemplate = BaseEntityQueryTemplateImpl{}

type BaseEntityQueryTemplateImpl struct {
	Id         *string                 `json:"id,omitempty"`
	Kind       EntityQueryTemplateKind `json:"kind"`
	Name       *string                 `json:"name,omitempty"`
	SystemData *systemdata.SystemData  `json:"systemData,omitempty"`
	Type       *string                 `json:"type,omitempty"`
}

func (s BaseEntityQueryTemplateImpl) EntityQueryTemplate() BaseEntityQueryTemplateImpl {
	return s
}

var _ EntityQueryTemplate = RawEntityQueryTemplateImpl{}

// RawEntityQueryTemplateImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawEntityQueryTemplateImpl struct {
	entityQueryTemplate BaseEntityQueryTemplateImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawEntityQueryTemplateImpl) EntityQueryTemplate() BaseEntityQueryTemplateImpl {
	return s.entityQueryTemplate
}

func (s RawEntityQueryTemplateImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalEntityQueryTemplateImplementation(input []byte) (EntityQueryTemplate, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EntityQueryTemplate into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Activity") {
		var out ActivityEntityQueryTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActivityEntityQueryTemplate: %+v", err)
		}
		return out, nil
	}

	var parent BaseEntityQueryTemplateImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEntityQueryTemplateImpl: %+v", err)
	}

	return RawEntityQueryTemplateImpl{
		entityQueryTemplate: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
