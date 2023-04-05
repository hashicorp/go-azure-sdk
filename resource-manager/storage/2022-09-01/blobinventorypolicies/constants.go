package blobinventorypolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Format string

const (
	FormatCsv     Format = "Csv"
	FormatParquet Format = "Parquet"
)

func PossibleValuesForFormat() []string {
	return []string{
		string(FormatCsv),
		string(FormatParquet),
	}
}

type InventoryRuleType string

const (
	InventoryRuleTypeInventory InventoryRuleType = "Inventory"
)

func PossibleValuesForInventoryRuleType() []string {
	return []string{
		string(InventoryRuleTypeInventory),
	}
}

type ObjectType string

const (
	ObjectTypeBlob      ObjectType = "Blob"
	ObjectTypeContainer ObjectType = "Container"
)

func PossibleValuesForObjectType() []string {
	return []string{
		string(ObjectTypeBlob),
		string(ObjectTypeContainer),
	}
}

type Schedule string

const (
	ScheduleDaily  Schedule = "Daily"
	ScheduleWeekly Schedule = "Weekly"
)

func PossibleValuesForSchedule() []string {
	return []string{
		string(ScheduleDaily),
		string(ScheduleWeekly),
	}
}
