package benefitutilizationsummaries

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitUtilizationSummary interface {
	BenefitUtilizationSummary() BaseBenefitUtilizationSummaryImpl
}

var _ BenefitUtilizationSummary = BaseBenefitUtilizationSummaryImpl{}

type BaseBenefitUtilizationSummaryImpl struct {
	Id   *string     `json:"id,omitempty"`
	Kind BenefitKind `json:"kind"`
	Name *string     `json:"name,omitempty"`
	Type *string     `json:"type,omitempty"`
}

func (s BaseBenefitUtilizationSummaryImpl) BenefitUtilizationSummary() BaseBenefitUtilizationSummaryImpl {
	return s
}

var _ BenefitUtilizationSummary = RawBenefitUtilizationSummaryImpl{}

// RawBenefitUtilizationSummaryImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBenefitUtilizationSummaryImpl struct {
	benefitUtilizationSummary BaseBenefitUtilizationSummaryImpl
	Type                      string
	Values                    map[string]interface{}
}

func (s RawBenefitUtilizationSummaryImpl) BenefitUtilizationSummary() BaseBenefitUtilizationSummaryImpl {
	return s.benefitUtilizationSummary
}

func UnmarshalBenefitUtilizationSummaryImplementation(input []byte) (BenefitUtilizationSummary, error) {
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

	var parent BaseBenefitUtilizationSummaryImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBenefitUtilizationSummaryImpl: %+v", err)
	}

	return RawBenefitUtilizationSummaryImpl{
		benefitUtilizationSummary: parent,
		Type:                      value,
		Values:                    temp,
	}, nil

}
