package backupresourcestorageconfigs

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageType string

const (
	StorageTypeGeoRedundant               StorageType = "GeoRedundant"
	StorageTypeInvalid                    StorageType = "Invalid"
	StorageTypeLocallyRedundant           StorageType = "LocallyRedundant"
	StorageTypeReadAccessGeoZoneRedundant StorageType = "ReadAccessGeoZoneRedundant"
	StorageTypeZoneRedundant              StorageType = "ZoneRedundant"
)

func PossibleValuesForStorageType() []string {
	return []string{
		string(StorageTypeGeoRedundant),
		string(StorageTypeInvalid),
		string(StorageTypeLocallyRedundant),
		string(StorageTypeReadAccessGeoZoneRedundant),
		string(StorageTypeZoneRedundant),
	}
}

func parseStorageType(input string) (*StorageType, error) {
	vals := map[string]StorageType{
		"georedundant":               StorageTypeGeoRedundant,
		"invalid":                    StorageTypeInvalid,
		"locallyredundant":           StorageTypeLocallyRedundant,
		"readaccessgeozoneredundant": StorageTypeReadAccessGeoZoneRedundant,
		"zoneredundant":              StorageTypeZoneRedundant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageType(input)
	return &out, nil
}

type StorageTypeState string

const (
	StorageTypeStateInvalid  StorageTypeState = "Invalid"
	StorageTypeStateLocked   StorageTypeState = "Locked"
	StorageTypeStateUnlocked StorageTypeState = "Unlocked"
)

func PossibleValuesForStorageTypeState() []string {
	return []string{
		string(StorageTypeStateInvalid),
		string(StorageTypeStateLocked),
		string(StorageTypeStateUnlocked),
	}
}

func parseStorageTypeState(input string) (*StorageTypeState, error) {
	vals := map[string]StorageTypeState{
		"invalid":  StorageTypeStateInvalid,
		"locked":   StorageTypeStateLocked,
		"unlocked": StorageTypeStateUnlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageTypeState(input)
	return &out, nil
}
