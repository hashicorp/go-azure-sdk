package metricdefinitions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AggregationType string

const (
	AggregationTypeAverage AggregationType = "Average"
	AggregationTypeCount   AggregationType = "Count"
	AggregationTypeMaximum AggregationType = "Maximum"
	AggregationTypeMinimum AggregationType = "Minimum"
	AggregationTypeNone    AggregationType = "None"
	AggregationTypeTotal   AggregationType = "Total"
)

func PossibleValuesForAggregationType() []string {
	return []string{
		string(AggregationTypeAverage),
		string(AggregationTypeCount),
		string(AggregationTypeMaximum),
		string(AggregationTypeMinimum),
		string(AggregationTypeNone),
		string(AggregationTypeTotal),
	}
}

func (s *AggregationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAggregationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAggregationType(input string) (*AggregationType, error) {
	vals := map[string]AggregationType{
		"average": AggregationTypeAverage,
		"count":   AggregationTypeCount,
		"maximum": AggregationTypeMaximum,
		"minimum": AggregationTypeMinimum,
		"none":    AggregationTypeNone,
		"total":   AggregationTypeTotal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AggregationType(input)
	return &out, nil
}

type MetricAggregationType string

const (
	MetricAggregationTypeAverage MetricAggregationType = "Average"
	MetricAggregationTypeCount   MetricAggregationType = "Count"
	MetricAggregationTypeMaximum MetricAggregationType = "Maximum"
	MetricAggregationTypeMinimum MetricAggregationType = "Minimum"
	MetricAggregationTypeNone    MetricAggregationType = "None"
	MetricAggregationTypeTotal   MetricAggregationType = "Total"
)

func PossibleValuesForMetricAggregationType() []string {
	return []string{
		string(MetricAggregationTypeAverage),
		string(MetricAggregationTypeCount),
		string(MetricAggregationTypeMaximum),
		string(MetricAggregationTypeMinimum),
		string(MetricAggregationTypeNone),
		string(MetricAggregationTypeTotal),
	}
}

func (s *MetricAggregationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMetricAggregationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMetricAggregationType(input string) (*MetricAggregationType, error) {
	vals := map[string]MetricAggregationType{
		"average": MetricAggregationTypeAverage,
		"count":   MetricAggregationTypeCount,
		"maximum": MetricAggregationTypeMaximum,
		"minimum": MetricAggregationTypeMinimum,
		"none":    MetricAggregationTypeNone,
		"total":   MetricAggregationTypeTotal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MetricAggregationType(input)
	return &out, nil
}

type MetricClass string

const (
	MetricClassAvailability MetricClass = "Availability"
	MetricClassErrors       MetricClass = "Errors"
	MetricClassLatency      MetricClass = "Latency"
	MetricClassSaturation   MetricClass = "Saturation"
	MetricClassTransactions MetricClass = "Transactions"
)

func PossibleValuesForMetricClass() []string {
	return []string{
		string(MetricClassAvailability),
		string(MetricClassErrors),
		string(MetricClassLatency),
		string(MetricClassSaturation),
		string(MetricClassTransactions),
	}
}

func (s *MetricClass) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMetricClass(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMetricClass(input string) (*MetricClass, error) {
	vals := map[string]MetricClass{
		"availability": MetricClassAvailability,
		"errors":       MetricClassErrors,
		"latency":      MetricClassLatency,
		"saturation":   MetricClassSaturation,
		"transactions": MetricClassTransactions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MetricClass(input)
	return &out, nil
}

type MetricUnit string

const (
	MetricUnitBitsPerSecond  MetricUnit = "BitsPerSecond"
	MetricUnitByteSeconds    MetricUnit = "ByteSeconds"
	MetricUnitBytes          MetricUnit = "Bytes"
	MetricUnitBytesPerSecond MetricUnit = "BytesPerSecond"
	MetricUnitCores          MetricUnit = "Cores"
	MetricUnitCount          MetricUnit = "Count"
	MetricUnitCountPerSecond MetricUnit = "CountPerSecond"
	MetricUnitMilliCores     MetricUnit = "MilliCores"
	MetricUnitMilliSeconds   MetricUnit = "MilliSeconds"
	MetricUnitNanoCores      MetricUnit = "NanoCores"
	MetricUnitPercent        MetricUnit = "Percent"
	MetricUnitSeconds        MetricUnit = "Seconds"
	MetricUnitUnspecified    MetricUnit = "Unspecified"
)

func PossibleValuesForMetricUnit() []string {
	return []string{
		string(MetricUnitBitsPerSecond),
		string(MetricUnitByteSeconds),
		string(MetricUnitBytes),
		string(MetricUnitBytesPerSecond),
		string(MetricUnitCores),
		string(MetricUnitCount),
		string(MetricUnitCountPerSecond),
		string(MetricUnitMilliCores),
		string(MetricUnitMilliSeconds),
		string(MetricUnitNanoCores),
		string(MetricUnitPercent),
		string(MetricUnitSeconds),
		string(MetricUnitUnspecified),
	}
}

func (s *MetricUnit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMetricUnit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMetricUnit(input string) (*MetricUnit, error) {
	vals := map[string]MetricUnit{
		"bitspersecond":  MetricUnitBitsPerSecond,
		"byteseconds":    MetricUnitByteSeconds,
		"bytes":          MetricUnitBytes,
		"bytespersecond": MetricUnitBytesPerSecond,
		"cores":          MetricUnitCores,
		"count":          MetricUnitCount,
		"countpersecond": MetricUnitCountPerSecond,
		"millicores":     MetricUnitMilliCores,
		"milliseconds":   MetricUnitMilliSeconds,
		"nanocores":      MetricUnitNanoCores,
		"percent":        MetricUnitPercent,
		"seconds":        MetricUnitSeconds,
		"unspecified":    MetricUnitUnspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MetricUnit(input)
	return &out, nil
}
