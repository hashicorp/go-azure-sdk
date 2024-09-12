package endpoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeliveryRuleConditionParameters = UrlPathMatchConditionParameters{}

type UrlPathMatchConditionParameters struct {
	MatchValues     *[]string       `json:"matchValues,omitempty"`
	NegateCondition *bool           `json:"negateCondition,omitempty"`
	Operator        UrlPathOperator `json:"operator"`
	Transforms      *[]Transform    `json:"transforms,omitempty"`

	// Fields inherited from DeliveryRuleConditionParameters
}

var _ json.Marshaler = UrlPathMatchConditionParameters{}

func (s UrlPathMatchConditionParameters) MarshalJSON() ([]byte, error) {
	type wrapper UrlPathMatchConditionParameters
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UrlPathMatchConditionParameters: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UrlPathMatchConditionParameters: %+v", err)
	}
	decoded["typeName"] = "DeliveryRuleUrlPathMatchConditionParameters"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UrlPathMatchConditionParameters: %+v", err)
	}

	return encoded, nil
}
