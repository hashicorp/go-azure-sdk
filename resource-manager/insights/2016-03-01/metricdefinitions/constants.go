package metricdefinitions

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
