package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAllowedDataIngestionLocations string

const (
	TargetedManagedAppProtectionAllowedDataIngestionLocations_Camera              TargetedManagedAppProtectionAllowedDataIngestionLocations = "camera"
	TargetedManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness TargetedManagedAppProtectionAllowedDataIngestionLocations = "oneDriveForBusiness"
	TargetedManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary        TargetedManagedAppProtectionAllowedDataIngestionLocations = "photoLibrary"
	TargetedManagedAppProtectionAllowedDataIngestionLocations_SharePoint          TargetedManagedAppProtectionAllowedDataIngestionLocations = "sharePoint"
)

func PossibleValuesForTargetedManagedAppProtectionAllowedDataIngestionLocations() []string {
	return []string{
		string(TargetedManagedAppProtectionAllowedDataIngestionLocations_Camera),
		string(TargetedManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness),
		string(TargetedManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary),
		string(TargetedManagedAppProtectionAllowedDataIngestionLocations_SharePoint),
	}
}

func (s *TargetedManagedAppProtectionAllowedDataIngestionLocations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionAllowedDataIngestionLocations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionAllowedDataIngestionLocations(input string) (*TargetedManagedAppProtectionAllowedDataIngestionLocations, error) {
	vals := map[string]TargetedManagedAppProtectionAllowedDataIngestionLocations{
		"camera":              TargetedManagedAppProtectionAllowedDataIngestionLocations_Camera,
		"onedriveforbusiness": TargetedManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness,
		"photolibrary":        TargetedManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary,
		"sharepoint":          TargetedManagedAppProtectionAllowedDataIngestionLocations_SharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAllowedDataIngestionLocations(input)
	return &out, nil
}
