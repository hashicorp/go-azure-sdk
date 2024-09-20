package endpoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeliveryRuleConditionParameters = UrlFileExtensionMatchConditionParameters{}

type UrlFileExtensionMatchConditionParameters struct {
	MatchValues     *[]string                `json:"matchValues,omitempty"`
	NegateCondition *bool                    `json:"negateCondition,omitempty"`
	Operator        UrlFileExtensionOperator `json:"operator"`
	Transforms      *[]Transform             `json:"transforms,omitempty"`

	// Fields inherited from DeliveryRuleConditionParameters

	TypeName DeliveryRuleConditionParametersType `json:"typeName"`
}

func (s UrlFileExtensionMatchConditionParameters) DeliveryRuleConditionParameters() BaseDeliveryRuleConditionParametersImpl {
	return BaseDeliveryRuleConditionParametersImpl{
		TypeName: s.TypeName,
	}
}

var _ json.Marshaler = UrlFileExtensionMatchConditionParameters{}

func (s UrlFileExtensionMatchConditionParameters) MarshalJSON() ([]byte, error) {
	type wrapper UrlFileExtensionMatchConditionParameters
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UrlFileExtensionMatchConditionParameters: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UrlFileExtensionMatchConditionParameters: %+v", err)
	}

	decoded["typeName"] = "DeliveryRuleUrlFileExtensionMatchConditionParameters"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UrlFileExtensionMatchConditionParameters: %+v", err)
	}

	return encoded, nil
}
