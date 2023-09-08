package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionAllowedDataStorageLocations string

const (
	ManagedAppProtectionAllowedDataStorageLocationsbox                 ManagedAppProtectionAllowedDataStorageLocations = "Box"
	ManagedAppProtectionAllowedDataStorageLocationslocalStorage        ManagedAppProtectionAllowedDataStorageLocations = "LocalStorage"
	ManagedAppProtectionAllowedDataStorageLocationsoneDriveForBusiness ManagedAppProtectionAllowedDataStorageLocations = "OneDriveForBusiness"
	ManagedAppProtectionAllowedDataStorageLocationsphotoLibrary        ManagedAppProtectionAllowedDataStorageLocations = "PhotoLibrary"
	ManagedAppProtectionAllowedDataStorageLocationssharePoint          ManagedAppProtectionAllowedDataStorageLocations = "SharePoint"
)

func PossibleValuesForManagedAppProtectionAllowedDataStorageLocations() []string {
	return []string{
		string(ManagedAppProtectionAllowedDataStorageLocationsbox),
		string(ManagedAppProtectionAllowedDataStorageLocationslocalStorage),
		string(ManagedAppProtectionAllowedDataStorageLocationsoneDriveForBusiness),
		string(ManagedAppProtectionAllowedDataStorageLocationsphotoLibrary),
		string(ManagedAppProtectionAllowedDataStorageLocationssharePoint),
	}
}

func parseManagedAppProtectionAllowedDataStorageLocations(input string) (*ManagedAppProtectionAllowedDataStorageLocations, error) {
	vals := map[string]ManagedAppProtectionAllowedDataStorageLocations{
		"box":                 ManagedAppProtectionAllowedDataStorageLocationsbox,
		"localstorage":        ManagedAppProtectionAllowedDataStorageLocationslocalStorage,
		"onedriveforbusiness": ManagedAppProtectionAllowedDataStorageLocationsoneDriveForBusiness,
		"photolibrary":        ManagedAppProtectionAllowedDataStorageLocationsphotoLibrary,
		"sharepoint":          ManagedAppProtectionAllowedDataStorageLocationssharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionAllowedDataStorageLocations(input)
	return &out, nil
}
