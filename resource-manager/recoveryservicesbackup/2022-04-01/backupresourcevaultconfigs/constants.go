package backupresourcevaultconfigs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnhancedSecurityState string

const (
	EnhancedSecurityStateDisabled EnhancedSecurityState = "Disabled"
	EnhancedSecurityStateEnabled  EnhancedSecurityState = "Enabled"
	EnhancedSecurityStateInvalid  EnhancedSecurityState = "Invalid"
)

func PossibleValuesForEnhancedSecurityState() []string {
	return []string{
		string(EnhancedSecurityStateDisabled),
		string(EnhancedSecurityStateEnabled),
		string(EnhancedSecurityStateInvalid),
	}
}

type SoftDeleteFeatureState string

const (
	SoftDeleteFeatureStateDisabled SoftDeleteFeatureState = "Disabled"
	SoftDeleteFeatureStateEnabled  SoftDeleteFeatureState = "Enabled"
	SoftDeleteFeatureStateInvalid  SoftDeleteFeatureState = "Invalid"
)

func PossibleValuesForSoftDeleteFeatureState() []string {
	return []string{
		string(SoftDeleteFeatureStateDisabled),
		string(SoftDeleteFeatureStateEnabled),
		string(SoftDeleteFeatureStateInvalid),
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
