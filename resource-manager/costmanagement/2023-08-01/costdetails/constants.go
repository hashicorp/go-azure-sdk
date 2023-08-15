package costdetails

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CostDetailsDataFormat string

const (
	CostDetailsDataFormatCsv CostDetailsDataFormat = "Csv"
)

func PossibleValuesForCostDetailsDataFormat() []string {
	return []string{
		string(CostDetailsDataFormatCsv),
	}
}

func (s *CostDetailsDataFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCostDetailsDataFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCostDetailsDataFormat(input string) (*CostDetailsDataFormat, error) {
	vals := map[string]CostDetailsDataFormat{
		"csv": CostDetailsDataFormatCsv,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CostDetailsDataFormat(input)
	return &out, nil
}

type CostDetailsMetricType string

const (
	CostDetailsMetricTypeActualCost    CostDetailsMetricType = "ActualCost"
	CostDetailsMetricTypeAmortizedCost CostDetailsMetricType = "AmortizedCost"
)

func PossibleValuesForCostDetailsMetricType() []string {
	return []string{
		string(CostDetailsMetricTypeActualCost),
		string(CostDetailsMetricTypeAmortizedCost),
	}
}

func (s *CostDetailsMetricType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCostDetailsMetricType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCostDetailsMetricType(input string) (*CostDetailsMetricType, error) {
	vals := map[string]CostDetailsMetricType{
		"actualcost":    CostDetailsMetricTypeActualCost,
		"amortizedcost": CostDetailsMetricTypeAmortizedCost,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CostDetailsMetricType(input)
	return &out, nil
}

type CostDetailsStatusType string

const (
	CostDetailsStatusTypeCompleted   CostDetailsStatusType = "Completed"
	CostDetailsStatusTypeFailed      CostDetailsStatusType = "Failed"
	CostDetailsStatusTypeNoDataFound CostDetailsStatusType = "NoDataFound"
)

func PossibleValuesForCostDetailsStatusType() []string {
	return []string{
		string(CostDetailsStatusTypeCompleted),
		string(CostDetailsStatusTypeFailed),
		string(CostDetailsStatusTypeNoDataFound),
	}
}

func (s *CostDetailsStatusType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCostDetailsStatusType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCostDetailsStatusType(input string) (*CostDetailsStatusType, error) {
	vals := map[string]CostDetailsStatusType{
		"completed":   CostDetailsStatusTypeCompleted,
		"failed":      CostDetailsStatusTypeFailed,
		"nodatafound": CostDetailsStatusTypeNoDataFound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CostDetailsStatusType(input)
	return &out, nil
}
