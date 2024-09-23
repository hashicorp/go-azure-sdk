package alerts

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceIdentifier interface {
	ResourceIdentifier() BaseResourceIdentifierImpl
}

var _ ResourceIdentifier = BaseResourceIdentifierImpl{}

type BaseResourceIdentifierImpl struct {
	Type ResourceIdentifierType `json:"type"`
}

func (s BaseResourceIdentifierImpl) ResourceIdentifier() BaseResourceIdentifierImpl {
	return s
}

var _ ResourceIdentifier = RawResourceIdentifierImpl{}

// RawResourceIdentifierImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawResourceIdentifierImpl struct {
	resourceIdentifier BaseResourceIdentifierImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawResourceIdentifierImpl) ResourceIdentifier() BaseResourceIdentifierImpl {
	return s.resourceIdentifier
}

func UnmarshalResourceIdentifierImplementation(input []byte) (ResourceIdentifier, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ResourceIdentifier into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AzureResource") {
		var out AzureResourceIdentifier
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureResourceIdentifier: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "LogAnalytics") {
		var out LogAnalyticsIdentifier
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LogAnalyticsIdentifier: %+v", err)
		}
		return out, nil
	}

	var parent BaseResourceIdentifierImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseResourceIdentifierImpl: %+v", err)
	}

	return RawResourceIdentifierImpl{
		resourceIdentifier: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
