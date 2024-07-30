package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionAllowedDataIngestionLocations string

const (
	ManagedAppProtectionAllowedDataIngestionLocations_Camera              ManagedAppProtectionAllowedDataIngestionLocations = "camera"
	ManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness ManagedAppProtectionAllowedDataIngestionLocations = "oneDriveForBusiness"
	ManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary        ManagedAppProtectionAllowedDataIngestionLocations = "photoLibrary"
	ManagedAppProtectionAllowedDataIngestionLocations_SharePoint          ManagedAppProtectionAllowedDataIngestionLocations = "sharePoint"
)

func PossibleValuesForManagedAppProtectionAllowedDataIngestionLocations() []string {
	return []string{
		string(ManagedAppProtectionAllowedDataIngestionLocations_Camera),
		string(ManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness),
		string(ManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary),
		string(ManagedAppProtectionAllowedDataIngestionLocations_SharePoint),
	}
}

func (s *ManagedAppProtectionAllowedDataIngestionLocations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionAllowedDataIngestionLocations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionAllowedDataIngestionLocations(input string) (*ManagedAppProtectionAllowedDataIngestionLocations, error) {
	vals := map[string]ManagedAppProtectionAllowedDataIngestionLocations{
		"camera":              ManagedAppProtectionAllowedDataIngestionLocations_Camera,
		"onedriveforbusiness": ManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness,
		"photolibrary":        ManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary,
		"sharepoint":          ManagedAppProtectionAllowedDataIngestionLocations_SharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionAllowedDataIngestionLocations(input)
	return &out, nil
}
