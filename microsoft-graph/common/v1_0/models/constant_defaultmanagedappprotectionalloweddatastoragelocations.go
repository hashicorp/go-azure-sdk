package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAllowedDataStorageLocations string

const (
	DefaultManagedAppProtectionAllowedDataStorageLocationsbox                 DefaultManagedAppProtectionAllowedDataStorageLocations = "Box"
	DefaultManagedAppProtectionAllowedDataStorageLocationslocalStorage        DefaultManagedAppProtectionAllowedDataStorageLocations = "LocalStorage"
	DefaultManagedAppProtectionAllowedDataStorageLocationsoneDriveForBusiness DefaultManagedAppProtectionAllowedDataStorageLocations = "OneDriveForBusiness"
	DefaultManagedAppProtectionAllowedDataStorageLocationssharePoint          DefaultManagedAppProtectionAllowedDataStorageLocations = "SharePoint"
)

func PossibleValuesForDefaultManagedAppProtectionAllowedDataStorageLocations() []string {
	return []string{
		string(DefaultManagedAppProtectionAllowedDataStorageLocationsbox),
		string(DefaultManagedAppProtectionAllowedDataStorageLocationslocalStorage),
		string(DefaultManagedAppProtectionAllowedDataStorageLocationsoneDriveForBusiness),
		string(DefaultManagedAppProtectionAllowedDataStorageLocationssharePoint),
	}
}

func parseDefaultManagedAppProtectionAllowedDataStorageLocations(input string) (*DefaultManagedAppProtectionAllowedDataStorageLocations, error) {
	vals := map[string]DefaultManagedAppProtectionAllowedDataStorageLocations{
		"box":                 DefaultManagedAppProtectionAllowedDataStorageLocationsbox,
		"localstorage":        DefaultManagedAppProtectionAllowedDataStorageLocationslocalStorage,
		"onedriveforbusiness": DefaultManagedAppProtectionAllowedDataStorageLocationsoneDriveForBusiness,
		"sharepoint":          DefaultManagedAppProtectionAllowedDataStorageLocationssharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAllowedDataStorageLocations(input)
	return &out, nil
}
