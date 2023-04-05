package vaultusages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsagesUnit string

const (
	UsagesUnitBytes          UsagesUnit = "Bytes"
	UsagesUnitBytesPerSecond UsagesUnit = "BytesPerSecond"
	UsagesUnitCount          UsagesUnit = "Count"
	UsagesUnitCountPerSecond UsagesUnit = "CountPerSecond"
	UsagesUnitPercent        UsagesUnit = "Percent"
	UsagesUnitSeconds        UsagesUnit = "Seconds"
)

func PossibleValuesForUsagesUnit() []string {
	return []string{
		string(UsagesUnitBytes),
		string(UsagesUnitBytesPerSecond),
		string(UsagesUnitCount),
		string(UsagesUnitCountPerSecond),
		string(UsagesUnitPercent),
		string(UsagesUnitSeconds),
	}
}
