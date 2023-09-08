package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAllowedDataStorageLocations string

const (
	AndroidManagedAppProtectionAllowedDataStorageLocationsbox                 AndroidManagedAppProtectionAllowedDataStorageLocations = "Box"
	AndroidManagedAppProtectionAllowedDataStorageLocationslocalStorage        AndroidManagedAppProtectionAllowedDataStorageLocations = "LocalStorage"
	AndroidManagedAppProtectionAllowedDataStorageLocationsoneDriveForBusiness AndroidManagedAppProtectionAllowedDataStorageLocations = "OneDriveForBusiness"
	AndroidManagedAppProtectionAllowedDataStorageLocationsphotoLibrary        AndroidManagedAppProtectionAllowedDataStorageLocations = "PhotoLibrary"
	AndroidManagedAppProtectionAllowedDataStorageLocationssharePoint          AndroidManagedAppProtectionAllowedDataStorageLocations = "SharePoint"
)

func PossibleValuesForAndroidManagedAppProtectionAllowedDataStorageLocations() []string {
	return []string{
		string(AndroidManagedAppProtectionAllowedDataStorageLocationsbox),
		string(AndroidManagedAppProtectionAllowedDataStorageLocationslocalStorage),
		string(AndroidManagedAppProtectionAllowedDataStorageLocationsoneDriveForBusiness),
		string(AndroidManagedAppProtectionAllowedDataStorageLocationsphotoLibrary),
		string(AndroidManagedAppProtectionAllowedDataStorageLocationssharePoint),
	}
}

func parseAndroidManagedAppProtectionAllowedDataStorageLocations(input string) (*AndroidManagedAppProtectionAllowedDataStorageLocations, error) {
	vals := map[string]AndroidManagedAppProtectionAllowedDataStorageLocations{
		"box":                 AndroidManagedAppProtectionAllowedDataStorageLocationsbox,
		"localstorage":        AndroidManagedAppProtectionAllowedDataStorageLocationslocalStorage,
		"onedriveforbusiness": AndroidManagedAppProtectionAllowedDataStorageLocationsoneDriveForBusiness,
		"photolibrary":        AndroidManagedAppProtectionAllowedDataStorageLocationsphotoLibrary,
		"sharepoint":          AndroidManagedAppProtectionAllowedDataStorageLocationssharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAllowedDataStorageLocations(input)
	return &out, nil
}
