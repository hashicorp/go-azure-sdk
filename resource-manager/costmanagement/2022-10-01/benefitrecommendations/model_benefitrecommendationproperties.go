package benefitrecommendations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitRecommendationProperties interface {
}

func unmarshalBenefitRecommendationPropertiesImplementation(input []byte) (BenefitRecommendationProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BenefitRecommendationProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["scope"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Shared") {
		var out SharedScopeBenefitRecommendationProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharedScopeBenefitRecommendationProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Single") {
		var out SingleScopeBenefitRecommendationProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SingleScopeBenefitRecommendationProperties: %+v", err)
		}
		return out, nil
	}

	type RawBenefitRecommendationPropertiesImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawBenefitRecommendationPropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
