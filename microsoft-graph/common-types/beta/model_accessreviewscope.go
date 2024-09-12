package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewScope interface {
	AccessReviewScope() BaseAccessReviewScopeImpl
}

var _ AccessReviewScope = BaseAccessReviewScopeImpl{}

type BaseAccessReviewScopeImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseAccessReviewScopeImpl) AccessReviewScope() BaseAccessReviewScopeImpl {
	return s
}

var _ AccessReviewScope = RawAccessReviewScopeImpl{}

// RawAccessReviewScopeImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAccessReviewScopeImpl struct {
	accessReviewScope BaseAccessReviewScopeImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawAccessReviewScopeImpl) AccessReviewScope() BaseAccessReviewScopeImpl {
	return s.accessReviewScope
}

func UnmarshalAccessReviewScopeImplementation(input []byte) (AccessReviewScope, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewScope into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewQueryScope") {
		var out AccessReviewQueryScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewQueryScope: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewReviewerScope") {
		var out AccessReviewReviewerScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewReviewerScope: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.principalResourceMembershipsScope") {
		var out PrincipalResourceMembershipsScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrincipalResourceMembershipsScope: %+v", err)
		}
		return out, nil
	}

	var parent BaseAccessReviewScopeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAccessReviewScopeImpl: %+v", err)
	}

	return RawAccessReviewScopeImpl{
		accessReviewScope: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
