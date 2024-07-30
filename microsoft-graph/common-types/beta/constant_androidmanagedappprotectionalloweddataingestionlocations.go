package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAllowedDataIngestionLocations string

const (
	AndroidManagedAppProtectionAllowedDataIngestionLocations_Camera              AndroidManagedAppProtectionAllowedDataIngestionLocations = "camera"
	AndroidManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness AndroidManagedAppProtectionAllowedDataIngestionLocations = "oneDriveForBusiness"
	AndroidManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary        AndroidManagedAppProtectionAllowedDataIngestionLocations = "photoLibrary"
	AndroidManagedAppProtectionAllowedDataIngestionLocations_SharePoint          AndroidManagedAppProtectionAllowedDataIngestionLocations = "sharePoint"
)

func PossibleValuesForAndroidManagedAppProtectionAllowedDataIngestionLocations() []string {
	return []string{
		string(AndroidManagedAppProtectionAllowedDataIngestionLocations_Camera),
		string(AndroidManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness),
		string(AndroidManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary),
		string(AndroidManagedAppProtectionAllowedDataIngestionLocations_SharePoint),
	}
}

func (s *AndroidManagedAppProtectionAllowedDataIngestionLocations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAllowedDataIngestionLocations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAllowedDataIngestionLocations(input string) (*AndroidManagedAppProtectionAllowedDataIngestionLocations, error) {
	vals := map[string]AndroidManagedAppProtectionAllowedDataIngestionLocations{
		"camera":              AndroidManagedAppProtectionAllowedDataIngestionLocations_Camera,
		"onedriveforbusiness": AndroidManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness,
		"photolibrary":        AndroidManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary,
		"sharepoint":          AndroidManagedAppProtectionAllowedDataIngestionLocations_SharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAllowedDataIngestionLocations(input)
	return &out, nil
}
