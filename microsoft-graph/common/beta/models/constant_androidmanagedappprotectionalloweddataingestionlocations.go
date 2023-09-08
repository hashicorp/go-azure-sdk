package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAllowedDataIngestionLocations string

const (
	AndroidManagedAppProtectionAllowedDataIngestionLocationscamera              AndroidManagedAppProtectionAllowedDataIngestionLocations = "Camera"
	AndroidManagedAppProtectionAllowedDataIngestionLocationsoneDriveForBusiness AndroidManagedAppProtectionAllowedDataIngestionLocations = "OneDriveForBusiness"
	AndroidManagedAppProtectionAllowedDataIngestionLocationsphotoLibrary        AndroidManagedAppProtectionAllowedDataIngestionLocations = "PhotoLibrary"
	AndroidManagedAppProtectionAllowedDataIngestionLocationssharePoint          AndroidManagedAppProtectionAllowedDataIngestionLocations = "SharePoint"
)

func PossibleValuesForAndroidManagedAppProtectionAllowedDataIngestionLocations() []string {
	return []string{
		string(AndroidManagedAppProtectionAllowedDataIngestionLocationscamera),
		string(AndroidManagedAppProtectionAllowedDataIngestionLocationsoneDriveForBusiness),
		string(AndroidManagedAppProtectionAllowedDataIngestionLocationsphotoLibrary),
		string(AndroidManagedAppProtectionAllowedDataIngestionLocationssharePoint),
	}
}

func parseAndroidManagedAppProtectionAllowedDataIngestionLocations(input string) (*AndroidManagedAppProtectionAllowedDataIngestionLocations, error) {
	vals := map[string]AndroidManagedAppProtectionAllowedDataIngestionLocations{
		"camera":              AndroidManagedAppProtectionAllowedDataIngestionLocationscamera,
		"onedriveforbusiness": AndroidManagedAppProtectionAllowedDataIngestionLocationsoneDriveForBusiness,
		"photolibrary":        AndroidManagedAppProtectionAllowedDataIngestionLocationsphotoLibrary,
		"sharepoint":          AndroidManagedAppProtectionAllowedDataIngestionLocationssharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAllowedDataIngestionLocations(input)
	return &out, nil
}
