package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessRelatedResource interface {
	NetworkaccessRelatedResource() BaseNetworkaccessRelatedResourceImpl
}

var _ NetworkaccessRelatedResource = BaseNetworkaccessRelatedResourceImpl{}

type BaseNetworkaccessRelatedResourceImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseNetworkaccessRelatedResourceImpl) NetworkaccessRelatedResource() BaseNetworkaccessRelatedResourceImpl {
	return s
}

var _ NetworkaccessRelatedResource = RawNetworkaccessRelatedResourceImpl{}

// RawNetworkaccessRelatedResourceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawNetworkaccessRelatedResourceImpl struct {
	networkaccessRelatedResource BaseNetworkaccessRelatedResourceImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawNetworkaccessRelatedResourceImpl) NetworkaccessRelatedResource() BaseNetworkaccessRelatedResourceImpl {
	return s.networkaccessRelatedResource
}

func UnmarshalNetworkaccessRelatedResourceImplementation(input []byte) (NetworkaccessRelatedResource, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessRelatedResource into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.relatedDestination") {
		var out NetworkaccessRelatedDestination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessRelatedDestination: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.relatedDevice") {
		var out NetworkaccessRelatedDevice
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessRelatedDevice: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.relatedMalware") {
		var out NetworkaccessRelatedMalware
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessRelatedMalware: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.relatedProcess") {
		var out NetworkaccessRelatedProcess
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessRelatedProcess: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.relatedRemoteNetwork") {
		var out NetworkaccessRelatedRemoteNetwork
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessRelatedRemoteNetwork: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.relatedTenant") {
		var out NetworkaccessRelatedTenant
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessRelatedTenant: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.relatedThreatIntelligence") {
		var out NetworkaccessRelatedThreatIntelligence
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessRelatedThreatIntelligence: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.relatedToken") {
		var out NetworkaccessRelatedToken
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessRelatedToken: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.relatedUser") {
		var out NetworkaccessRelatedUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessRelatedUser: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.relatedWebCategory") {
		var out NetworkaccessRelatedWebCategory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessRelatedWebCategory: %+v", err)
		}
		return out, nil
	}

	var parent BaseNetworkaccessRelatedResourceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseNetworkaccessRelatedResourceImpl: %+v", err)
	}

	return RawNetworkaccessRelatedResourceImpl{
		networkaccessRelatedResource: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
