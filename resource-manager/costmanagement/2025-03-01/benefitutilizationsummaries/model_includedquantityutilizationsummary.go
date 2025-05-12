package benefitutilizationsummaries

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BenefitUtilizationSummary = IncludedQuantityUtilizationSummary{}

type IncludedQuantityUtilizationSummary struct {
	Properties *IncludedQuantityUtilizationSummaryProperties `json:"properties,omitempty"`

	// Fields inherited from BenefitUtilizationSummary

	Id   *string     `json:"id,omitempty"`
	Kind BenefitKind `json:"kind"`
	Name *string     `json:"name,omitempty"`
	Type *string     `json:"type,omitempty"`
}

func (s IncludedQuantityUtilizationSummary) BenefitUtilizationSummary() BaseBenefitUtilizationSummaryImpl {
	return BaseBenefitUtilizationSummaryImpl{
		Id:   s.Id,
		Kind: s.Kind,
		Name: s.Name,
		Type: s.Type,
	}
}

var _ json.Marshaler = IncludedQuantityUtilizationSummary{}

func (s IncludedQuantityUtilizationSummary) MarshalJSON() ([]byte, error) {
	type wrapper IncludedQuantityUtilizationSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IncludedQuantityUtilizationSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IncludedQuantityUtilizationSummary: %+v", err)
	}

	decoded["kind"] = "IncludedQuantity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IncludedQuantityUtilizationSummary: %+v", err)
	}

	return encoded, nil
}
