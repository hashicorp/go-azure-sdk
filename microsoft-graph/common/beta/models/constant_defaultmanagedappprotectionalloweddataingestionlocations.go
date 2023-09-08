package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAllowedDataIngestionLocations string

const (
	DefaultManagedAppProtectionAllowedDataIngestionLocationscamera              DefaultManagedAppProtectionAllowedDataIngestionLocations = "Camera"
	DefaultManagedAppProtectionAllowedDataIngestionLocationsoneDriveForBusiness DefaultManagedAppProtectionAllowedDataIngestionLocations = "OneDriveForBusiness"
	DefaultManagedAppProtectionAllowedDataIngestionLocationsphotoLibrary        DefaultManagedAppProtectionAllowedDataIngestionLocations = "PhotoLibrary"
	DefaultManagedAppProtectionAllowedDataIngestionLocationssharePoint          DefaultManagedAppProtectionAllowedDataIngestionLocations = "SharePoint"
)

func PossibleValuesForDefaultManagedAppProtectionAllowedDataIngestionLocations() []string {
	return []string{
		string(DefaultManagedAppProtectionAllowedDataIngestionLocationscamera),
		string(DefaultManagedAppProtectionAllowedDataIngestionLocationsoneDriveForBusiness),
		string(DefaultManagedAppProtectionAllowedDataIngestionLocationsphotoLibrary),
		string(DefaultManagedAppProtectionAllowedDataIngestionLocationssharePoint),
	}
}

func parseDefaultManagedAppProtectionAllowedDataIngestionLocations(input string) (*DefaultManagedAppProtectionAllowedDataIngestionLocations, error) {
	vals := map[string]DefaultManagedAppProtectionAllowedDataIngestionLocations{
		"camera":              DefaultManagedAppProtectionAllowedDataIngestionLocationscamera,
		"onedriveforbusiness": DefaultManagedAppProtectionAllowedDataIngestionLocationsoneDriveForBusiness,
		"photolibrary":        DefaultManagedAppProtectionAllowedDataIngestionLocationsphotoLibrary,
		"sharepoint":          DefaultManagedAppProtectionAllowedDataIngestionLocationssharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAllowedDataIngestionLocations(input)
	return &out, nil
}
