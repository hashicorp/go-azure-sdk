package backupresourcestorageconfigsnoncrr

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DedupState string

const (
	DedupStateDisabled DedupState = "Disabled"
	DedupStateEnabled  DedupState = "Enabled"
	DedupStateInvalid  DedupState = "Invalid"
)

func PossibleValuesForDedupState() []string {
	return []string{
		string(DedupStateDisabled),
		string(DedupStateEnabled),
		string(DedupStateInvalid),
	}
}

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

type XcoolState string

const (
	XcoolStateDisabled XcoolState = "Disabled"
	XcoolStateEnabled  XcoolState = "Enabled"
	XcoolStateInvalid  XcoolState = "Invalid"
)

func PossibleValuesForXcoolState() []string {
	return []string{
		string(XcoolStateDisabled),
		string(XcoolStateEnabled),
		string(XcoolStateInvalid),
	}
}
