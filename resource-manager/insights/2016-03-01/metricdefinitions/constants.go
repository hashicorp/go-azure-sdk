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

type Unit string

const (
	UnitBitsPerSecond  Unit = "BitsPerSecond"
	UnitByteSeconds    Unit = "ByteSeconds"
	UnitBytes          Unit = "Bytes"
	UnitBytesPerSecond Unit = "BytesPerSecond"
	UnitCores          Unit = "Cores"
	UnitCount          Unit = "Count"
	UnitCountPerSecond Unit = "CountPerSecond"
	UnitMilliCores     Unit = "MilliCores"
	UnitMilliSeconds   Unit = "MilliSeconds"
	UnitNanoCores      Unit = "NanoCores"
	UnitPercent        Unit = "Percent"
	UnitSeconds        Unit = "Seconds"
	UnitUnspecified    Unit = "Unspecified"
)

func PossibleValuesForUnit() []string {
	return []string{
		string(UnitBitsPerSecond),
		string(UnitByteSeconds),
		string(UnitBytes),
		string(UnitBytesPerSecond),
		string(UnitCores),
		string(UnitCount),
		string(UnitCountPerSecond),
		string(UnitMilliCores),
		string(UnitMilliSeconds),
		string(UnitNanoCores),
		string(UnitPercent),
		string(UnitSeconds),
		string(UnitUnspecified),
	}
}

func (s *Unit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUnit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUnit(input string) (*Unit, error) {
	vals := map[string]Unit{
		"bitspersecond":  UnitBitsPerSecond,
		"byteseconds":    UnitByteSeconds,
		"bytes":          UnitBytes,
		"bytespersecond": UnitBytesPerSecond,
		"cores":          UnitCores,
		"count":          UnitCount,
		"countpersecond": UnitCountPerSecond,
		"millicores":     UnitMilliCores,
		"milliseconds":   UnitMilliSeconds,
		"nanocores":      UnitNanoCores,
		"percent":        UnitPercent,
		"seconds":        UnitSeconds,
		"unspecified":    UnitUnspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Unit(input)
	return &out, nil
}
