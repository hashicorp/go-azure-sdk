package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAllowedDataIngestionLocations string

const (
	TargetedManagedAppProtectionAllowedDataIngestionLocationscamera              TargetedManagedAppProtectionAllowedDataIngestionLocations = "Camera"
	TargetedManagedAppProtectionAllowedDataIngestionLocationsoneDriveForBusiness TargetedManagedAppProtectionAllowedDataIngestionLocations = "OneDriveForBusiness"
	TargetedManagedAppProtectionAllowedDataIngestionLocationsphotoLibrary        TargetedManagedAppProtectionAllowedDataIngestionLocations = "PhotoLibrary"
	TargetedManagedAppProtectionAllowedDataIngestionLocationssharePoint          TargetedManagedAppProtectionAllowedDataIngestionLocations = "SharePoint"
)

func PossibleValuesForTargetedManagedAppProtectionAllowedDataIngestionLocations() []string {
	return []string{
		string(TargetedManagedAppProtectionAllowedDataIngestionLocationscamera),
		string(TargetedManagedAppProtectionAllowedDataIngestionLocationsoneDriveForBusiness),
		string(TargetedManagedAppProtectionAllowedDataIngestionLocationsphotoLibrary),
		string(TargetedManagedAppProtectionAllowedDataIngestionLocationssharePoint),
	}
}

func parseTargetedManagedAppProtectionAllowedDataIngestionLocations(input string) (*TargetedManagedAppProtectionAllowedDataIngestionLocations, error) {
	vals := map[string]TargetedManagedAppProtectionAllowedDataIngestionLocations{
		"camera":              TargetedManagedAppProtectionAllowedDataIngestionLocationscamera,
		"onedriveforbusiness": TargetedManagedAppProtectionAllowedDataIngestionLocationsoneDriveForBusiness,
		"photolibrary":        TargetedManagedAppProtectionAllowedDataIngestionLocationsphotoLibrary,
		"sharepoint":          TargetedManagedAppProtectionAllowedDataIngestionLocationssharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAllowedDataIngestionLocations(input)
	return &out, nil
}
