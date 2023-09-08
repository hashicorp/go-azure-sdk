package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAllowedDataStorageLocations string

const (
	TargetedManagedAppProtectionAllowedDataStorageLocationsbox                 TargetedManagedAppProtectionAllowedDataStorageLocations = "Box"
	TargetedManagedAppProtectionAllowedDataStorageLocationslocalStorage        TargetedManagedAppProtectionAllowedDataStorageLocations = "LocalStorage"
	TargetedManagedAppProtectionAllowedDataStorageLocationsoneDriveForBusiness TargetedManagedAppProtectionAllowedDataStorageLocations = "OneDriveForBusiness"
	TargetedManagedAppProtectionAllowedDataStorageLocationsphotoLibrary        TargetedManagedAppProtectionAllowedDataStorageLocations = "PhotoLibrary"
	TargetedManagedAppProtectionAllowedDataStorageLocationssharePoint          TargetedManagedAppProtectionAllowedDataStorageLocations = "SharePoint"
)

func PossibleValuesForTargetedManagedAppProtectionAllowedDataStorageLocations() []string {
	return []string{
		string(TargetedManagedAppProtectionAllowedDataStorageLocationsbox),
		string(TargetedManagedAppProtectionAllowedDataStorageLocationslocalStorage),
		string(TargetedManagedAppProtectionAllowedDataStorageLocationsoneDriveForBusiness),
		string(TargetedManagedAppProtectionAllowedDataStorageLocationsphotoLibrary),
		string(TargetedManagedAppProtectionAllowedDataStorageLocationssharePoint),
	}
}

func parseTargetedManagedAppProtectionAllowedDataStorageLocations(input string) (*TargetedManagedAppProtectionAllowedDataStorageLocations, error) {
	vals := map[string]TargetedManagedAppProtectionAllowedDataStorageLocations{
		"box":                 TargetedManagedAppProtectionAllowedDataStorageLocationsbox,
		"localstorage":        TargetedManagedAppProtectionAllowedDataStorageLocationslocalStorage,
		"onedriveforbusiness": TargetedManagedAppProtectionAllowedDataStorageLocationsoneDriveForBusiness,
		"photolibrary":        TargetedManagedAppProtectionAllowedDataStorageLocationsphotoLibrary,
		"sharepoint":          TargetedManagedAppProtectionAllowedDataStorageLocationssharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAllowedDataStorageLocations(input)
	return &out, nil
}
