package metrics

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricResultType string

const (
	MetricResultTypeData     MetricResultType = "Data"
	MetricResultTypeMetadata MetricResultType = "Metadata"
)

func PossibleValuesForMetricResultType() []string {
	return []string{
		string(MetricResultTypeData),
		string(MetricResultTypeMetadata),
	}
}

func parseMetricResultType(input string) (*MetricResultType, error) {
	vals := map[string]MetricResultType{
		"data":     MetricResultTypeData,
		"metadata": MetricResultTypeMetadata,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MetricResultType(input)
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

type ResultType string

const (
	ResultTypeData     ResultType = "Data"
	ResultTypeMetadata ResultType = "Metadata"
)

func PossibleValuesForResultType() []string {
	return []string{
		string(ResultTypeData),
		string(ResultTypeMetadata),
	}
}

func parseResultType(input string) (*ResultType, error) {
	vals := map[string]ResultType{
		"data":     ResultTypeData,
		"metadata": ResultTypeMetadata,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResultType(input)
	return &out, nil
}
