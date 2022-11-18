package benefitutilizationsummaries

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitUtilizationSummary interface {
}

func unmarshalBenefitUtilizationSummaryImplementation(input []byte) (BenefitUtilizationSummary, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BenefitUtilizationSummary into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "IncludedQuantity") {
		var out IncludedQuantityUtilizationSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IncludedQuantityUtilizationSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SavingsPlan") {
		var out SavingsPlanUtilizationSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SavingsPlanUtilizationSummary: %+v", err)
		}
		return out, nil
	}

	type RawBenefitUtilizationSummaryImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawBenefitUtilizationSummaryImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
