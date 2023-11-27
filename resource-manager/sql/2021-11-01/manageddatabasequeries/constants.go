package manageddatabasequeries

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryMetricUnitType string

const (
	QueryMetricUnitTypeCount        QueryMetricUnitType = "count"
	QueryMetricUnitTypeKB           QueryMetricUnitType = "KB"
	QueryMetricUnitTypeMicroseconds QueryMetricUnitType = "microseconds"
	QueryMetricUnitTypePercentage   QueryMetricUnitType = "percentage"
)

func PossibleValuesForQueryMetricUnitType() []string {
	return []string{
		string(QueryMetricUnitTypeCount),
		string(QueryMetricUnitTypeKB),
		string(QueryMetricUnitTypeMicroseconds),
		string(QueryMetricUnitTypePercentage),
	}
}

func (s *QueryMetricUnitType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseQueryMetricUnitType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseQueryMetricUnitType(input string) (*QueryMetricUnitType, error) {
	vals := map[string]QueryMetricUnitType{
		"count":        QueryMetricUnitTypeCount,
		"kb":           QueryMetricUnitTypeKB,
		"microseconds": QueryMetricUnitTypeMicroseconds,
		"percentage":   QueryMetricUnitTypePercentage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := QueryMetricUnitType(input)
	return &out, nil
}

type QueryTimeGrainType string

const (
	QueryTimeGrainTypePOneD  QueryTimeGrainType = "P1D"
	QueryTimeGrainTypePTOneH QueryTimeGrainType = "PT1H"
)

func PossibleValuesForQueryTimeGrainType() []string {
	return []string{
		string(QueryTimeGrainTypePOneD),
		string(QueryTimeGrainTypePTOneH),
	}
}

func (s *QueryTimeGrainType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseQueryTimeGrainType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseQueryTimeGrainType(input string) (*QueryTimeGrainType, error) {
	vals := map[string]QueryTimeGrainType{
		"p1d":  QueryTimeGrainTypePOneD,
		"pt1h": QueryTimeGrainTypePTOneH,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := QueryTimeGrainType(input)
	return &out, nil
}
