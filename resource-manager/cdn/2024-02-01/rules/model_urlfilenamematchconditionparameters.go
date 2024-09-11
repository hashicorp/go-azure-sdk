package rules

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeliveryRuleConditionParameters = UrlFileNameMatchConditionParameters{}

type UrlFileNameMatchConditionParameters struct {
	MatchValues     *[]string           `json:"matchValues,omitempty"`
	NegateCondition *bool               `json:"negateCondition,omitempty"`
	Operator        UrlFileNameOperator `json:"operator"`
	Transforms      *[]Transform        `json:"transforms,omitempty"`

	// Fields inherited from DeliveryRuleConditionParameters
}

var _ json.Marshaler = UrlFileNameMatchConditionParameters{}

func (s UrlFileNameMatchConditionParameters) MarshalJSON() ([]byte, error) {
	type wrapper UrlFileNameMatchConditionParameters
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UrlFileNameMatchConditionParameters: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UrlFileNameMatchConditionParameters: %+v", err)
	}
	decoded["typeName"] = "DeliveryRuleUrlFilenameConditionParameters"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UrlFileNameMatchConditionParameters: %+v", err)
	}

	return encoded, nil
}
