package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAllowedDataIngestionLocations string

const (
	DefaultManagedAppProtectionAllowedDataIngestionLocations_Camera              DefaultManagedAppProtectionAllowedDataIngestionLocations = "camera"
	DefaultManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness DefaultManagedAppProtectionAllowedDataIngestionLocations = "oneDriveForBusiness"
	DefaultManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary        DefaultManagedAppProtectionAllowedDataIngestionLocations = "photoLibrary"
	DefaultManagedAppProtectionAllowedDataIngestionLocations_SharePoint          DefaultManagedAppProtectionAllowedDataIngestionLocations = "sharePoint"
)

func PossibleValuesForDefaultManagedAppProtectionAllowedDataIngestionLocations() []string {
	return []string{
		string(DefaultManagedAppProtectionAllowedDataIngestionLocations_Camera),
		string(DefaultManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness),
		string(DefaultManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary),
		string(DefaultManagedAppProtectionAllowedDataIngestionLocations_SharePoint),
	}
}

func (s *DefaultManagedAppProtectionAllowedDataIngestionLocations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAllowedDataIngestionLocations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAllowedDataIngestionLocations(input string) (*DefaultManagedAppProtectionAllowedDataIngestionLocations, error) {
	vals := map[string]DefaultManagedAppProtectionAllowedDataIngestionLocations{
		"camera":              DefaultManagedAppProtectionAllowedDataIngestionLocations_Camera,
		"onedriveforbusiness": DefaultManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness,
		"photolibrary":        DefaultManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary,
		"sharepoint":          DefaultManagedAppProtectionAllowedDataIngestionLocations_SharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAllowedDataIngestionLocations(input)
	return &out, nil
}
