package benefitutilizationsummaries

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BenefitUtilizationSummary = SavingsPlanUtilizationSummary{}

type SavingsPlanUtilizationSummary struct {
	Properties *SavingsPlanUtilizationSummaryProperties `json:"properties,omitempty"`

	// Fields inherited from BenefitUtilizationSummary
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

var _ json.Marshaler = SavingsPlanUtilizationSummary{}

func (s SavingsPlanUtilizationSummary) MarshalJSON() ([]byte, error) {
	type wrapper SavingsPlanUtilizationSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SavingsPlanUtilizationSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SavingsPlanUtilizationSummary: %+v", err)
	}
	decoded["kind"] = "SavingsPlan"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SavingsPlanUtilizationSummary: %+v", err)
	}

	return encoded, nil
}
