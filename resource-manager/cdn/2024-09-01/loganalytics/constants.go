package loganalytics

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogMetricsGranularity string

const (
	LogMetricsGranularityPOneD   LogMetricsGranularity = "P1D"
	LogMetricsGranularityPTFiveM LogMetricsGranularity = "PT5M"
	LogMetricsGranularityPTOneH  LogMetricsGranularity = "PT1H"
)

func PossibleValuesForLogMetricsGranularity() []string {
	return []string{
		string(LogMetricsGranularityPOneD),
		string(LogMetricsGranularityPTFiveM),
		string(LogMetricsGranularityPTOneH),
	}
}

func (s *LogMetricsGranularity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLogMetricsGranularity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLogMetricsGranularity(input string) (*LogMetricsGranularity, error) {
	vals := map[string]LogMetricsGranularity{
		"p1d":  LogMetricsGranularityPOneD,
		"pt5m": LogMetricsGranularityPTFiveM,
		"pt1h": LogMetricsGranularityPTOneH,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogMetricsGranularity(input)
	return &out, nil
}

type MetricsGranularity string

const (
	MetricsGranularityPOneD   MetricsGranularity = "P1D"
	MetricsGranularityPTFiveM MetricsGranularity = "PT5M"
	MetricsGranularityPTOneH  MetricsGranularity = "PT1H"
)

func PossibleValuesForMetricsGranularity() []string {
	return []string{
		string(MetricsGranularityPOneD),
		string(MetricsGranularityPTFiveM),
		string(MetricsGranularityPTOneH),
	}
}

func (s *MetricsGranularity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMetricsGranularity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMetricsGranularity(input string) (*MetricsGranularity, error) {
	vals := map[string]MetricsGranularity{
		"p1d":  MetricsGranularityPOneD,
		"pt5m": MetricsGranularityPTFiveM,
		"pt1h": MetricsGranularityPTOneH,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MetricsGranularity(input)
	return &out, nil
}

type MetricsSeriesUnit string

const (
	MetricsSeriesUnitBitsPerSecond MetricsSeriesUnit = "bitsPerSecond"
	MetricsSeriesUnitBytes         MetricsSeriesUnit = "bytes"
	MetricsSeriesUnitCount         MetricsSeriesUnit = "count"
	MetricsSeriesUnitMilliSeconds  MetricsSeriesUnit = "milliSeconds"
)

func PossibleValuesForMetricsSeriesUnit() []string {
	return []string{
		string(MetricsSeriesUnitBitsPerSecond),
		string(MetricsSeriesUnitBytes),
		string(MetricsSeriesUnitCount),
		string(MetricsSeriesUnitMilliSeconds),
	}
}

func (s *MetricsSeriesUnit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMetricsSeriesUnit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMetricsSeriesUnit(input string) (*MetricsSeriesUnit, error) {
	vals := map[string]MetricsSeriesUnit{
		"bitspersecond": MetricsSeriesUnitBitsPerSecond,
		"bytes":         MetricsSeriesUnitBytes,
		"count":         MetricsSeriesUnitCount,
		"milliseconds":  MetricsSeriesUnitMilliSeconds,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MetricsSeriesUnit(input)
	return &out, nil
}
