package wafloganalytics

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WafGranularity string

const (
	WafGranularityPOneD   WafGranularity = "P1D"
	WafGranularityPTFiveM WafGranularity = "PT5M"
	WafGranularityPTOneH  WafGranularity = "PT1H"
)

func PossibleValuesForWafGranularity() []string {
	return []string{
		string(WafGranularityPOneD),
		string(WafGranularityPTFiveM),
		string(WafGranularityPTOneH),
	}
}

func (s *WafGranularity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWafGranularity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWafGranularity(input string) (*WafGranularity, error) {
	vals := map[string]WafGranularity{
		"p1d":  WafGranularityPOneD,
		"pt5m": WafGranularityPTFiveM,
		"pt1h": WafGranularityPTOneH,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WafGranularity(input)
	return &out, nil
}

type WafMetricsGranularity string

const (
	WafMetricsGranularityPOneD   WafMetricsGranularity = "P1D"
	WafMetricsGranularityPTFiveM WafMetricsGranularity = "PT5M"
	WafMetricsGranularityPTOneH  WafMetricsGranularity = "PT1H"
)

func PossibleValuesForWafMetricsGranularity() []string {
	return []string{
		string(WafMetricsGranularityPOneD),
		string(WafMetricsGranularityPTFiveM),
		string(WafMetricsGranularityPTOneH),
	}
}

func (s *WafMetricsGranularity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWafMetricsGranularity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWafMetricsGranularity(input string) (*WafMetricsGranularity, error) {
	vals := map[string]WafMetricsGranularity{
		"p1d":  WafMetricsGranularityPOneD,
		"pt5m": WafMetricsGranularityPTFiveM,
		"pt1h": WafMetricsGranularityPTOneH,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WafMetricsGranularity(input)
	return &out, nil
}

type WafMetricsSeriesUnit string

const (
	WafMetricsSeriesUnitCount WafMetricsSeriesUnit = "count"
)

func PossibleValuesForWafMetricsSeriesUnit() []string {
	return []string{
		string(WafMetricsSeriesUnitCount),
	}
}

func (s *WafMetricsSeriesUnit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWafMetricsSeriesUnit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWafMetricsSeriesUnit(input string) (*WafMetricsSeriesUnit, error) {
	vals := map[string]WafMetricsSeriesUnit{
		"count": WafMetricsSeriesUnitCount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WafMetricsSeriesUnit(input)
	return &out, nil
}
