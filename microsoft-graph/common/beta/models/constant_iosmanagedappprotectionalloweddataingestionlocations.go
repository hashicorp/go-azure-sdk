package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAllowedDataIngestionLocations string

const (
	IosManagedAppProtectionAllowedDataIngestionLocationscamera              IosManagedAppProtectionAllowedDataIngestionLocations = "Camera"
	IosManagedAppProtectionAllowedDataIngestionLocationsoneDriveForBusiness IosManagedAppProtectionAllowedDataIngestionLocations = "OneDriveForBusiness"
	IosManagedAppProtectionAllowedDataIngestionLocationsphotoLibrary        IosManagedAppProtectionAllowedDataIngestionLocations = "PhotoLibrary"
	IosManagedAppProtectionAllowedDataIngestionLocationssharePoint          IosManagedAppProtectionAllowedDataIngestionLocations = "SharePoint"
)

func PossibleValuesForIosManagedAppProtectionAllowedDataIngestionLocations() []string {
	return []string{
		string(IosManagedAppProtectionAllowedDataIngestionLocationscamera),
		string(IosManagedAppProtectionAllowedDataIngestionLocationsoneDriveForBusiness),
		string(IosManagedAppProtectionAllowedDataIngestionLocationsphotoLibrary),
		string(IosManagedAppProtectionAllowedDataIngestionLocationssharePoint),
	}
}

func parseIosManagedAppProtectionAllowedDataIngestionLocations(input string) (*IosManagedAppProtectionAllowedDataIngestionLocations, error) {
	vals := map[string]IosManagedAppProtectionAllowedDataIngestionLocations{
		"camera":              IosManagedAppProtectionAllowedDataIngestionLocationscamera,
		"onedriveforbusiness": IosManagedAppProtectionAllowedDataIngestionLocationsoneDriveForBusiness,
		"photolibrary":        IosManagedAppProtectionAllowedDataIngestionLocationsphotoLibrary,
		"sharepoint":          IosManagedAppProtectionAllowedDataIngestionLocationssharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAllowedDataIngestionLocations(input)
	return &out, nil
}
